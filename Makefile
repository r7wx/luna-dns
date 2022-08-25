all:
	CGO_ENABLED=0 go build -trimpath -ldflags="-w -s" -o build/luna-dns cmd/luna-dns/main.go

clean:
	rm -rf dist