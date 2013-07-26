PROJECT=project
VERSION=$(shell go run main.go -v)
ARCHIVE=$(PROJECT)-$(VERSION).tar.gz

release:
	go get -u ./...
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	tar -zcvf $(ARCHIVE) $(PROJECT)
	s3cmd --acl-public put $(ARCHIVE) s3://bucket
	rm $(ARCHIVE)
	git checkout -

test:
	go test ./...
