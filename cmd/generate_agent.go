/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/programzheng/games/internal/service"
	"github.com/spf13/cobra"
)

// generateAgentCmd represents the generateAgent command
var generateAgentCmd = &cobra.Command{
	Use:   "generateAgent",
	Short: "generateAgent ${name} ${code}",
	Long:  `generateAgent by name and code`,
	Run: func(cmd *cobra.Command, args []string) {
		name := ""
		var code *string

		if len(args) == 0 {
			fmt.Println("name is required")
			return
		}
		name = args[0]
		if len(args) == 2 {
			code = &args[1]
		}
		err := service.GenerateAgent(name, code)
		if err != nil {
			fmt.Printf("generateAgent command error: %v", err)
			return
		}
		fmt.Println("generateAgent command success")
	},
}

func init() {
	rootCmd.AddCommand(generateAgentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateAgentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateAgentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
