# Título

## ordem de prioridade:
    - Adicionar balance no account + mapper do Balance
    - Criar transaction microservice
    - Adicionar kafka
    - Testes

### Tecnologias usadas:
- Docker
- Golang
- Kafka
- Postgres
- Keycloak
- Kong API-Gateway
- sqlc

## Subindo o ambiente:

**Explicar como pegar o  key do realm no keycloak!!!**

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
- Para autenticação foi utilizado o keycloak. Toda `account` criada no sistema será espelhada no keycloak. Assim, pode-se obter tokens, gerenciar sessões, permissões etc. Neste projeto, apenas fluxo de autenticação será utilizado.

- Microservices: arquitetura escolhida pela escabilidade, modularidade, elasticidade, tolerância a falhas e confiabilidade.

- Kong API-Gateway: direcionará as <em>requests</em> para o microservice adequado. Possui uma camada de segurança para validação do token do usuário + rate limit (por ip) para criação de `accounts`

- Event-Driven: ...

# Arquitetura

![image info](./assets/arch.png)
