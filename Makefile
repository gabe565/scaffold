generate:
	go generate

build:
	go build -ldflags="-w -s"

clean:
	rm -rf pkged.go installer out/
