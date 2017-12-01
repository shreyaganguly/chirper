chirper: main.go assets.go
	go build

assets.go: assets/
	go-bindata -pkg packed -o packed/assets.go assets/...

clean:
	rm -rf packed || true
	go clean
