package nc

type Subscriber interface {
	Chan() chan<- interface{}
	Done()
}

// Функция создает нового подписчика на уведомление
// о факте наступления события, при этом payload передается в функцию
// В payload могут быть любые данные в том числе и nil, по этой причини при возникновении события
// следует аккуратно проверять тип payload сообщения
func NewSubsriber(completion func(interface{})) Subscriber {
	ps := payloadSubscriber{
		done:       make(chan bool),
		ch:         make(chan interface{}),
		completion: completion,
	}

	go ps.wait()
	return &ps
}

type payloadSubscriber struct {
	done       chan bool
	ch         chan interface{}
	completion func(interface{})
}

func (ps *payloadSubscriber) Done() {
	ps.done <- true
}

func (ps *payloadSubscriber) Chan() chan<- interface{} {
	return ps.ch
}

func (ps *payloadSubscriber) wait() {
	done := false

	for {
		select {
		case <-ps.done:
			done = true
			ps = nil
		case v := <-ps.ch:
			ps.completion(v)
		}

		if done {
			break
		}
	}
}
