# Nymeria

## Description

Nymeria is an automatic and reliable deployment measure system, providing basic mechanisms for small teams to analyze their deployment pipeline data.

## Origins

Obvously, Nymeria is named after Arya Stark's direwolf in Game of Thrones. Although Nymeria has left Arya under unfortunate circunstances, this Nymeria is differente, and it will prove to be a valuable asset in environments where other references to GoT are necessary. Also, Nymeria is trustworthy and a great companion to your deployments.

## Technologies

The following technologies powers Nymeria:
* Go Programming Language
* Gorm Object Relationship Model
* Revel Web Framework
* Docker Containers
* Sqlite3

### Development

To help develop Nymeria, the following dependencies must be locally configured:
* Go Language Environment: [Go Language Install](https://golang.org/doc/install)
* Gorm ORM [Gorm ORM](http://jinzhu.me/gorm)
* Revel Web Framework [Revel Web Framework](https://https://revel.github.io)
* Sqlite3 [Sqlite3 Database](https://www.sqlite.org)
* Docker [Docker Community Edition Documentation](https://docs.docker.com/engine/installation).

In order to facilitate Nymeria's development, a Builder's Dockerfile is included in this repository, so if don't want to manually configure all those requirements, you may just simply execute `docker build -f Build-Dockerfile -t nymeria-dev .` inside this repository directory to generate a Docker Container Image with all the dependencies included, with vim batteries.

This docker container image shall be used to develop Nymeria in the following way:
1. Start the development docker container with `docker run --net host --rm -ti nymeria-dev`
2. Edit Nymeria code with Vim as you see fit.
3. Start Nymeria using `revel run github.com/rafaelbenvenuti/nymeria`
4. Wait almost 30 seconds for Revel and Database setup.
5. Enjoy

### Deployment
Nymeria's Default Dockerfile will create a Docker container image based on Alpine that is properly built for deployment.
This Dockerfile will use multi-stage builds to ensure that the final container image is small and self-contained, using around 32MB of disk space.
Althought Nymeria only supports sqlite3 at this moment, changing the database driver is very easy and all this careful thought about the final artifact generated for deployment will be important when deploying on orchestration systems like Swarm, Mesos and Kubernetes.

To run the deployment version, it is required to build nymeria using the default Dockerfile with `docker build -f Dockerfile -t nymeria .`.
After that, a new Docker image will be available for you to push into any Docker compatible repository. To run this version, you can use almost the same command for development, just need to change the image tag, like the following: `docker run --name nymeria --net host -ti nymeria`.
The above version of the command creates a container with the name `nymeria`, which you can manage with docker's start, stop and rm commands.

### Interacting


## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Nymeria main configuration file
        nymeria.conf  Override configuration file with user defined values
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
                deploys.go  Deploys controller
        models/       Database models directory
                deploy.go       Deploy model
        views/        Templates directory - Not used for APIs
    tests/            Test suites

## Todo
* Develop unit tests
* Monitoring with Prometheus
* Develop a systemd docker wrapper
* Configure a time series database backend
* Models for entities like accountable, components, status, etc...
