CURRENT_GIT_GROUP 	:= github.com/ucanme/fastgo
CURRENT_GIT_REPO  	:= fastgo
COMMONENVVAR		= GOOS=linux GOARCH=amd64
BUILDENVVAR			= CGO_ENABLED=0
GOTESTNOCACHE		?= GOCACHE=off
DOCKER_HUB			:= harbor.cn
DOCKER_GROUP		:= group
DOCKER_PROJECT		:= fastgo
BIN_NAME			:= fastgo
JSONITERTAG			:= -tags=jsoniter
GITTAG 				?= $(shell git describe --abbrev=0 --tags 2>/dev/null || echo NO_TAG)
GITTAG_NO_V			:= $(subst v,,$(GITTAG))
GITCOMMITHASH		?= $(shell git rev-parse --short HEAD)
DATE_TIME			:= $(shell date '+%Y-%m-%dT%H:%M:%S')
COMMIT_COUNT		:= $(shell git rev-list --all --count)
DOCKER_VERSION		:= $(GITTAG_NO_V).$(COMMIT_COUNT).$(GITCOMMITHASH)
VERSIONTAG			:= -ldflags "-X main.BuildTime=$(DATE_TIME) -X main.BuildGitHash=$(GITCOMMITHASH) -X main.BuildGitTag=$(GITTAG) -X main.BuildGitCount=$(COMMIT_COUNT)"
# PACKAGES			:= `go list ./... | grep -v /vendor/`
# VETPACKAGES			:= `go list ./... | grep -v /vendor/ | grep -v /examples/`


.PHONY:  test clean pb all build

all: build

build:
	go build $(JSONITERTAG) -o bin/$(BIN_NAME) $(VERSIONTAG) .

linux_build:
	go build $(JSONITERTAG) -o bin/$(BIN_NAME) $(VERSIONTAG) .  `
# 	$(COMMONENVVAR) $(BUILDENVVAR) go build $(JSONITERTAG) -o bin/$(BIN_NAME)-linux $(VERSIONTAG) .

test:
	go test -v ./


docker:
	docker build -t $(DOCKER_HUB)/$(DOCKER_GROUP)/$(DOCKER_PROJECT):$(DOCKER_VERSION) -f Dockerfile .

docker-push:
	docker push $(DOCKER_HUB)/$(DOCKER_GROUP)/$(DOCKER_PROJECT):$(DOCKER_VERSION)

fmt:
	gofmt -w ${GOFILES}

fmt-check:
	@diff=$$(goimports -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

vet:
	go vet $(VETPACKAGES)

clean:
	@rm -rf bin

