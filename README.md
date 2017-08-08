# somesvgtest with docker and nginx proxy

This is a test repo showing that nginx-proxy does not set svg content-type, discussion on https://github.com/jwilder/nginx-proxy/issues/691

Two services are started behind the proxy: first one serving svg from another nginx and the second serving it from a go server (that does not set content type).

`docker-compose up --build ` will launch all services, `checksvgs.sh` will curl svg from different hosts showing content-type: 

`curl -I -L -H "Host: mysvgfile.local" localhost/logo_interne1.svg` will show Content-Type: image/svg+xml because it is set by the inner nginx:
```
HTTP/1.1 200 OK
Server: nginx/1.11.8
Date: Tue, 08 Aug 2017 15:18:52 GMT
Content-Type: image/svg+xml
Content-Length: 6877
Connection: keep-alive
Last-Modified: Tue, 08 Aug 2017 15:17:11 GMT
ETag: "5989d5f7-1add"
Accept-Ranges: bytes
```
`curl -I -L -H "Host: goservesvg.local" localhost/circle.svg` will show text/xml because go has not set any content type and proxy nginx does not set it...
```
HTTP/1.1 200 OK
Server: nginx/1.11.8
Date: Tue, 08 Aug 2017 15:18:52 GMT
Content-Type: text/xml; charset=utf-8
Content-Length: 242
Connection: keep-alive
```