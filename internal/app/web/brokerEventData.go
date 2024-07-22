package web

type BrokerEventData struct {
	Channel chan BrokerEvent
}

func NewBrokerEventData() *BrokerEventData {
	return &BrokerEventData{make(chan BrokerEvent)}
}
