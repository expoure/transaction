# Título

## ordem de prioridade:
    - Testes + validações de negócio  transaction
        - se vai ser + ou - o amount etc
        - a quantidade de casas decimais
    - Swagger
    - Adicionar collection do insomnia no repo
    - Criar auth com jwt
    - Criar interface pro producer kafka/qualquer outro

### Tecnologias usadas:
- Docker
- Golang
- Kafka
- Postgres
- Krakend API-Gateway
- sqlc + pgx/v5

## Subindo o ambiente:

Crie a rede para o docker:
```bash
docker network create internal-net
```

Inicie o ambiente com:
```bash
docker-compose up -d
```
ou:
```bash
docker compose up -d
```

### Escolhas técnicas:
- Microservices: arquitetura escolhida pela escabilidade, modularidade, elasticidade, tolerância a falhas e confiabilidade.

- Krakend API-Gateway: direcionará as <em>requests</em> para o microservice adequado. 
- Event-Driven: ...

## Execução dos testes:

### Testes de integração com db:
```bash
make run-db-tests
```

# Arquitetura

![image info](./assets/arch.png)
