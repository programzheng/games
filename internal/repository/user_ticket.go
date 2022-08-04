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

	var userTickets []model.UserTicket
	for rows.Next() {
		var userTicket model.UserTicket
		err = rows.Scan(
			&userTicket.ID,
			&userTicket.TicketID,
			&userTicket.Code,
		)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return nil, err
		}
		userTickets = append(userTickets, userTicket)
	}

	return userTickets, nil
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
