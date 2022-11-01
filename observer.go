/*
Observer is helpful when we want to notify about an action being completed;
instead of a children process constantly asking if the action has been done,
here it's the pattern process the one to trigger its observers
*/
package main

import (
	"fmt"
	"time"
)

type Topic interface {
	register(observer Observer)
	broadcast()
	UpdateAvailability(status bool)
}

type Observer interface {
	updateValue(string)
	getID() string
}

type Item struct {
	name        string
	observers   []Observer
	isAvailable bool
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i *Item) broadcast() {
	for _, v := range i.observers {
		v.updateValue(i.name)
	}
}

func (i *Item) UpdateAvailability(status bool) {
	i.isAvailable = status
	if i.isAvailable {
		i.broadcast()
	}
}

func NewTopic(name string) Topic {
	return &Item{
		name: name,
	}
}

type EmailNotification struct {
	id string
}

func (en *EmailNotification) getID() string {
	return en.id
}

func (en *EmailNotification) updateValue(value string) {
	fmt.Printf("Action triggered (email notification id: %s) by %s\n", en.id, value)
}

func NewObserver(id string) Observer {
	return &EmailNotification{
		id: id,
	}
}

func main() {
	i1 := NewTopic("book1")
	o1 := NewObserver("en1")
	i1.register(o1)
	time.Sleep(5 * time.Second)
	i1.UpdateAvailability(true)
}
