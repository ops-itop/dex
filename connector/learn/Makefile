all: build build-example
build:
	cd ../../cmd/dex && go build -mod=vendor -o ../../bin/dex
build-example:
	cd ../../cmd/example-app && go build -mod=vendor -o ../../bin/example-app
test:
	cp -ru ../../web .
	../../bin/dex serve config-dev.yaml &
	../../bin/example-app --listen http://0.0.0.0:5550 --redirect-uri http://127.0.0.1:5550/callback