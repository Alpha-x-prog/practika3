# Docker Swarm - Полезные команды

## Управление стеком

### Посмотреть статус сервисов
```bash
docker stack services app
```

### Посмотреть детальный статус задач
```bash
docker stack ps app
```

### Посмотреть логи
```bash
# Логи сервиса
docker service logs app_service-a

# Только последние 100 строк
docker service logs --tail 100 app_nginx

# Follow режим (как tail -f)
docker service logs -f app_service-b
```

### Посмотреть настройки сервиса
```bash
docker service inspect app_nginx
```

### Масштабирование сервиса
```bash
# Увеличить до 3 реплик
docker service scale app_service-a=3

# Уменьшить до 1 реплики
docker service scale app_service-a=1
```

## Обновление стека

### После изменения stack.yaml
```bash
docker stack deploy -c stack.yaml app
```

## Удаление

### Удалить стек
```bash
docker stack rm app
```

### Остановить весь swarm
```bash
docker swarm leave --force
```

## Полезная информация

### Посмотреть все сервисы в swarm
```bash
docker service ls
```

### Посмотреть все контейнеры в стеке
```bash
docker stack ps app
```

### Список узлов
```bash
docker node ls
```

## Тестирование

### Проверка endpoints
```bash
curl http://localhost/a/
curl http://localhost/b/
curl http://localhost/healthz
```

### Проверить что request'ы распределяются между репликами
```bash
# Можно смотреть логи разных реплик
docker service logs app_service-a
# Каждый запрос может пойти на разную реплику благодаря load balancing
```

