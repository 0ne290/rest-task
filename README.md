## Как запускать:
1. Создайте файл .env в корне проекта. Пример содержимого файла:
```console
LOG_LEVEL=info

AUTH_KEY=secret_key_for_decoding_jwt

LISTEN_ADDRESS=:8080
WRITE_TIMEOUT=15s
SERVER_NAME=homework-4

DB_HOST=postgres
DB_PORT=5432
DB_NAME=homework
DB_USER=test_user
DB_PASSWORD=securepassword
DB_SSL_MODE=disable
DB_POOL_MAX_CONNS=10
DB_POOL_MAX_CONN_LIFETIME=300s
DB_POOL_MAX_CONN_IDLE_TIME=150s
```
2. Выполните команду docker compose up.

## Приложение запустится по адресу http://localhost:8080. Спецификацию можно посмотреть по адресу http://localhost:8080/swagger.

## Все эндпоинты v1/tasks защищены Bearer-авторизацией. Формат хэдера следующий: "Authorization: Bearer <JWT_Access_Token>".
## Т. е. чтобы работать с тасками, сперва надо получить Access Token от [сервиса аутентификации](https://github.com/0ne290/grpc-auth). С тасками может взаимодействовать только тот юзер, который их создал - иначе 401.
