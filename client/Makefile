build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build -o talos-demo

run: build
	./talos-demo

test:
	go test