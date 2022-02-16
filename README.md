Тестовое задание GO

- Описать proto файл с сервисом из 3 методов: добавить пользователя, удалить пользователя, список пользователей
- Реализовать gRPC сервис на основе proto файла на Go
- Для хранения данных использовать PostgreSQL на запрос получения списка пользователей данные будут кешироваться в redis на минуту и брать из редис
- При добавлении пользователя делать лог в clickHouse
- Добавление логов в clickHouse делать через очередь RabbitMQ

ПО: Git, Docker, Postman

Клонировать проект:
git clone https://github.com/s3rzh/go-grpc-user-service.git

Перейти в директорию проекта.

В корне проекта, в терминале выполнить след. команды:
Пересобрать контейнеры:
sudo docker-compose up -d --build

Создать образ для миграций:
docker build -t migrator ./api/migrator

Запустить Postgres миграции:
docker run --rm --network host migrator -path=./migrations/ -database "postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable" up

Запустить Clickhouse миграции:
docker exec user_clickhouse bash -c "clickhouse-client -mn < /var/lib/clickhouse/migrations/1_init_schema.up.sql"

В Postman

Открыть New -> gRPC Request указать locahost:8080, импортировать user.proto, выбрать метод и сгенерироваться сообщение в зависимости от выбранного метода автоматически (Generate Example Message) или вручную, например:

для метода CreateUser:
{
    "age": 25,
    "email": "my@eml.com"
}

для метода GetUsers (пустое сообщение):
{}

для метода DeleteUser:
{
    "email": "my@eml.com"
}


- в качестве клиента использовал Postman (поэтому всегда возврашал nil в ответе ошибки).
- тк задание тестовое, пароли оставил в файле конфига.

Откат Postgres миграций:
docker run --rm --network host migrator -path=./migrations/ -database "postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable" down -all

Откат Clickhouse миграций:
docker exec user_clickhouse bash -c "clickhouse-client -mn < /var/lib/clickhouse/migrations/1_init_schema.down.sql"

Админ панель RabbitMQ
http://localhost:15674/
guest/guest

DBeaver для просмотра PostgreSQL и ClickHouse

Утилита redis-cli для просмотра кэша:
docker exec -it user_redis bash
redis-cli
GET user:list


дата: февраль 2022