run:
	go run cmd/auction/main.go

mongo:
	docker container run --rm --name auctionsdb -d -p 27017:27017 mongo:8.0.8