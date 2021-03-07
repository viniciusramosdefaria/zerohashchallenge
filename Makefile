REGISTRY = $(DOCKER_REGISTRY)
RELEASE = $(RELEASE_VERSION)
ROUNDS = $(NUMBER_OF_ROUNDS)

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build

#Binary names
BINARY_NAME=pictionary

#Main Files
CMD_PATH=./cmd/pictionary/main.go

# Docker
DOCKERCMD=docker
DOCKERBUILD=${DOCKERCMD} build
DOCKERPUSH=${DOCKERCMD} push
DOCKERRUN=${DOCKERCMD} run -p 3000:3000

build: build-docker

build-local:
	GOCACHE='/tmp/.cache' CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GOBUILD} -o ./bin/${BINARY_NAME} -v ${CMD_PATH}

run-local:
	./bin/${BINARY_NAME}

build-docker:
	${DOCKERBUILD} -t "${REGISTRY}/${BINARY_NAME}:${RELEASE}" .

publish: build-docker
	${DOCKERPUSH} "${REGISTRY}/${BINARY_NAME}:${RELEASE}"

run-docker:
	${DOCKERRUN} "${REGISTRY}/${BINARY_NAME}:${RELEASE}"
test:
	curl --location --request POST 'http://localhost:3000/test' --header 'Content-Type: application/json' --data-raw '{"metadata": {"rounds":  "${ROUNDS}"}}'
