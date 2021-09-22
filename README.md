# Volume-weighted average price (VWAMP) calculation engine

  - [About the project](#about-the-project)
    - [Design](#design)
  - [Getting started](#getting-started)
    - [Layout](#layout)
  - [Notes](#notes)


## About the project
The goal of this project is to create a real-time VWAP (volume-weighted average price) calculation engine which pull coinbase websocket feed to stream in trade executions and update the VWAP for each trading pair as updates become available.

### Design

This project follows the [Clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) convention.

## Getting started

### Prerequisites
- installed [Golang 1.17](https://golang.org/) 
- or run using [Docker](https://www.docker.com/)

### Start application
Run using Go
```sh
make run // to start application
make test // to run all the tests
```
Run using Docker
```sh
make docker-build
make docker-run
make docker-test
```


### Layout

```tree
├── .gitignore
├── CHANGELOG.md
├── Makefile
├── README.md
├── service
│   └── vamp.service.go
├── release
│   ├── template-admin.yaml
│   └── template-controller.yaml
├── test
│   ├── README.md
│   └── test_make.sh
└── external
    ├── notifier
    └── provider
```

A brief description of the layout:

* `.gitignore` varies per project, but all projects need to ignore `bin` directory.
* `CHANGELOG.md` contains auto-generated changelog information.
* `README.md` is a detailed description of the project.
* `pkg` places most of project business logic.
* `test` holds all tests.
* `external` for all external services i.e. feed data provider and notifier.

## Notes

* Makefile **MUST NOT** change well-defined command semantics, see Makefile for details.
