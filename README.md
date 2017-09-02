# Nymeria

### Description

Nymeria is an automatic and reliable deployment measure system, providing basic mechanisms for small teams to analyze their deployment pipeline data.

### Origins

Nymeria is named after Arya Stark's direwolf in Game of Thrones.

## How to develop

Nymeria is developed using the Revel web framework. Revel is based on the Go Programming Language. Also, the Gorm ORM is used to store the application data.

You can find further instructions on how to install and configure the Go Development Environment in xxxxxx. It's necessary to install the revel cli, by issuing the following command:
* go get github.com/revel/revel

In order to develop Nymeria locally, the following dependencies are necessary:
* GCC
* Stdlib
* Sqlite3

##### Nymeria uses the Go Dep package management ( link ), so all dependecies are included in the source code. If you need to upgrade any of them, just run:
* dep ensure -update dependecy_name.

##### In order to run Nymeria locally, you can use:
* revel run github.com/rafaelbenvenuti/Nymeria

Nymeria will be listening on all interfaces at port 9000.

## How to prepare Nymeria for deployment

Nymeria's build process for deployment is based on Docker containers, so you'll need a Docker runtime engine that supports multi-stage builds. In order to install Docker, you may follow the instructions provided at [Docker Community Edition Documentation]().

To build the Docker image, run the following command: Docker build . -t <tag>

Nymeria's Dockerfile will create a new Docker container image that is

## How to deploy Nymeria

You may run docker run

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Nymeria main configuration file
        nymeria.conf  Override configuration file with user defined values
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        models/       Database models directory
        views/        Templates directory
    tests/            Test suites
