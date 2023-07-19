# Título

## ordem de prioridade:
    - Adicionar balance no account + mapper do Balance [x]
    - Criar transaction microservice
    - Adicionar kafka
    - Testes + validações de negócio
    - Swagger
    - Adicionar collection do insomnia

### Tecnologias usadas:
- Docker
- Golang
- Kafka
- Postgres
- Krakend API-Gateway
- sqlc

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

# Arquitetura

![image info](./assets/arch.png)
