package nc

import "sync"

// Функция создает новый центр для отправки уведомлений
func NewNotificationCenter() *NotificationsCenter {
	return &NotificationsCenter{
		wg:        sync.WaitGroup{},
		registred: make(map[string][]Subscriber),
	}
}

// Структура представляет из себя центр уведомлений
type NotificationsCenter struct {
	wg        sync.WaitGroup
	registred map[string][]Subscriber
}

// Метод производит добаление нового подпищика для события в центре уведомлений
func (nc *NotificationsCenter) Subscribe(name string, subscriber Subscriber) {
	nc.operation(func() {
		subscribers := nc.registred[name]
		if len(subscribers) == 0 {
			subscribers = []Subscriber{}
		}

		subscribers = append(subscribers, subscriber)
		nc.registred[name] = subscribers
	})
}

// Метод производит удаление конкретного подпищика.
func (nc *NotificationsCenter) Unsubscribe(name string, subscriber Subscriber) {
	nc.operation(func() {
		subscribers := nc.registred[name]
		if len(subscribers) == 0 {
			return
		}

		updated := make([]Subscriber, 0, len(subscribers)-1)

		for _, s := range subscribers {
			if s != subscriber {
				updated = append(updated, s)
			} else {
				s.Done()
			}
		}

		nc.registred[name] = updated
	})
}

// Метод служит для отправки уведомления о событии
func (nc *NotificationsCenter) Send(name string, payload interface{}) {
	subscribers := nc.registred[name]
	if len(subscribers) == 0 {
		return
	}

	for _, subscriber := range subscribers {
		if subscriber == nil {
			continue
		}

		subscriber.Chan() <- payload
	}
}

func (nc *NotificationsCenter) operation(operation func()) {
	nc.wg.Wait()
	nc.wg.Add(1)

	operation()

	nc.wg.Done()
}
