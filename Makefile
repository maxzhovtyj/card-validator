generator-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto

DEPLOY_USER=root
DEPLOY_HOST=194.164.59.123
DEPLOY_PATH=/root/card-validator

card-validator-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/card-validator-linux-amd64 ./cmd/server/

deploy-card-validator: card-validator-linux
	rsync ./config/config.yml ./deploy.sh ./bin/card-validator-linux-amd64 $(DEPLOY_USER)@$(DEPLOY_HOST):$(DEPLOY_PATH)
	ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null $(DEPLOY_USER)@$(DEPLOY_HOST) "cd $(DEPLOY_PATH) && bash ./deploy.sh"

test-api:
	go run ./cmd/client/main.go

start:
	docker build -t card-validator .
	docker stop validator
	docker rm validator
	docker run --name validator -d -p 7799:7799 card-validator
