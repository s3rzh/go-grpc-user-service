Тестовое задание GO

- Описать proto файл с сервисом из 3 методов: добавить пользователя, удалить пользователя, список пользователей
- Реализовать gRPC сервис на основе proto файла на Go
- Для хранения данных использовать PostgreSQL на запрос получения списка пользователей данные будут кешироваться в redis на минуту и брать из редис
- При добавлении пользователя делать лог в clickHouse
- Добавление логов в clickHouse делать через очередь RabbitMQ

ПО: Git, Docker, Postman

Клонировать проект:
git clone git@gitlab.com:dev-area/horse-exchange.git

Перейти в папку проекта и cоздать в корне файл .env

Добавить в него
DB_PASSWORD=qwerty

В корне проекта, выполнить след. команды:
docker-compose up -d --build horse-exchange-app

Создание образа для миграций
docker build -t migrator ./api/migrator

Запустить миграции
docker run --network host migrator -path=./migrations/ -database "postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable" up

В Postman

открыть grpc




Для отката миграций
docker run --network host migrator -path=./migrations/ -database "postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable" down -all