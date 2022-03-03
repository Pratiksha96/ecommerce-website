build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build main.go

run: build
	brew services start mongodb-community && go run main.go start-server