Тестовое задание GO

- Описать proto файл с сервисом из 3 методов: добавить пользователя, удалить пользователя, список пользователей
- Реализовать gRPC сервис на основе proto файла на Go
- Для хранения данных использовать PostgreSQL на запрос получения списка пользователей данные будут кешироваться в redis на минуту и брать из редис
- При добавлении пользователя делать лог в clickHouse
- Добавление логов в clickHouse делать через очередь RabbitMQ

ПО: Git, Docker, Postman

Клонировать проект:
git clone https://github.com/s3rzh/go-grpc-user-service.git

Перейти в папку проекта.

В корне проекта, выполнить след. команды:
docker-compose up -d --build app

Создание образа для миграций.
docker build -t migrator ./api/migrator

Запустить миграции.
docker run --network host migrator -path=./migrations/ -database "postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable" up

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



В качестве клиента использовал Postman (поэтому всегда возврашал nil в ответе ошибки)

P.S. тк задание тестовое, пароли оставил в файле конфига.



Для отката миграций
docker run --network host migrator -path=./migrations/ -database "postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable" down -all

админ панель RabbitMQ
http://localhost:15674/
guest/guest
