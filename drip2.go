package main

import (
	"fmt"
	"time"
)

type DripTwoService struct {
}

type DripTwo struct {
	policyNumber int
	firstName    string
	lastName     string
}

func (d DripTwoService) GetFactory() ([]DripTwo, error) {
	var dripTwo []DripTwo
	for i := 0; i < 10; i++ {
		dripTwo = append(dripTwo, DripTwo{policyNumber: i, firstName: "John", lastName: fmt.Sprintf("Doe %d", i)})
	}
	return dripTwo, nil
}

func (d DripTwoService) GetSchedule() []DripSchedule {
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

func (d DripTwoService) InsertCacheRecord(record DripTwo, timeToSend time.Time, number int) {
	fmt.Printf("DRIP 2: Inserting cache record for %s %s, time to send: %s, number: %d \n", record.firstName, record.lastName, timeToSend, number)
}
