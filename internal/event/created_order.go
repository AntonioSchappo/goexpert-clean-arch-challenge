package event

import "time"

type CreatedOrder struct {
	Name    string
	Payload interface{}
}

func NewCreatedOrder() *CreatedOrder {
	return &CreatedOrder{
		Name: "CreatedOrder",
	}
}

func (e *CreatedOrder) GetName() string {
	return e.Name
}

func (e *CreatedOrder) GetPayload() interface{} {
	return e.Payload
}

func (e *CreatedOrder) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *CreatedOrder) GetDateTime() time.Time {
	return time.Now()
}
