# Zero Hash challenge

As part of the interviewing process to work at Zero Hash I did this little project.

## What I did

### Web server structure

I converted all game generation to Golang so I could test it properly, after that I created the webserver responsible for running such application, it has the following endpoints:

HTTP Method   | URI Path     | Description |
------------- | -------------|-------------|
GET           | /            | Shows a static page |
GET           | /metrics     | Shows metrics regarding the application (I included information about how many games each team won)|
POST           | /test       | Runs the game, you have to specify how many turns you want 

### Project structure

```
project
│
└───pkg
│   │
│   └───metrics (responsible for creating metrics and give info about it)
│   |    │   metrics.go
│   |    │   metrics_http.go
|   |___phttp (defines the interface the other http handlers must follow)
|   |    |   phttp.go   
│   |___runner (responsible for running the game, logging about it and generating metrics)
│   |    │   runner.go
│   |    │   runner_http.go
└───cmd
│   │
│   └───pictionary(defines what the application is gonna be like and serves it on :3000)
│   |    │   main.go

```

## Development

### Build and run locally
```
make build-local
make run-local
```
The project was built and tested locally using `go1.16`.

### Build and run on docker
```
DOCKER_REGISTRY=<docker-registry-name> RELEASE_VERSION=<image-tag> make build-docker
DOCKER_REGISTRY=<docker-registry-name> RELEASE_VERSION=<image-tag> make run-docker
```
For example:
```
DOCKER_REGISTRY="viniciusramosdefaria" RELEASE_VERSION="0.0.1" make build-docker
DOCKER_REGISTRY="viniciusramosdefaria" RELEASE_VERSION="0.0.1" make run-docker
```
### Test
```
NUMBER_OF_ROUNDS=<number-of-rounds> make test
```
For example:
```
NUMBER_OF_ROUNDS=7 make test
```

## Future work

- Breaking down methods for the runner package
- Writing tests
- Use templates to show info on the front page


