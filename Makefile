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
	serverless deploy \
	--aws-profile maxaldtools  \
	--allowedOrigin http://localhost:8080 \
	--stage dev \
	--ssmEncryptionKeyId 7a8dbbad-bae9-480c-af13-d7a14bb70a71 \
	--frontendBaseUrl http://localhost:8080 \
	--emailFrom maximilien.andile.demo@gmail.com \
	--adminEmails maximilien.andile.demo@gmail.com

genMocks:
	mockgen -source=internal/storage/interface.go -destination=internal/storage/mock.go -package=storage
	mockgen -source=internal/uniqueid/interface.go -destination=internal/uniqueid/mock.go -package=uniqueid