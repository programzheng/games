/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/programzheng/games/internal/service"
	"github.com/spf13/cobra"
)

// getThirdPartyUserCmd represents the getThirdPartyUser command
var getThirdPartyUserCmd = &cobra.Command{
	Use:   "getThirdPartyUser",
	Short: "getThirdPart8523bb6dfc45f5898531e90d5e03a074yUser ${code} ${third_party_id}",
	Long:  `getThirdPartyUser by code and third party id`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("agent code is required")
			return
		}
		code := args[0]
		if len(args) == 1 {
			fmt.Println("third party id is required")
			return
		}
		thirdPartyID := args[1]

		parameters := service.GetThirdPartyUserParameters{
			AgentCode:    code,
			ThirdPartyID: thirdPartyID,
		}
		user, err := service.GetThirdPartyUser(&parameters)
		if err != nil {
			panic(fmt.Sprintf("service.GenerateThirdPartyUser error: %v", err))
		}

		fmt.Printf("generateThirdPartyUser success\nuser: %v\n", user)
	},
}

func init() {
	rootCmd.AddCommand(getThirdPartyUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getThirdPartyUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getThirdPartyUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
