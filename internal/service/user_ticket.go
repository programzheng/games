package service

import (
	"fmt"
	"strings"

	"github.com/programzheng/games/internal/model"
	"github.com/programzheng/games/internal/repository"
	"github.com/programzheng/games/pkg/helper"
)

func GetUserTicketsByUserIDs(userIDs []int) ([]model.UserTicket, error) {
	uIDs := helper.ConvertIntSliceToStringSlice(userIDs)

	userTicketsWhere := fmt.Sprintf("WHERE user_id IN (%s) ", strings.Trim(strings.Join(uIDs, ","), "[]"))
	userTickets, err := repository.GetUserTickets("*", userTicketsWhere)
	if err != nil {
		return nil, err
	}

	return userTickets, nil
}

func GetUserTicketsAndTicketsByUserIDs(userIDs []int) ([]model.UserTicket, []model.Ticket, error) {
	uIDs := helper.ConvertIntSliceToStringSlice(userIDs)

	userTicketsWhere := fmt.Sprintf("WHERE user_tickets.user_id IN (%s) ", strings.Trim(strings.Join(uIDs, ","), "[]"))
	userTickets, tickets, err := repository.GetUserTicketsAndTickets("user_tickets.*, tickets.id, tickets.name", userTicketsWhere)
	if err != nil {
		return nil, nil, err
	}

	return userTickets, tickets, nil
}
