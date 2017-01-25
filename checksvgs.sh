echo the following will show Content-Type: image/svg+xml because it is set by the "inner" nginx
curl -I -L -H "Host: mysvgfile.local" localhost/logo_interne1.svg

echo the following will show text/xml because go has not set any content type and proxy nginx does not set it...
curl -I -L -H "Host: goservesvg.local" localhost/circle.svg
