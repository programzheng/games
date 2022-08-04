package service

import (
	"log"
	"testing"
)

func TestAssignRandomIssuedTicketToUser(t *testing.T) {
	userID := 1
	result, err := AssignRandomIssuedTicketToUser(userID)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("TestAssignRandomIssuedTicketToUser success result:%v", result)
}
