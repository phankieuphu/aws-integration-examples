# App variables
APP_NAME := cicd-example
APP_PATH := ./main.go
IMAGE := $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/$(APP_NAME)
TAG := latest

all: build docker push

# Build Go binary
build:
	go mod tidy
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(APP_NAME) $(APP_PATH)

# Run local
run:
	go run $(APP_PATH)

# Build Docker image
docker:
	docker build -t $(IMAGE):$(TAG) .

# Push Docker image to ECR
push:
	aws ecr get-login-password --region $(AWS_REGION) | docker login --username AWS --password-stdin $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com
	docker push $(IMAGE):$(TAG)

clean:
	rm -rf bin
