# Работа с RebbitMQ из Go

## Быстрый старт:
1) Стартуем RebbitMQ (password/user:guest)
```bash
docker-compose up --build
```
2) Запускаем sender
3) Запускаем consumer
4) Отправляем Get запрос на sender
```
http://10.10.0.136:3000/send?msg=Respect
```

//TODO завернуть sender и consumer в docker и сделать общий docker-compose