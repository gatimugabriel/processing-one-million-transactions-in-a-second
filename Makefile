BINARY_NAME=1mts
SEED_BINARY=seeder
GOBUILD=go build
GOTEST=go test
GOCLEAN=go clean

.PHONY: seed test build run clean

build:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/transactions

seed:
	$(GOBUILD) -o $(SEED_BINARY) ./cmd/seed
	./$(SEED_BINARY)

run: build
	./$(BINARY_NAME)

test:
	$(GOTEST) ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME) $(SEED_BINARY)