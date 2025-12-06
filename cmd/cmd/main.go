package main

import (
	"net/http"

	"github.com/gliderlabs/ssh"
	"github.com/malivvan/cui/server"
	"github.com/malivvan/cui/service"
)

func main() {
	s, err := service.NewServer(service.Config{
		TCP: 8080,
		SSH: &ssh.Server{
			Addr: ":8080",
		},
		HTTP: &http.Server{
			Addr: ":8080",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hel\n"))
			}),
		},
	})
	if err != nil {
		panic(err)
	}
	s.Wait()

}
