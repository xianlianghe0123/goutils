check:
	go vet ./...

test:
	go test -cover ./...

cover:
	go tool cover -html coverprofile