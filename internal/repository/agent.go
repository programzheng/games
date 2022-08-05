package repository

import (
	"fmt"
	"log"

	"github.com/programzheng/games/internal/model"
)

const agentTableName = "agents"

func CreateAgent(columns string, values string) error {
	sql := insertSyntax(agentTableName, columns, values)

	_, err := DB.Exec(sql)

	return err
}

func GetAgentByCode(code string) (*model.Agent, error) {
	sql := getFirstSyntax(agentTableName, "*", "WHERE code = '"+code+"'")

	rows, err := DB.Query(sql)
	if err != nil {
		log.Fatalf("FindAgentByCode error: %v", err)
	}
	defer rows.Close()

	var agent model.Agent
	for rows.Next() {
		err = rows.Scan(
			&agent.ID,
			&agent.Code,
			&agent.Name,
			&agent.CreatedAt,
			&agent.UpdatedAt,
			&agent.DeletedAt,
		)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return nil, err
		}
	}

	return &agent, nil
}
