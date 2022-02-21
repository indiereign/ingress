build:
	@mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/ingress-controller ./cmd/caddy

build-docker: build
	docker build -t $(tag) -t $(tag2) -f Dockerfile.tilt .
