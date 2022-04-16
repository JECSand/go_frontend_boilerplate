# golang-frontend-boilerplate

boilerplate for golang frontend development

## Setup

1- Ensure a non-root user is setup with password-less sudo powers

2- Clone this repo
```bash
$ git clone https://gitlab.synercloud.io/general/golang-frontend-boilerplate.git
$ cd golang-frontend-boilerplate
```

### Installation

1- Ubuntu Linux:
- Run Setup and Install Scripts
```bash
$ . ./setup.sh
```

### Configure

1- copy conf.json.example
```bash
$ cp conf.json.example conf.json
$ vi conf.json
```

2- open conf.json in a text editor and edit json settings as needed

## Test
```bash
$ go test ./src/spa_app/cmd/app
```

## Run
### Development
```bash
$ go build ./src/spa_app/cmd/app
$ ./app
$ curl http://localhost:8080
```

### Production

1- Start App
```bash
$ sh start.sh
$ curl http://localhost:8080
```

2- Stop App
```bash
$ sh stop.sh
```
