# gosolve-recruitment-task
Repo including code for the GoSolve recruitment task.

## Setup Project

### Requirements

1. Clone this repo

```bash
git clone https://github.com/kevinsantana/gosolve-recruitment-task.git
```

2. Copy [.env.example](.env.example) to `.env` and export them with

```bash
make envvars
```

### Run
1. Run api

```bash
make run
```

2. Either export [gosolve.postman_collection](docs/postman/gosolve.postman_collection.json) or make a curl request

```bash
curl --location 'localhost:5060/api/v1/1150/search'
```
