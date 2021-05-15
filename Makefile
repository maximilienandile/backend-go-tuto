run:
	echo "Run triggered"
	echo "Golang rocks"
	go run main.go

build:
	echo "Building for Linux"
	env GOOS=linux go build -o bin/api api/main.go
	env GOOS=linux go build -o bin/hello testLambda/main.go

deploy: build
	serverless deploy --aws-profile maxaldtools

deploy_dev: build
	serverless deploy --aws-profile maxaldtools --allowedOrigin http://localhost:8080 --stage dev

genMocks:
	mockgen -source=internal/storage/interface.go -destination=internal/storage/mock.go -package=storage