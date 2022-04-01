build:
	cd ./backend && go mod download && CGO_ENABLED=0 GOOS=linux go build main.go

run: build
	cd ./backend && brew services start mongodb-community && go run main.go start-server

test: test
	cd ./backend && go test ./...