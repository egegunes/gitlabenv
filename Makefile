build:
	go build
clean:
	rm gitlabenv
release:
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/gitlabenv-$(shell git describe)-linux-amd64
	GOOS=linux GOARCH=386 go build -o dist/gitlabenv-$(shell git describe)-linux-386
	GOOS=darwin GOARCH=amd64 go build -o dist/gitlabenv-$(shell git describe)-darwin-amd64
	GOOS=darwin GOARCH=386 go build -o dist/gitlabenv-$(shell git describe)-darwin-386
	GOOS=windows GOARCH=amd64 go build -o dist/gitlabenv-$(shell git describe)-windows-amd64
	GOOS=windows GOARCH=386 go build -o dist/gitlabenv-$(shell git describe)-windows-386
