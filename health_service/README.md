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
- Якщо сервіс не відповідає більше ніж 10 секунд то можна рахувати його неробочим
