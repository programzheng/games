package repository

import (
	"fmt"

	"github.com/programzheng/games/internal/model"
)

const ticketTableName = "tickets"
const userTicketsTableName = "user_tickets"

var DB = GetDB()

func CreateTicketByName(name string) error {
	sql := generateSyntax("INSERT INTO `%s` (name) VALUES %s", ticketTableName, name)
	_, err := DB.Exec(sql)

	return err
}

func CreateTicketsByName(names []string) error {
	sql := generateSyntax("INSERT INTO `%s` (name) VALUES ", ticketTableName)

	for key, name := range names {
		sql += "('" + name + "')"

		if key != (len(names) - 1) {
			sql += ","
		}
	}
	_, err := DB.Exec(sql)

	return err
}

func GetTickets(sel string) ([]model.Ticket, error) {
	sql := generateSyntax("SELECT %s FROM `%s`", sel, ticketTableName)

	rows, err := DB.Query(sql)
	if err != nil {
		fmt.Printf("Query failed,err:%v\n", err)
		return nil, err
	}
	defer rows.Close()

	var tickets []model.Ticket
	for rows.Next() {
		var ticket model.Ticket
		err = rows.Scan(
			&ticket.ID,
			&ticket.Name,
			&ticket.CreateAt,
			&ticket.UpdateAt,
			&ticket.DeleteAt,
		)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func CreateUserTickets(columns string, values string) error {
	sql := generateSyntax("INSERT INTO `%s` (%s) VALUES %s", userTicketsTableName, columns, values)

	_, err := DB.Exec(sql)

	return err
}
