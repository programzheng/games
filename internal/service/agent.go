package service

import (
	"fmt"
	"time"

	"github.com/programzheng/games/internal/model"
	"github.com/programzheng/games/internal/repository"
	"github.com/programzheng/games/pkg/helper"
)

func GetAgentByCode(code string) (*model.Agent, error) {
	agent, err := repository.GetAgentByCode(code)
	if err != nil {
		return nil, err
	}
	return agent, nil
}

func GenerateAgent(name string, code *string) error {
	if code == nil {
		unixNano := time.Now().UnixNano()
		t := helper.CreateMD5(name + helper.ConvertToString(unixNano))
		code = &t
	}

	err := repository.CreateAgent("name, code", fmt.Sprintf("('%s', '%s')", name, *code))
	if err != nil {
		return err
	}

	return nil
}
