package service

import (
	"fmt"
	"math/rand"
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

func IssuedRandomTickets(count uint) error {
	tickets, err := repository.GetTickets("*")
	if err != nil {
		return err
	}

	for i := 0; i < int(count); i++ {
		unix := time.Now().UnixMicro()
		rand.Seed(unix)

		ticket := tickets[rand.Intn(len(tickets))]
		sid := helper.ConvertToString(ticket.ID)
		secret := helper.CreateMD5(int(ticket.ID) + int(unix))

		colums := "ticket_id,code"
		values := fmt.Sprintf("(%s,'%s')", sid, secret)
		err := repository.CreateUserTickets(colums, values)
		if err != nil {
			return err
		}
	}

	return nil
}
