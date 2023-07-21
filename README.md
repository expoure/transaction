# Pismo - Rotina de Transações

## Index

- [Executando](#executando)
  * [Makefile](#makefile)
  * [Prefiro eu mesmo fazer](#prefiro-eu-mesmo-fazer)
- [Testes](#testes)
  * [Integração com Banco Postgres](#integração-com-db)
  * [Unitários](#unitários)
- [Arquitetura](#arquitetura)
- [Tecnologias usadas](#tecnologias-usadas)
- [Escolhas técnicas](#escolhas-técnicas)
    * [Microservices](#microservice)
    * [API-GATEWAY](#api-gateway)
    * [Dinheiro do tipo...INTEGER!](#dinheiro-do-tipointeger)
    * [Event-Driven](#event-driven)

## Executando;

### Makefile

O comando abaixo irá subir todos os serviços _docker_ necessários:
```bash
make up-all
```
### Prefiro eu mesmo fazer

Caso prefira subir o ambiente sem o _Makefile_ acima:
```bash
docker network create internal-net
docker compose up -d --scale kafkaui=0
```
Dessa forma, a rede docker que o projeto necessita será criada e o container do [kafka-ui](https://github.com/provectus/kafka-ui) será ignorado, uma vez que está presente no projeto apenas para **acompanhamento** do fluxo de eventos.

## Testes:

### Integração com db:
```bash
make run-db-tests
```

### Unitários:
```bash
make TODO
```

## Arquitetura

![image info](./assets/arch.png)

## Tecnologias usadas
- Docker
- Golang (1.20)
- Kafka
- Postgres
- Krakend API-Gateway
- sqlc + pgx/v5
- gomock(uber)

## Escolhas técnicas

### Microservice:

Arquitetura escolhida pela escabilidade, modularidade, elasticidade, tolerância a falhas, testabilidade e confiabilidade.

### API-GATEWAY

Uma Api-Gateway entrega muitas vantagens, neste pequeno projeto serve para direcionar as _requests_ para o _microservice_ adequado.

### Dinheiro do tipo...INTEGER!

Como estamos lidando com real e centavos (ou quaisquer que sejam os equivalentes), e eles geralmente são representados por um número decimal, pode parecer óbvio usar float ou decimal, pois eles são projetados para representar números que incluem casas decimais. No entanto, se você entender um pouco sobre como float funciona no nível do hardware, verá por que essa não é a melhor abordagem.

![image info](./assets/golang_float.png)

Você pode rodar o exemplo [aqui](https://go.dev/play/p/IrhUSV1CZGC) e também ler mais informações neste ótimo [artigo](https://blog.codeminer42.com/be-cool-dont-use-float-double-for-storing-monetary-values)

### Event-Driven
Trabalhar com _microservices_ pode ser muito complexo dependendo do domínio da aplicação. Um dos grandes problemas desta arquitetura são as chamadas síncronas entre serviços, que podem gerar lentidão no sistema como um todo ou folharem devido a falhas de rede. Event-Driven é descrito por Mark Richards e Neal Ford em [Fundamentals of Software Architecture: An Engineering Approach](https://www.goodreads.com/book/show/44144493-fundamentals-of-software-architecture) como uma `arquitetura`. Nesta arquitetura, cada ação gera um evento e este será usado por outra ação que também irá gerar um evento e assim por diante.</p>

Devido a esta característica, _microservices_ "casam" bem como uma arquitetura baseada em eventos, pois os erros de rede são drasticamente diminuídos e tudo acontece de forma assíncrona.
</p>