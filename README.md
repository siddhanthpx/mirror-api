# Fastest Debian Mirror API

Just a simple API for finding the fastest download mirror provided [here](www.debian.org/mirror/list).

Using regex I first fetched all the mirror links from debian then for every link I ran a goroutine ( lightweight green thread ), the first goroutine that finishes would be the fastest so we terminate the process.

## Installation

Clone the repo and build

```bash
go build -o mirror-api .
./mirror-api
```

## Usage

Which starts the service on `localhost:8080/fastest-mirror`
