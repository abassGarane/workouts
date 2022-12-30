build:
	@go build -o bin/muscles -ldflags "-w -s" cmd/web/*

run: build
	./bin/muscles

clean:
	rm -rf ./bin


test:
	go test -v -cover ./...
