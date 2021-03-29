package utils

import (
	"fmt"
	"sync"
)

type DataEvent struct {
	Data  interface{}
	Topic string
}

type DataChannel chan DataEvent
type DataChannelSlice []DataChannel

type EventBus struct {
	subscribers map[string]DataChannelSlice
	m           sync.RWMutex
}

// NewEventBus
func NewEventBus() *EventBus {
	e := &EventBus{
		subscribers: map[string]DataChannelSlice{},
	}

	return e
}

// Subscribe to a given topic
func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.m.Lock()
	defer eb.m.Unlock()

	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
}

// Unsubscribe from a given topic
func (eb *EventBus) Unsubscribe(topic string, ch DataChannel) {

	eb.m.Lock()
	defer eb.m.Unlock()

	if subs, found := eb.subscribers[topic]; found {
		for i, sub := range subs {
			if sub == ch {
				eb.subscribers[topic] = append(subs[:i], subs[i+1:]...)
				return
			}
		}
	} else {
		fmt.Println("Subscriber not found:", topic, ch)
	}
}

// Publish ...
func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.m.Lock()
	defer eb.m.Unlock()

	if subs, found := eb.subscribers[topic]; found {

		// make our own slice of data channels
		subscribers := append(DataChannelSlice{}, subs...)
		go func(d DataEvent, dcs DataChannelSlice) {
			for _, ch := range dcs {
				ch <- d
			}
		}(DataEvent{Topic: topic, Data: data}, subscribers)
	}
}
