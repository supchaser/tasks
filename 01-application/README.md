# Cloud.ru Camp 2025 — Echo Server

Небольшое Go-приложение, которое на HTTP 8000 отдаёт:
- имя хоста,
- IP-адрес,
- значение переменной окружения `AUTHOR`.

## Функции

- `GET /` — HTML-страница с хостом, IP и AUTHOR  
- `GET /healthz` — проверка статуса (возвращает `ok` и HTTP 200)  
- Graceful shutdown на SIGINT/SIGTERM  

## Требования

- Go ≥ 1.24  
- GNU Make  
- Docker  

## Локальная сборка и запуск

1. Клонируйте репозиторий и перейдите в него:
   ```bash
   git clone https://github.com/supchaser/tasks.git
   cd tasks/01-application

2. Соберите бинарь:
   ```bash
   make build

3. Запустите:
   ```bash
   ./bin/echo-server

4. Откройте в браузере [localhost](http://localhost:8000)
