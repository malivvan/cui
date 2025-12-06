package service

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/gliderlabs/ssh"
	"github.com/malivvan/cui/service/reuse"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

type Handler func(net.Conn) error

type Config struct {
	TCP  uint16
	UDP  uint16
	UNIX string
	SSH  *ssh.Server
	GRPC *grpc.Server
	HTTP *http.Server
	TLS  *tls.Config
	ERR  func(error)
}

func (cfg Config) error(err error) {
	if cfg.ERR != nil {
		cfg.ERR(err)
	}
}

func (cfg Config) tcpAddr() string {
	return fmt.Sprintf(":%d", cfg.TCP)
}

func (cfg Config) udpAddr() string {
	return fmt.Sprintf(":%d", cfg.UDP)
}

func (cfg Config) unixAddr() string {
	return cfg.UNIX
}

type Server struct {
	wg  sync.WaitGroup
	die chan struct{}
	err chan error
	cfg Config
	tcp struct {
		net.Listener
		mux  cmux.CMux
		ssh  net.Listener
		grpc net.Listener
		http struct {
			net.Listener
			ws net.Listener
		}
		tls struct {
			mux   cmux.CMux
			ssh   net.Listener
			grpc  net.Listener
			https struct {
				net.Listener
				wss net.Listener
			}
		}
	}
}

type Listener struct {
	net.Listener
	mutex sync.Mutex
	conns []net.Conn
}

func NewServer(cfg Config) (s *Server, err error) {
	s = &Server{
		cfg: cfg,
		err: make(chan error, 1),
	}
	s.tcp.Listener, err = reuse.Listen("tcp", s.cfg.tcpAddr())
	if err != nil {
		return nil, err
	}
	log.Printf("Listening on %s", s.cfg.tcpAddr())
	s.tcp.mux = cmux.New(s.tcp.Listener)
	if s.cfg.SSH != nil {
		s.tcp.ssh = s.tcp.mux.Match(cmux.PrefixMatcher("SSH-"))
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			log.Printf("SSH server listening on %s", s.cfg.tcpAddr())
			s.err <- s.cfg.SSH.Serve(s.tcp.ssh)
		}()
	}
	if s.cfg.GRPC != nil {
		s.tcp.grpc = s.tcp.mux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			log.Printf("gRPC server listening on %s", s.cfg.tcpAddr())
			s.err <- s.cfg.GRPC.Serve(s.tcp.grpc)
		}()
	}
	if s.cfg.HTTP != nil {
		s.tcp.http.Listener = s.tcp.mux.Match(cmux.HTTP1Fast(), cmux.HTTP2())
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			log.Printf("HTTP server listening on %s", s.cfg.tcpAddr())
			s.err <- s.cfg.HTTP.Serve(s.tcp.http)
		}()
	}
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		log.Printf("TCP multiplexer listening on %s", s.cfg.tcpAddr())
		s.err <- s.tcp.mux.Serve()
	}()
	if s.cfg.TLS != nil {
		s.tcp.tls.mux = cmux.New(tls.NewListener(s.tcp.mux.Match(cmux.Any()), s.cfg.TLS))
		if s.cfg.SSH != nil {
			s.tcp.tls.ssh = s.tcp.tls.mux.Match(cmux.PrefixMatcher("SSH-"))
			s.wg.Add(1)
			go func() {
				defer s.wg.Done()
				log.Printf("SSH over TLS server listening on %s", s.cfg.tcpAddr())
				s.err <- s.cfg.SSH.Serve(s.tcp.tls.ssh)
			}()
		}
		if s.cfg.GRPC != nil {
			s.tcp.tls.grpc = s.tcp.tls.mux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
			s.wg.Add(1)
			go func() {
				defer s.wg.Done()
				log.Printf("gRPC over TLS server listening on %s", s.cfg.tcpAddr())
				s.err <- s.cfg.GRPC.Serve(s.tcp.tls.grpc)
			}()
		}
		if s.cfg.HTTP != nil {
			s.tcp.tls.https.Listener = s.tcp.tls.mux.Match(cmux.Any())
			s.wg.Add(1)
			go func() {
				defer s.wg.Done()
				log.Printf("HTTPS server listening on %s", s.cfg.tcpAddr())
				s.err <- s.cfg.HTTP.Serve(s.tcp.tls.https)
			}()
		}
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			log.Printf("TLS multiplexer listening on %s", s.cfg.tcpAddr())
			s.err <- s.tcp.tls.mux.Serve()
		}()
	}

	shutdown := func() {
		if s.tcp.mux != nil {
			s.tcp.mux.Close()
			s.tcp.mux = nil
		}
		if s.tcp.Listener != nil {
			_ = s.tcp.Listener.Close()
			s.tcp.Listener = nil
		}
	}
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		select {
		case err := <-s.err:
			if err != nil {
				s.cfg.error(err)
			}
			shutdown()

		case <-s.die:
			shutdown()
		}
	}()
	return s, nil
}

func (s *Server) Wait() {
	s.wg.Wait()
}
