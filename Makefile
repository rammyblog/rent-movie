build:
	go build -o bin/rent-movie

run: build
	./bin/rent-movie

seed: build
	 .bin/rent-movie --seed


test: 
	go test -v ./...