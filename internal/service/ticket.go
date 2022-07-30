package service

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/programzheng/games/internal/repository"
	"github.com/programzheng/games/pkg/helper"
)

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

func IssuedRandomTickets(count int) error {
	tickets, err := repository.GetTickets("*")
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
