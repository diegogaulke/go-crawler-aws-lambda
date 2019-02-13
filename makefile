GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
DEPLOYMENT_ZIP=deployment.zip
BINARY_NAME=go-crawler-aws-lambda
BINARY_UNIX=$(BINARY_NAME)_unix

all: clean build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
	zip $(DEPLOYMENT_ZIP) $(BINARY_NAME)
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(DEPLOYMENT_ZIP)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
test: 
	$(GOTEST) -v ./...
# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
	zip $(DEPLOYMENT_ZIP) $(BINARY_UNIX)