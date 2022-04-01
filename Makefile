build-server:
	cd ./backend && go mod download && CGO_ENABLED=0 GOOS=linux go build main.go

build-client: 
	cd ./frontend && npm install

run-server: build-server
	cd ./backend && brew services start mongodb-community && go run main.go start-server
	
run-client: 
	cd ./frontend && npm start

test-server: test-server
	cd ./backend && go test ./...
