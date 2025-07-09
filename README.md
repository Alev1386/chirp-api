# Chirp API

**Chirp API** — учебный RESTful backend-сервис на Go для управления параметрами ЛЧМ-сигналов.

## Запуск проекта

1. Клонируйте репозиторий:

```bash
git clone git@github.com:Alev1386/chirp-api.git
cd chirp-api
````

2. Инициализируйте и скачайте зависимости (если нужно):

```bash
go mod tidy
```

3. Запустите сервер:

```bash
go run main.go
```

4. Проверьте работу:

Откройте в браузере или выполните curl:

```bash
curl http://localhost:8080/health
```

Должен быть ответ:

```json
{"status":."ok".}
```
