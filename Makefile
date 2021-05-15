run:
	echo "Run triggered"
	echo "Golang rocks"
	go run main.go

build:
	env GOOS=linux go build -o bin/api main.go