.PHONY: build test deps build-release package-release release

build:
	go build -o bin/sifaka main.go

test:
	go test -v ./...

deps:
	go get

release: build-release package-release
	@echo "Release build and packaged"

build-release:
	GOOS=darwin  GOARCH=amd64 go build -o release/osx-amd64/sifaka main.go
	#GOOS=darwin  GOARCH=arm64 go build -o release/osx-arm64/sifaka main.go
	GOOS=linux   GOARCH=amd64 go build -o release/linux-amd64/sifaka main.go

package-release:
	tar -czvf release/sifaka.osx-amd64.tar.gz --directory=release/osx-amd64/ sifaka
	#tar -czvf release/sifaka.osx-arm64.tar.gz --directory=release/osx-arm64/ sifaka
	tar -czvf release/sifaka.linux-amd64.tar.gz --directory=release/linux-amd64/ sifaka
