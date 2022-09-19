# Fullcycle - Arquitetura Hexagonal
O objetivo da arquitetura hexagonal é encapsular a lógica, de maneira que nada externo acesse-a diretamente, então, o meio de um usuário acessar uma informação gerada pela camada de domínio é através de um serviço.

## Divindo a aplicação em camadas:
 - Interface (Port -> entities, domain...) - Toda logica de dominio "regra de negócios" em um unico lugar;
 - Serviço (Adapter -> cli, databases, webserver...) - Meio-campo entre a camada de dominio e requisições de informações.

## Pré-requisitos
- Golang: `https://go.dev/dl/`
- Docker: `https://www.docker.com/`

## Utilização em desenvolvimento local
O arquivo `Makefile` que existe na raiz do projeto, tem todos os comandos necessários mapeados.

- Para criar container:
  - `$ docker-compose up -d`
 - Para acessar o bash do container:
   - `$ docker exec -it app-product bash` ou `$ make docker-bash`
