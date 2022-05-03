build:
	@mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/ingress-controller ./cmd/caddy

build-docker: build
	docker build -t $(tag) -t $(tag2) -f Dockerfile.tilt .

copy-dev:
	cp -r -v ../ingress-config/errors ./errors
	cp -v ../ingress-config/GeoIP2-Country.mmdb  ./data

build-dev: copy-dev build