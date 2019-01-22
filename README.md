# tax_calculator
This project contains 2 APIs to be used by front-end engineers to develop an application that store and display tax amounts.

## Getting Started
These instruction will guide how to run this project.

### Prerequisites
```
- Go Language
- docker
```

### Installing
```shell
$ git clone https://github.com/MartinToruan/tax_calculator.git
$ cd tax_calculator
$ docker-compose up -d
```

## Running the tests
Run the unit test by running command below:
```shell
$ go test `go list ./... | grep -v 'dockerfile\|svc'` --cover
```
