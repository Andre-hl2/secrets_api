# Secrets API

This is a Go REST API that implements the "Guesser Game" functionality.

This is a learning project, I'm using this to get myself comfortable with the Go language, understanding its syntax and how to properly setup and develop a project with it

## The Guesser Game

The Guesser games goes as it follows

- Someone creates a new secret, by sending a word to the service and not telling it to anyone 🤫

- Other folks can then try to guess what the secret is by sending guesses requests to the API 🤔

- Whoever guess the word correctly first will be stored as the correct guesser for that secret 🎉

## Running the service

In order to run this service you'll need Go installed. After this then run:

```
cd server
cp .env.example .env
go install
go run .
```

With this you should have the application running with the `memory` store at the address `http://0.0.0.0:8080`. You can then verify that it's running correctly by using:

```
$ curl http://0.0.0.0:8080/health     

{"status":"pass","healthy":true}
```

You can check 

### Env Vars

| Name | Description | Example |
| ------------ | ------------------------------------------------------- | ---------------------------------------- |
| SERVER_PORT  | Which port the application will run on                  | 8080                                     |
| DATA_SOURCE  | which store will the application use                    | must be one of [ `memory`, `postgres`]   |
| DATABASE_URL | if data source is set as `postgres`, provide the DB URL | `postgresql://USER:PW@127.0.0.1:5432/DB` |


## Read more about the project:

## [📂 server](./server/README.md)

## [📂 devops](./devops/README.md)
