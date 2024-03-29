package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/programzheng/games/internal/model"
	"github.com/programzheng/games/internal/repository"
	"github.com/programzheng/games/pkg/helper"
)

type GenerateUserParameters struct {
	Name     string
	Account  string
	Password string
}

type GetThirdPartyUserParameters struct {
	AgentCode    string
	ThirdPartyID string
}

type GenerateThirdPartyUserParameters struct {
	AgentCode    string
	ThirdPartyID string
}

func GenerateUser(parameters *GenerateUserParameters) error {
	if parameters == nil {
		log.Fatalf("service.GenerateUser parameters is nil")
	}

	err := repository.CreateUser("id", "(null)")
	if err != nil {
		return err
	}

	return nil
}

func GenerateThirdPartyUser(parameters *GenerateThirdPartyUserParameters) error {
	err := checkParameters(parameters, "service.GenerateThirdPartyUser parameters is nil")
	if err != nil {
		return err
	}
	agent, err := repository.GetAgentByCode(parameters.AgentCode)
	if err != nil {
		return err
	}
	if agent == nil {
		return fmt.Errorf("not found agent by code: %s", parameters.AgentCode)
	}

	err = repository.CreateUser("id", "(null)")
	if err != nil {
		return err
	}
	userID, err := repository.GetUserLastInsertID()
	if err != nil {
		return err
	}

	err = repository.CreateUserAgent("agent_id, user_id, third_party_id", fmt.Sprintf(
		"(%d, %d, '%s')", agent.ID, *userID, parameters.ThirdPartyID,
	))
	if err != nil {
		return err
	}

	return nil
}

func GenerateUserAgent(agentID int, userID int, thirdPartyID string) error {
	err := repository.CreateUserAgent("agent_id, user_id, third_party_id",
		fmt.Sprintf("(%d, %d, '%s')",
			agentID,
			userID,
			thirdPartyID,
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func GenerateThirdPartyUserAgent(agentID int, thirdPartyID string) error {
	err := repository.CreateUserAgent("agent_id, third_party_id",
		fmt.Sprintf("(%d, '%s')",
			agentID,
			thirdPartyID,
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func GetThirdPartyUser(parameters *GetThirdPartyUserParameters) (*model.User, error) {
	err := checkParameters(parameters, "service.GetThirdPartyUser parameters is nil")
	if err != nil {
		return nil, err
	}

	agent, err := repository.GetAgentByCode(parameters.AgentCode)
	if err != nil {
		return nil, err
	}
	if agent.ID == 0 {
		return nil, fmt.Errorf("not found agent by code: %s", parameters.AgentCode)
	}

	userAgent, err := repository.GetUserAgentByAgentIDAndThirdPartyID(
		agent.ID, parameters.ThirdPartyID,
	)
	if err != nil {
		return nil, err
	}
	if userAgent.ID == 0 {
		return nil, nil
	}

	user, err := repository.GetUserByID("*", userAgent.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByAgentCode(agentCode string) ([]model.User, error) {
	agent, err := repository.GetAgentByCode(agentCode)
	if err != nil {
		return nil, err
	}
	if agent.ID == 0 {
		return nil, fmt.Errorf("not found agent by code: %s", agentCode)
	}

	userAgents, err := repository.GetUserAgents(
		"user_id",
		fmt.Sprintf("WHERE agent_id = %d", agent.ID),
	)
	if err != nil {
		return nil, err
	}

	userIDs := []int{}
	for _, userAgent := range userAgents {
		userIDs = append(userIDs, int(userAgent.UserID))
	}
	uIDs := helper.ConvertIntSliceToStringSlice(userIDs)

	users, err := repository.GetUsers(
		"id",
		fmt.Sprintf("WHERE id IN (%s)", strings.Trim(strings.Join(uIDs, ","), "[]")),
	)
	if err != nil {
		return nil, err
	}

	return users, nil
}
