package main

import (
	"fmt"
	"time"
)

type DripSchedule struct {
	Number     int
	TimeToSend time.Time
}

type DripDispatcherItem interface {
	GetHtmlContent() string
	StopConditionMet() bool
}

type DripServiceInterface[FactoryStruct any, DispatcherStruct DripDispatcherItem] interface {
	GetFactory() ([]FactoryStruct, error)
	GetSchedule() []DripSchedule
	InsertCacheRecord(record FactoryStruct, timeToSend time.Time, number int) error
	UpdateCacheRecord(record DispatcherStruct, delivered, cancelled bool) error
	GetQueue() []DispatcherStruct
	SendEmail(item DispatcherStruct) error
	HandleStopCondition(item DispatcherStruct) error
}

func main() {
	var dripOneServiceS DripServiceInterface[DripOne, DripOneDispatcher] = NewDripOneService(nil)

	scheduleDripEmails(dripOneServiceS)
	fmt.Println("Sending drip emails")
	sendDripEmails(dripOneServiceS)
}

func scheduleDripEmails[T any, V DripDispatcherItem](service DripServiceInterface[T, V]) {
	scheduledDripEmails, err := service.GetFactory()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, dripEmail := range scheduledDripEmails {
		fmt.Println(dripEmail)
		for _, dripItem := range service.GetSchedule() {
			service.InsertCacheRecord(dripEmail, dripItem.TimeToSend, dripItem.Number)
		}
	}
}

func sendDripEmails[T any, V DripDispatcherItem](service DripServiceInterface[T, V]) {
	queue := service.GetQueue()
	for _, item := range queue {
		if !item.StopConditionMet() {
			err := service.SendEmail(item)
			if err != nil {
				fmt.Println("Error sending email for item", item, err)
			}
		}

		err := service.UpdateCacheRecord(item, true, false)
		if err != nil {
			fmt.Println("Error updating cache record for item", item, err)
		}

		if item.StopConditionMet() {
			fmt.Println("Stop condition met")
			err = service.HandleStopCondition(item)
			if err != nil {
				fmt.Println("Error handling stop condition for item", item, err)
			}
		}
	}
}
