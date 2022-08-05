package repository

import (
	"fmt"

	"github.com/programzheng/games/internal/model"
)

const ticketTableName = "tickets"
const userTicketsTableName = "user_tickets"

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

func GetTicket(columns string, wheres string) (*model.Ticket, error) {
	sql := getFirstSyntax(ticketTableName, columns, wheres)

	rows, err := DB.Query(sql)
	if err != nil {
		fmt.Printf("Query failed,err:%v\n", err)
		return nil, err
	}
	defer rows.Close()
	rowsColumns, err := rows.Columns()
	if err != nil {
		fmt.Printf("Query failed,err:%v\n", err)
		return nil, err
	}

	var ticket model.Ticket
	for rows.Next() {
		for _, col := range rowsColumns {
			switch col {
			case "id":
				err = rows.Scan(
					&ticket.ID,
				)
			case "name":
				err = rows.Scan(
					&ticket.Name,
				)
			case "created_at":
				err = rows.Scan(
					&ticket.CreatedAt,
				)
			case "updated_at":
				err = rows.Scan(
					&ticket.UpdatedAt,
				)
			case "deleted_at":
				err = rows.Scan(
					&ticket.DeletedAt,
				)
			}

			if err != nil {
				fmt.Printf("Scan failed,err:%v\n", err)
				return nil, err
			}
		}
	}

	return &ticket, nil
}

func GetTickets(columns string, wheres string) ([]model.Ticket, error) {
	sql := getAllSyntax(ticketTableName, columns, wheres)

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
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
			&ticket.DeletedAt,
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
	sql := insertSyntax(userTicketsTableName, columns, values)

	_, err := DB.Exec(sql)

	return err
}
