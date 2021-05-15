run:
	echo "Run triggered"
	echo "Golang rocks"
	go run main.go

build:
	env GOOS=linux go build -o bin/api main.go


# TODO : before building the binary for linux display on the screen "Building for linux"