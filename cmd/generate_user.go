/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/programzheng/games/internal/service"
	"github.com/spf13/cobra"
)

// generateUserCmd represents the generateUser command
var generateUserCmd = &cobra.Command{
	Use:   "generateUser",
	Short: "generateUser",
	Long:  `generateUser`,
	Run: func(cmd *cobra.Command, args []string) {
		err := service.GenerateUser(&service.GenerateUserParameters{})
		if err != nil {
			log.Fatalf("generateAgent command error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
