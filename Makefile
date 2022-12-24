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
	--param="allowedOrigins=http://localhost:8080" \
	--stage dev \
	--param="ssmEncryptionKeyId=7a8dbbad-bae9-480c-af13-d7a14bb70a71" \
	--param="frontendBaseUrl=http://localhost:8080" \
	--param="emailFrom=maximilien.andile.demo@gmail.com" \
	--param="adminEmails=maximilien.andile.demo@gmail.com"

deploy_prod: build
	serverless deploy \
	--aws-profile maxaldtools  \
	--param="allowedOrigins=https://www.gopher-ecommerce.com,http://localhost:8080" \
	--stage prod \
	--param="ssmEncryptionKeyId=0370bb46-ea37-40fa-98f6-47af96d93599" \
	--param="frontendBaseUrl=https://gopher-ecommerce.com" \
	--param="emailFrom=maximilien.andile.demo@gmail.com" \
	--param="adminEmails=maximilien.andile.demo@gmail.com"

genMocks:
	mockgen -source=internal/storage/interface.go -destination=internal/storage/mock.go -package=storage
	mockgen -source=internal/uniqueid/interface.go -destination=internal/uniqueid/mock.go -package=uniqueid