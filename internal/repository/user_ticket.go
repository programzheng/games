package repository

import (
	"fmt"

	"github.com/programzheng/games/internal/model"
)

const userTicketTableName = "user_tickets"

func GetUserTickets(columns string, wheres string) ([]model.UserTicket, error) {
	sql := getAllSyntax(userTicketTableName, columns, wheres)

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

	var userTickets []model.UserTicket
	for rows.Next() {
		var userTicket model.UserTicket
		if columns == "*" {
			err = rows.Scan(
				&userTicket.ID,
				&userTicket.UserID,
				&userTicket.TicketID,
				&userTicket.Code,
				&userTicket.CreatedAt,
				&userTicket.UpdatedAt,
				&userTicket.DeletedAt,
			)
		} else {
			for _, col := range rowsColumns {
				switch col {
				case "id":
					err = rows.Scan(
						&userTicket.ID,
					)
				case "user_id":
					err = rows.Scan(
						&userTicket.UserID,
					)
				case "ticket_id":
					err = rows.Scan(
						&userTicket.TicketID,
					)
				case "code":
					err = rows.Scan(
						&userTicket.Code,
					)
				case "created_at":
					err = rows.Scan(
						&userTicket.CreatedAt,
					)
				case "updated_at":
					err = rows.Scan(
						&userTicket.UpdatedAt,
					)
				case "deleted_at":
					err = rows.Scan(
						&userTicket.DeletedAt,
					)
				}
			}
		}

		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return nil, err
		}

		userTickets = append(userTickets, userTicket)
	}

	return userTickets, nil
}

func GetUserTicketsAndTickets(columns string, wheres string) ([]model.UserTicket, []model.Ticket, error) {
	sql := getAllWithJoinSyntax(userTicketTableName, columns, wheres, "JOIN tickets ON user_tickets.ticket_id = tickets.id")

	rows, err := DB.Query(sql)
	if err != nil {
		fmt.Printf("Query failed,err:%v\n", err)
		return nil, nil, err
	}
	defer rows.Close()

	var userTickets []model.UserTicket
	var tickets []model.Ticket
	for rows.Next() {
		var userTicket model.UserTicket
		var ticket model.Ticket
		err = rows.Scan(
			&userTicket.ID,
			&userTicket.UserID,
			&userTicket.TicketID,
			&userTicket.Code,
			&userTicket.CreatedAt,
			&userTicket.UpdatedAt,
			&userTicket.DeletedAt,
			&ticket.ID,
			&ticket.Name,
		)

		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return nil, nil, err
		}
		userTickets = append(userTickets, userTicket)
		tickets = append(tickets, ticket)
	}

	return userTickets, tickets, nil
}

func UpdateUserTickets(updates string, wheres string) (int, error) {
	sql := updateSyntax(userTicketTableName, updates, wheres)

	res, err := DB.Exec(sql)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
