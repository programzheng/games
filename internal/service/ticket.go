package service

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/programzheng/games/internal/model"
	"github.com/programzheng/games/internal/repository"
	"github.com/programzheng/games/pkg/helper"
)

type AssignRandomIssuedTicketToUserResult struct {
	Name string
	Code string
}

func GenerateTickets(names []string) {
	ticketNames := ""
	for index, name := range names {
		ticketNames += fmt.Sprintf("('%s')", name)
		if index < len(names)-1 {
			ticketNames += ", "
		}
	}

	if ticketNames == "" {
		panic("ticket names is empty")
	}

	err := repository.CreateTicketByName(ticketNames)
	if err != nil {
		panic(err)
	}

}

func GetTicketsByIDs(ids []int) ([]model.Ticket, error) {
	ticketsIDs := helper.ConvertIntSliceToStringSlice(ids)

	ticketsWhere := fmt.Sprintf("WHERE id IN (%s) ", strings.Trim(strings.Join(ticketsIDs, ","), "[]"))
	tickets, err := repository.GetTickets("*", ticketsWhere)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func AssignRandomIssuedTicketToUser(userID int) (*AssignRandomIssuedTicketToUserResult, error) {
	noOwnerUserTickets, err := repository.GetUserTickets("id, ticket_id, code", "WHERE user_id IS NULL")
	if err != nil {
		return nil, err
	}
	if len(noOwnerUserTickets) == 0 {
		return nil, fmt.Errorf("userTickets not found")
	}

	unixNano := time.Now().UnixNano()
	rand.Seed(unixNano)
	noOwnerUserTicket := noOwnerUserTickets[rand.Intn(len(noOwnerUserTickets))]

	updateCount, err := repository.UpdateUserTickets(
		fmt.Sprintf("user_id = %d", userID),
		fmt.Sprintf("WHERE id = %d", noOwnerUserTicket.ID),
	)
	if err != nil {
		return nil, err
	}
	if updateCount == 0 {
		return nil, errors.New("no updates")
	}

	ticket, err := repository.GetTicket(
		"name",
		fmt.Sprintf("WHERE id = %d", noOwnerUserTicket.TicketID),
	)
	if err != nil {
		return nil, err
	}

	return &AssignRandomIssuedTicketToUserResult{
		Name: ticket.Name,
		Code: noOwnerUserTicket.Code,
	}, nil
}

func IssuedRandomTickets(count int) error {
	tickets, err := repository.GetTickets("*", "")
	if err != nil {
		return err
	}

	//default a goroutine
	routineNum := 1
	if count/runtime.NumCPU() > 0 {
		routineNum = count / runtime.NumCPU()
	}
	issuedNumber := count / routineNum
	remainIssuedNumber := count % routineNum
	if remainIssuedNumber > 0 {
		routineNum++
	}

	done := make(chan struct{}, routineNum)
	defer close(done)

	rows := ""
	var mux sync.RWMutex
	for num := 0; num < routineNum; num++ {
		//last routine
		if remainIssuedNumber > 0 && num == (routineNum-1) {
			issuedNumber = remainIssuedNumber
		}
		go func(currentNum int, issuedNumber int, currentDone chan struct{}) {
			for i := 0; i < int(issuedNumber); i++ {
				unixNano := time.Now().UnixNano()
				rand.Seed(unixNano)

				ticket := tickets[rand.Intn(len(tickets))]
				sid := helper.ConvertToString(ticket.ID)
				secret := helper.CreateMD5(currentNum + i + rand.Intn(int(unixNano)))

				values := fmt.Sprintf("(%s,'%s'),", sid, secret)

				mux.Lock()
				rows += values
				mux.Unlock()
			}
			currentDone <- struct{}{}
		}(num, issuedNumber, done)
	}
	for i := 0; i < routineNum; i++ {
		<-done
	}
	rows = strings.TrimSuffix(rows, ",")
	columns := "ticket_id,code"
	err = repository.CreateUserTickets(columns, rows)
	if err != nil {
		return err
	}

	return nil
}
