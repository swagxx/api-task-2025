# Quotes API

REST API для управления цитатами: добавление, получение, выбор случайной, удаление.

---

## Функционал
- Добавление цитат
- Получение всех цитат
- Получение случайной цитаты
- Получение цитат по автору 
- Удаление цитат по ID
---


## Стэк
- GO, без использования сторонних библиотек
- PostgreSQL
- Unit-тесты 
---


## Как запустить?

### Клонировать проект 

```
git clone https://github.com/swagxx/api-task-2025.git
cd api-task-2025
```

### Настройка базы данных 
```sql
CREATE DATABASE db_quotes;
```
P.S Для корректной работы БД, необходимо прописать
```sql
go get github.com/lib/pq
```
Это драйвер, для подключения к БД

### Настройка .env файла
Укажите свои данные в example.env и перенесите их в .env
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=db_quotes
```

### Запустите приложение
```
go run ./cmd/main.go
```

## Работа с API 

### Добавить цитату
```
curl.exe -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d "{"author": "Confucius", "quote": "Life is simple."}"
```

### Получить все цитаты
```
curl http://localhost:8080/quotes
```

### Получить случайную цитату
```
curl http://localhost:8080/quotes/random
```

### Получить цитату по автору 
```
curl http://localhost:8080/quotes?author=Confucius
```

### Удалить цитату по ID
```
curl -X DELETE http://localhost:8080/quotes/1
```

## Также присутствуют тесты Хендлера

## Ключевые требования соблюдены, СПАСИБО ЗА ВНИМАНИЕ