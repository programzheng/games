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
