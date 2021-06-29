# Fastest Debian Mirror API

Just a simple API for finding the fastest download mirror provided <a href="https://www.debian.org/mirror/list" target="_blank">here</a>.


Using regex I first fetched all the mirror links from debian then for every link I ran a goroutine ( lightweight green thread ), the first goroutine that finishes would be the fastest so we terminate the process.

## Installation

If you have go on your machine just clone the repo and build :

```bash
go build -o mirror-api .
./mirror-api
```

or you can use the Dockerfile to create an image to spawn the container from :

```bash
docker build -t mirror-api:latest .
docker run -d --name mirror-api -p 8080:8080 mirror-api:latest
```

## Usage

Which starts the service on `localhost:8080/fastest-mirror`
