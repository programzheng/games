package repository

import (
	"fmt"

	"github.com/programzheng/games/internal/model"
)

const userTableName = "users"
const userAgentTableName = "user_agents"

func CreateUser(columns string, values string) error {
	sql := insertSyntax(userTableName, columns, values)

	_, err := DB.Exec(sql)

	return err
}

func GetUsers(columns string, wheres string) ([]model.User, error) {
	sql := getAllSyntax(userTableName, columns, wheres)

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

	var users []model.User
	for rows.Next() {
		var user model.User
		if columns == "*" {
			err = rows.Scan(
				&user.ID,
				&user.CreatedAt,
				&user.UpdatedAt,
				&user.DeletedAt,
			)
		} else {
			for _, col := range rowsColumns {
				switch col {
				case "id":
					err = rows.Scan(
						&user.ID,
					)
				case "created_at":
					err = rows.Scan(
						&user.CreatedAt,
					)
				case "updated_at":
					err = rows.Scan(
						&user.UpdatedAt,
					)
				case "deleted_at":
					err = rows.Scan(
						&user.DeletedAt,
					)
				}
			}
		}

		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserLastInsertID() (*int, error) {
	sql := getLastInsertID(userTableName)

	rows, err := DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int
	for rows.Next() {
		err := rows.Scan(
			&id,
		)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return nil, err
		}
	}

	return &id, nil
}

func GetUserByID(columns string, id uint) (*model.User, error) {
	sql := getFirstSyntax(userTableName, columns, fmt.Sprintf("WHERE id = %d", id))

	rows, err := DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user model.User
	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return nil, err
		}
	}

	return &user, nil
}

func CreateUserAgent(columns string, values string) error {
	sql := insertSyntax(userAgentTableName, columns, values)

	_, err := DB.Exec(sql)

	return err
}

func GetUserAgents(columns string, wheres string) ([]model.UserAgent, error) {
	sql := getAllSyntax(userAgentTableName, columns, wheres)

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

	var userAgents []model.UserAgent
	for rows.Next() {
		var userAgent model.UserAgent
		if columns == "*" {
			err = rows.Scan(
				&userAgent.ID,
				&userAgent.AgentID,
				&userAgent.UserID,
				&userAgent.ThirdPartyID,
				&userAgent.CreatedAt,
				&userAgent.UpdatedAt,
				&userAgent.DeletedAt,
			)
		} else {
			for _, col := range rowsColumns {
				switch col {
				case "id":
					err = rows.Scan(
						&userAgent.ID,
					)
				case "agent_id":
					err = rows.Scan(
						&userAgent.AgentID,
					)
				case "user_id":
					err = rows.Scan(
						&userAgent.UserID,
					)
				case "third_party_id":
					err = rows.Scan(
						&userAgent.ThirdPartyID,
					)
				case "created_at":
					err = rows.Scan(
						&userAgent.CreatedAt,
					)
				case "updated_at":
					err = rows.Scan(
						&userAgent.UpdatedAt,
					)
				case "deleted_at":
					err = rows.Scan(
						&userAgent.DeletedAt,
					)
				}
			}
		}

		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return nil, err
		}

		userAgents = append(userAgents, userAgent)
	}

	return userAgents, nil
}

func GetUserAgentByAgentIDAndThirdPartyID(agentID uint, thirdPartyID string) (*model.UserAgent, error) {
	sql := getFirstSyntax(userAgentTableName, "*", fmt.Sprintf(
		"WHERE agent_id = '%d' AND third_party_id = '%s'", agentID, thirdPartyID,
	))

	rows, err := DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userAgent model.UserAgent
	for rows.Next() {
		err = rows.Scan(
			&userAgent.ID,
			&userAgent.AgentID,
			&userAgent.UserID,
			&userAgent.ThirdPartyID,
			&userAgent.CreatedAt,
			&userAgent.UpdatedAt,
			&userAgent.DeletedAt,
		)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return nil, err
		}
	}

	return &userAgent, nil
}

func GetUserAgentByThirdPartyID(thirdPartyID string) (*model.UserAgent, error) {
	sql := getFirstSyntax(userAgentTableName, "*", "WHERE third_party_id = '"+thirdPartyID+"'")

	rows, err := DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userAgent model.UserAgent
	for rows.Next() {
		err = rows.Scan(
			&userAgent.ID,
			&userAgent.UserID,
			&userAgent.ThirdPartyID,
			&userAgent.CreatedAt,
			&userAgent.UpdatedAt,
			&userAgent.DeletedAt,
		)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return nil, err
		}
	}

	return &userAgent, nil
}
