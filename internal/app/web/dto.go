package web

type BrokerEvent struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}
