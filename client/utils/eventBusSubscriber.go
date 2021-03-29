package utils

import "fmt"

type EventBusSubscriber struct {
	topic  string
	eb     *EventBus
	dataCh DataChannel
	doneCh chan bool
	Name   string //for debugging
}

func NewEventBusSubscriber(topic string, eb *EventBus) *EventBusSubscriber {
	ebs := &EventBusSubscriber{
		topic:  topic,
		eb:     eb,
		doneCh: make(chan bool),
		dataCh: make(chan DataEvent),
	}
	return ebs
}

func (s *EventBusSubscriber) Start(fcn func(d DataEvent)) {

	s.eb.Subscribe(s.topic, s.dataCh)

	go func() {
		for {
			select {
			case d := <-s.dataCh:
				//@todo confirm the topics match? (should be safe)
				fcn(d)
			case <-s.doneCh:
				if s.Name != "" {
					fmt.Println(s.Name, ": done.")
				}
				return
			}
		}
	}()
}

func (s *EventBusSubscriber) Stop() {
	s.doneCh <- true
	s.eb.Unsubscribe(s.topic, s.dataCh)
}
