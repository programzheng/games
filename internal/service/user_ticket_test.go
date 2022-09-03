package service

import (
	"log"
	"testing"
)

func TestGetUserTicketsByUserIDs(t *testing.T) {
	result, err := GetUserTicketsByUserIDs([]int{1})
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("TestGetUserTicketsByUserIDs success result:%v", result)
}
