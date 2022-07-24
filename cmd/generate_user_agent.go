/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/programzheng/games/internal/service"
	"github.com/programzheng/games/pkg/helper"
	"github.com/spf13/cobra"
)

// generateUserAgentCmd represents the generateUserAgent command
var generateUserAgentCmd = &cobra.Command{
	Use:   "generateUserAgent",
	Short: "generateUserAgent ${agent_id} ${user_id} ${third_party_id}",
	Long:  `generateUserAgent for third party service`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("agent code is required")
		}
		if len(args) == 1 {
			log.Fatal("user id is required")
		}
		if len(args) == 2 {
			log.Fatal("third party id is required")
		}
		agentCode := args[0]
		userID := helper.ConvertToInt(args[1])
		thirdPartyID := args[2]

		agent, err := service.GetAgentByCode(agentCode)
		if err != nil {
			log.Fatalf("service.GetAgentByCode error: %v", err)
		}

		err = service.GenerateUserAgent(int(agent.ID), userID, thirdPartyID)
		if err != nil {
			log.Fatalf("service.GenerateUserAgent error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateUserAgentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateUserAgentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateUserAgentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
