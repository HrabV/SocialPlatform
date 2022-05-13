# Сервіс для моніторингу здоров'я всіх сервісів Socia

# Технологічний стек
- Node.JS
- Express
- Axios
- MongoDB (`mongoose` npm library)
- Socket.IO
- React.JS
- Environment Variables (.env file) -> `dotenv` (`STAT_SERVICE_MONGO_DB_URL`)

## Кінцеві точки (наприклад)
- http://127.0.0.1:5000/healthcheck

## Як це має працювати
- Сервіс відправляє запит на /healthcheck і очікує 204 статус код
- (Випадок 1) Якщо статус код - 204 тоді сервіс працює
- (Випадок 2) Якщо статус код 404 то тоді сервіс не працює
