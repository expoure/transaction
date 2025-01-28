# Pismo - Transaction Routine

## Index

- [Running](#running)
  * [Environment Variables](#environment-variables)
  * [Makefile](#makefile)
  * [I Prefer Doing It Myself](#i-prefer-doing-it-myself)
  * [Using the APIs](#using-the-apis)
- [Tests](#tests)
  * [Postgres DB Integration](#integrated)
  * [Unit Tests](#unit)
- [Architecture](#architecture)
- [Technologies Used](#technologies-used)
- [Technical Choices](#technical-choices)
    * [Hexagonal](#hexagonal)
    * [Microservices](#microservices)
    * [API Gateway](#api-gateway)
    * [Money as...INTEGER!](#money-as-integer)
    * [Event-Driven](#event-driven)

## Running

### Environment Variables

Ensure that the environment variables used, especially those related to ports, do not conflict with any other services on your machine.  
The files can be found in:  
`services/*/.env`.


### Makefile

The command below will start all the necessary _Docker_ services:  
```bash
make up-all
```

### I Prefer Doing It Myself

If you prefer to set up the environment without using the _Makefile_ above:
```bash
docker network create internal-net
docker compose up -d --scale kafkaui=0 --scale db-test=0
```
This way, the Docker network required by the project will be created, and the containers for [kafka-ui](https://github.com/provectus/kafka-ui) and the test database will be ignored, as they are included in the project solely for **monitoring** the event flow and running integration tests.

### Using the APIs

You can import the _Insomnia_ collection located in this repository or use the commands below directly in the terminal.  
**curl must be installed!**

#### Creating an _Account_

```bash
curl --request POST \
  --url http://localhost:8080/v1/accounts \
  --header 'Content-Type: application/json' \
  --data '{
	"documentNumber": "12348378734"
}'
```

#### Get _account_ by _id_

```bash
curl --request GET \
  --url http://localhost:8080/v1/accounts/you-uuid-here
```

#### Creating _transaction_ for an _account_

```bash
curl --request POST \
  --url http://localhost:8080/v1/transactions \
  --header 'Content-Type: application/json' \
  --data '{
	"accountId": "you-uuid-here",
	"operationTypeId": 1,
	"amount": -5.00
}'
```


## Tests:

### Integrated:
```bash
make run-db-tests
```

### Unit:
```bash
make run-unit-tests
```

## Architecture

![image info](./assets/arch.png)

## Technologies used
- Docker
- Golang (1.20)
- Kafka
- Postgres
- Krakend API-Gateway
- sqlc + pgx/v5
- gomock(uber)

## Technical Choices

### Hexagonal
This architectural design helps with the testability and flexibility of our application, allowing greater focus on the domain without concerns about external components.

### Microservices

This architecture was chosen for its scalability, modularity, elasticity, fault tolerance, testability, and reliability.

### API-Gateway

An API-Gateway provides many advantages. In this small project, it serves to route _requests_ to the appropriate _microservice_.

### Money as INTEGER!

Since we are dealing with currency (real and cents or their equivalents), typically represented by a decimal number, it might seem obvious to use float or decimal because they are designed to represent numbers with decimal places. However, if you understand how float works at the hardware level, you'll see why this is not the best approach.

![image info](./assets/golang_float.png)

You can run the example [here](https://go.dev/play/p/IrhUSV1CZGC) and read more about it in this excellent [article](https://blog.codeminer42.com/be-cool-dont-use-float-double-for-storing-monetary-values).

### Event-Driven
Working with _microservices_ can be very complex depending on the application domain. One of the major challenges of this architecture is synchronous calls between services, which can slow down the system as a whole or fail due to network issues. Event-Driven is described by Mark Richards and Neal Ford in [Fundamentals of Software Architecture: An Engineering Approach](https://www.goodreads.com/book/show/44144493-fundamentals-of-software-architecture) as an `architecture`. In this architecture, each action generates an event, which will be used by another action that also generates an event, and so on.

Because of this characteristic, _microservices_ align well with an event-driven architecture, as network errors are drastically reduced, and everything happens asynchronously.
