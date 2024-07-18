# Desafio Back-end Frete Rápido

### Instruções para execução local
Definir as variáveis de ambiente no arquivo .env
```env
DB_HOST=172.16.10.2
DB_PORT=3306
DB_USERNAME=
DB_PASSWORD=
DB_DATABASE=

FR_COTACAO_ENDPOINT=

TEST_CPNJ=
TEST_TOKEN=
TEST_PLATFORM_CODE=
TEST_ZIP_CODE=
```
Executar o build da imagem do container da aplicação
```shell
docker build -f Dockerfile . -t fr-cotacao
```
Para executar os containers foi disponibilizado o arquivo `docker-compose.yml`

```yaml
services:
  api:
    image: fr-cotacao
    container_name: quote_api-api
    restart: unless-stopped
    env_file:
      - .env
    depends_on:
      - db
    ports:
      - 8080:80
    networks:
      quote_net:
        ipv4_address: 172.16.10.3

  db:
    image: bitnami/mariadb:11.4
    container_name: quote_api-db
    expose:
      - 3306
    volumes:
      - mariadb_data:/bitnami/mariadb
    environment:
      MARIADB_USER: ${DB_USERNAME}
      MARIADB_PASSWORD: ${DB_PASSWORD}
      MARIADB_DATABASE: ${DB_DATABASE}
      ALLOW_EMPTY_PASSWORD: yes
      TZ: 'America/Sao_Paulo'
    networks:
      quote_net:
        ipv4_address: 172.16.10.2

volumes:
  mariadb_data:
    driver: local

networks:
  quote_net:
    driver: bridge
    ipam:
      config:
        - subnet: 172.16.10.0/28
```
Crie o schema do banco de dados
```shell
docker exec -i quote_api-db mysql -u root -v quotes < scripts/sql/schema.sql
```
Agora a API está pronta para ser utilizada

Endpoints

Documentação da API

`http://172.16.10.3/swagger/index.html`

`http://localhost:8080/swagger/index.html`

Cotação

`POST http://172.16.10.3/api/v1/quote`

`POST http://localhost:8080/api/v1/quote`

Metricas

`GET http://172.16.10.3/api/v1/metrics`

`GET http://localhost:8080/api/v1/metrics`

## Informações Adicionais

Para praticidade realizei deploy da API e uma imagem no docker hub

URL: `http://cotacao.ianbrito.com.br`

IMAGE: `https://hub.docker.com/r/ianbrito/fr-cotacao`

Documentação da API

`http://cotacao.ianbrito.com.br/swagger/index.html`

Cotação

`POST http://cotacao.ianbrito.com.br/api/v1/quote`

Metricas

`GET http://cotacao.ianbrito.com.br/api/v1/metrics`