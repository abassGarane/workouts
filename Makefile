build:
	@go build -o bin/muscles -ldflags "-w -s" cmd/web/*

run: build
	./bin/muscles

clean:
	rm -rf ./bin


test:
	go test -v -cover ./...

up:
	sudo docker-compose up -d

down:
	sudo docker-compose down

running:
	sudo docker ps

