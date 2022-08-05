package service

type PlayAssignTicketForThirdPartyUserResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func PlayAssignTicketForThirdPartyUser(agentCode string, thirdPartyID string) (*PlayAssignTicketForThirdPartyUserResponse, error) {
	agent, err := GetAgentByCode(agentCode)
	if err != nil {
		return nil, err
	}

	user, err := GetThirdPartyUser(&GetThirdPartyUserParameters{
		AgentCode:    agent.Code,
		ThirdPartyID: thirdPartyID,
	})
	if err != nil {
		return nil, err
	}

	if user == nil {
		err := GenerateThirdPartyUser(&GenerateThirdPartyUserParameters{
			AgentCode:    agent.Code,
			ThirdPartyID: thirdPartyID,
		})
		if err != nil {
			return nil, err
		}
		user, err = GetThirdPartyUser(&GetThirdPartyUserParameters{
			AgentCode:    agent.Code,
			ThirdPartyID: thirdPartyID,
		})
		if err != nil {
			return nil, err
		}
	}

	assignRandomIssuedTicketToUserResponse, err := AssignRandomIssuedTicketToUser(int(user.ID))
	if err != nil {
		return nil, err
	}

	return &PlayAssignTicketForThirdPartyUserResponse{
		Code: assignRandomIssuedTicketToUserResponse.Code,
		Name: assignRandomIssuedTicketToUserResponse.Name,
	}, nil
}
