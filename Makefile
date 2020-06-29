all: generate build

build:
	go build -ldflags="-w -s"

generate:
	go generate

clean:
	rm -rf pkged.go scaffold out/
