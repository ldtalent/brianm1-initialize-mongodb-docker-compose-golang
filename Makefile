# Pull mongo image from docker hub
mongo:
	docker run --name mongo4 \
	-p 8081:27017 \
	-v mongo-data:/data/db \
	-e MONGO_INITDB_ROOT_USERNAME=root \
	-e MONGO_INITDB_ROOT_PASSWORD=rootpassword \
	-d mongo:4

serve:
	go run main.go

.PHONY: mongo serve