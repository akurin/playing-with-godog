build:
	GO111MODULE=on go build -o bin/app cmd/app/main.go

run:
	GO111MODULE=on go run cmd/app/main.go

test:
	GO111MODULE=on go test ./...
