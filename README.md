# email-service

## Project structure and explanation

`service.go`

- interface `MailService` contains the SendMail definition

`failover_handler_service.go`

- on instantiation takes in multiple instances of `MailService`.
- contains one function which handles failing over between services (in case a service is down).

`mailgun_service.go` and `sendgrid_service.go`

- Implement the `MailService` for each provider.

`model.go`

- contains a Config struct which contains two other structs with fields for each service (Mailgun/SendGrid)
- contains our Mail struct with basic fields (To, Subject, Message)

`config.yml`

- contains api keys and sender info per service

`controller.go`

- Handles swagger endpoint and POST endpoint (which calls the FailOverHandlerService and handles the response accordingly)

`main.go`

- Handles decoding of `config.yml` content into `Config` struct
- Create new instances of the services and inject them into the `FailOverHandlerService`
- Creates instance of the `Controller` and calls `HttpRoutes()`

## How to run

##### Note: You first need to setup Sendgrid and Mailgun accounts, after that you can add the API keys to `config.yml`

### Docker

`docker-compose up -d`

### Locally

- `cd src`
- `go run cmd/main.go`

## How to build

### Docker

`docker-compose build`

### Locally

- `cd src`
- `go build cmd/main.go`

## How to test

- `cd src`
- `go test ./...`

## Swagger

After running the application, swagger docs can be found here: http://localhost:8080/swagger/index.html

> `curl -X POST "http://localhost:8080/api" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"message\": \"👋\", \"subject\": \"Test 🍻\", \"to\": \"xyz@gmail.com\"}"`
