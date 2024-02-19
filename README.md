# GoYandexTask2.0
Project for Yandex Go course
## Как это работает
[![1.jpg](https://i.postimg.cc/YSkqgqx2/1.jpg)](https://postimg.cc/21cDNYJg)
- Frontend статический, index.html - главная страница, реализована навигация
- Backend сервер запускается консольно и выполняет функционал оператора.
- Хранение данных производится внутри json файлов, лежащих в корне(expressions, operations, power).

## Запросы

Сервер поднимается на http://localhost:8080

- Отправить выражение POST http://localhost:8080/send

- Получить полный список всех выражений GET http://localhost:8080/expressions

- Получить операции и их время GET http://localhost:8080/operations

- Задать время для операций POST http://localhost:8080/operations

- Получить список вычислительных ресурсов GET http://localhost:8080/power

- Заглушка для OPTIONS запросов в угоду CORS OPTION http://localhost:8080/options

## Запуск

Для запуска сервера следует использовать
```
go run Backend/main.go
```
## Уровень реализации
~~Сомнительный~~

### Сделано:
- Относительно неплохой UI
- Настроены запросы с клиента на оркестратор
- Настроена перезапись json при помощи запросов
### Не сделано:
- Времени на написание демона не хватило, поэтому вычисления не производятся.
- СУБД не используется
- Docker не используется
- Нет валидации выражений
