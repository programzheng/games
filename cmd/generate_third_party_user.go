/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/programzheng/games/internal/service"
	"github.com/spf13/cobra"
)

// generateThirdPartyUserCmd represents the generateThirdPartyUser command
var generateThirdPartyUserCmd = &cobra.Command{
	Use:   "generateThirdPartyUser",
	Short: "generateThirdPartyUser ${code} ${third_party_id}",
	Long:  `generateThirdPartyUser by code and third party id`,
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

		parameters := service.GenerateThirdPartyUserParameters{
			AgentCode:    code,
			ThirdPartyID: thirdPartyID,
		}
		err := service.GenerateThirdPartyUser(&parameters)
		if err != nil {
			panic(fmt.Sprintf("service.GenerateThirdPartyUser error: %v", err))
		}

		fmt.Println("generateThirdPartyUser success")
	},
}

func init() {
	rootCmd.AddCommand(generateThirdPartyUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateThirdPartyUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateThirdPartyUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
