# cui - A terminal UI framework in Go
- https://github.com/libp2p/go-reuseport/tree/40284dc62c72ec11b336d718c84852f0825c7732
- https://github.com/andybalholm/cascadia/tree/5263deb988702df34b4de5b8cd2fe53add4bea3d


## Package structure
- `cmd/` contains the cui application entry point.
- `editor/` contains a small text editor component.
- `internal/` contains internal helpers and utilities.
- `markup/` implements a custom markup language and rendering engine.
- `runtime/` fast and secure script language supporting routines and channels
- `service/` implements a background service manager.
- `storage/` implements a simple key-value storage engine.
- `terminal/` contains a terminal emulator component.

## Vendor candidates

```
require (
github.com/ScaleFT/sshkeys v0.0.0-20200327173127-6142f742bca5
github.com/danieljoos/wincred v1.1.0
github.com/davecgh/go-spew v1.1.1
github.com/dchest/blake2b v1.0.0
github.com/flynn/noise v1.0.0
github.com/godbus/dbus v4.1.0+incompatible
github.com/golang/protobuf v1.5.2 // indirect
github.com/google/go-cmp v0.5.6
github.com/keybase/go-keychain v0.0.0-20201121013009-976c83ec27a6
github.com/keybase/saltpack v0.0.0-20210611181147-9dd0a21addc6
github.com/keys-pub/secretservice v0.0.0-20200519003656-26e44b8df47f
github.com/kr/text v0.2.0 // indirect
github.com/pkg/errors v0.9.1
github.com/stretchr/objx v0.3.0 // indirect
github.com/stretchr/testify v1.7.0
github.com/tyler-smith/go-bip39 v1.1.0
github.com/vmihailenco/msgpack/v4 v4.3.12
github.com/vmihailenco/tagparser v0.1.2 // indirect
golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e
golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
golang.org/x/term v0.0.0-20210317153231-de623e64d2a6 // indirect
golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
google.golang.org/appengine v1.6.7 // indirect
google.golang.org/protobuf v1.27.1 // indirect
gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
```

### general
https://github.com/ajaxray/geek-life
-https://github.com/jplozf/gosh
- https://github.com/desertbit/grumble/blob/v1.1.3/app.go#L207
- https://github.com/nlepage/go-wasm-http-server?tab=readme-ov-file for service worker
- https://github.com/magodo/go-wasmww for workers and shared workers
- https://github.com/life4/gweb way newer
- https://github.com/tarndt/wasmws for websocket net.Conn

### webrtc
- https://github.com/suutaku/sshx/tree/master
- https://github.com/trzsz/trzsz-ssh
- https://github.com/cloudfoundry/socks5-proxy
- https://github.com/wencaiwulue/sshvpn

### widgets
- https://github.com/dimonomid/nerdlog

### misc
- https://github.com/phuslu/iploc?tab=readme-ov-file
- https://github.com/denisbrodbeck/machineid

## Packages

| path                  | repository                                                                                                                   | commit                                     |    license     |
|:----------------------|:-----------------------------------------------------------------------------------------------------------------------------|:-------------------------------------------|:--------------:|
| **/**                 | [codeberg.org/tslocum/cview](https://codeberg.org/tslocum/cview/src/commit/242e7c1f1b61a4b3722a1afb45ca1165aefa9a59)         | `242e7c1f1b61a4b3722a1afb45ca1165aefa9a59` |     `MIT`      |
| **/bind.go**          | [codeberg.org/tslocum/cbind](https://codeberg.org/tslocum/cbind/src/commit/5cd49d3cfccbe4eefaab8a5282826aa95100aa42)         | `5cd49d3cfccbe4eefaab8a5282826aa95100aa42` |     `MIT`      |
| **/vte/**             | [git.sr.ht/~rockorager/tcell-term](https://git.sr.ht/~rockorager/tcell-term/commit/6805a6d75db82c2e1f51c1fb97170c26daf7aea0) | `6805a6d75db82c2e1f51c1fb97170c26daf7aea0` |     `MIT`      |
| **/vte/pty**          | [github.com/wellcomez/aiopty](https://github.com/wellcomez/aiopty/tree/afbcf1124b2cb834b75236fd0a5e6b56a2790e03)             | `afbcf1124b2cb834b75236fd0a5e6b56a2790e03` |     `MIT`      |
| **/vte/term**         | [github.com/wellcomez/aiopty](https://github.com/wellcomez/aiopty/tree/afbcf1124b2cb834b75236fd0a5e6b56a2790e03)             | `afbcf1124b2cb834b75236fd0a5e6b56a2790e03` |     `MIT`      |
| **/vte/ioctl**        | [github.com/wellcomez/aiopty](https://github.com/wellcomez/aiopty/tree/afbcf1124b2cb834b75236fd0a5e6b56a2790e03)             | `afbcf1124b2cb834b75236fd0a5e6b56a2790e03` |     `MIT`      |
| **/editor**           | [github.com/wellcomez/femto](https://github.com/wellcomez/femto/tree/8413a0288bcb042fd0de5cbbcb9893c16a01ee69)               | `8413a0288bcb042fd0de5cbbcb9893c16a01ee69` |     `MIT`      |                                     
| **/chart**            | [github.com/navidys/tvxwidgets](https://github.com/navidys/tvxwidgets/commit/96bcc0450684693eebd4f8e3e95fcc40eae2dbaa)       | `96bcc0450684693eebd4f8e3e95fcc40eae2dbaa` |     `MIT`      | 
| **/internal/etree**   | [github.com/beevik/etree](https://github.com/beevik/etree/tree/4032e04c8f2e2f35e43ce5d772fcef14a5df4d74)                     | `4032e04c8f2e2f35e43ce5d772fcef14a5df4d74` | `BSD-2-Clause` |


## tcell

- https://github.com/gdamore/tcell/tree/2c42625c556c4d03dca885420270723bb1829db1
- add encoding package
- https://github.com/gdamore/encoding/tree/6770ff7f5dae83f6e1bec40dc177c0f347df5139
- backport changes from micro tcell
- https://github.com/micro-editor/tcell/tree/20b75e27dba7691d064bda718a9c3a68d760531d