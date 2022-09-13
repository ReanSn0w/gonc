# GONC
Простой центр уведомлений основанный на каналах

### Пример использования
```go
// создаем сущьность для реакции на событие
subscriber := nc.NewSubscriber(func() {
	print("Событие произошло")
})

// подписываемся
nc.Default().Subscribe("event_name", subscriber)

// далеко-далеко, но все же в вашем приложении
// инициируем
nc.Default().Send("event_name", nil)
```