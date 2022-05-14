# Сервіс для файлів

## Стек технологій

- Golang 1.18
- Minio (Object Storage)
- GoFiber (golang lib for rest api)

## Змінні оточення

- `SS_SERVER_ADDR` - адреса сервера (порт)
- `SS_MINIO_ENDPOINT` - адреса сервера Minio
- `SS_MINIO_ACCESS_KEY` - логін для Minio
- `SS_MINIO_SECRET_KEY` - секретний ключ для Minio
- `SS_MINIO_USE_SSL` - чи використовувати ssl для Minio
- `SS_SERVER_HEADER` - заголовок сервера
- `SS_IS_DEVELOPMENT` - режим сервіс ("dev" або "prod")

## Як запустити ?

```bash
make run
```