# Título

### Tecnologias usadas:
- Docker
- Golang
- Kafka
- Postgres
- Krakend API-Gateway
- sqlc + pgx/v5
- gomock(uber)

## Subindo o ambiente:

```bash
make up-all
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

### Testes unitários:
```bash
make TODO
```

# Arquitetura

![image info](./assets/arch.png)
