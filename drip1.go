package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type DripOneService struct {
	db *sql.DB
}

func NewDripOneService(db *sql.DB) DripOneService {
	return DripOneService{db: db}
}

type DripOne struct {
	quoteNumber int
	firstName   string
	lastName    string
}

type DripOneDispatcher struct {
	quoteNumber int
	number      int
	timeToSend  time.Time
}

func (d DripOneService) GetFactory() ([]DripOne, error) {
	var dripOne []DripOne
	for i := 0; i < 10; i++ {
		dripOne = append(dripOne, DripOne{quoteNumber: i, firstName: "Jane", lastName: fmt.Sprintf("Doe %d", i)})
	}
	return dripOne, nil
}

func (d DripOneService) GetSchedule() []DripSchedule {
	var dripItems []DripSchedule
	for i := 0; i < 5; i++ {
		var timeToSend time.Time
		if i == 0 {
			timeToSend = time.Now()
		} else {
			timeToSend = time.Now().Add(time.Duration(i*24) * time.Hour)
		}

		dripItems = append(dripItems, DripSchedule{Number: i + 1, TimeToSend: timeToSend})
	}
	return dripItems
}

func (d DripOneService) InsertCacheRecord(record DripOne, timeToSend time.Time, number int) error {
	fmt.Printf("DRIP 1: Inserting cache record for %s %s, time to send: %s, number: %d \n", record.firstName, record.lastName, timeToSend, number)
	return nil
}

func (d DripOneService) UpdateCacheRecord(record DripOneDispatcher, delivered bool, cancelled bool) error {
	fmt.Printf("DRIP 1: Updating cache record for quote number %d, number: %d, time to send: %s, delivered: %t, cancelled: %t \n", record.quoteNumber, record.number, record.timeToSend, delivered, cancelled)
	return nil
}

func (d DripOneService) GetQueue() []DripOneDispatcher {
	var queue []DripOneDispatcher
	for i := 0; i < 10; i++ {
		queue = append(queue, DripOneDispatcher{quoteNumber: i, number: i + 1, timeToSend: time.Now()})
	}
	return queue
}

func (d DripOneService) SendEmail(r DripOneDispatcher) error {
	bodyContent := r.GetHtmlContent()
	fmt.Printf("DRIP 1: Sending email for quote number %d, number: %d, time to send: %s - %s \n", r.quoteNumber, r.number, r.timeToSend, bodyContent)
	return errors.New("error sending email")
}

func (d DripOneService) HandleStopCondition(r DripOneDispatcher) error {
	fmt.Printf("DRIP 1: Handle stop condition for quote number %d, number: %d, time to send: %s \n", r.quoteNumber, r.number, r.timeToSend)
	return nil
}

func (d DripOneDispatcher) GetHtmlContent() string {
	return "This is the HTML content"
}

func (d DripOneDispatcher) StopConditionMet() bool {
	return d.number == 5
}
