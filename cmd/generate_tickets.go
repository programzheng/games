/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/programzheng/games/internal/service"
	"github.com/spf13/cobra"
)

// generateTicketsCmd represents the generateTickets command
var generateTicketsCmd = &cobra.Command{
	Use:   "generateTickets",
	Short: "generate tickets by name to database",
	Long:  `generateTickets ${name1},${name2},${name3}......`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("at least a ticket name")
			return
		}
		ticketNames := strings.Split(args[0], ",")
		service.GenerateTickets(ticketNames)
		fmt.Println("generateTickets command success")
	},
}

func init() {
	rootCmd.AddCommand(generateTicketsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateTicketsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateTicketsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
