# Build
## Build binaries
```shell script
make build
```

## Build Deb Package
```shell script
apt install make devscripts debhelper build-essential dh-systemd
debuild -us -uc -b
```

# Usage
## Start nginx-js-challenge backend
```shell script
./nginx-js-challenge -address=unix:/run/nginx-js-challenge.sock
```
## Test nginx-js-challenge backend
```shell script
curl --unix-socket /run/nginx-js-challenge.sock http:/example.com
```

## Nginx configuration
The ./nginx dir contains the vhost configuration template.
