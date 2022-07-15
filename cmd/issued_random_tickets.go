/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/programzheng/games/internal/service"
	"github.com/programzheng/games/pkg/helper"
	"github.com/spf13/cobra"
)

// issuedRandomTicketsCmd represents the issuedRandomTickets command
var issuedRandomTicketsCmd = &cobra.Command{
	Use:   "issuedRandomTickets",
	Short: "issued random tickets by count",
	Long:  `issuedRandomTickets ${count}`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("count is required")
			return
		}
		count := helper.ConvertToInt(args[0])
		if count <= 0 {
			fmt.Println("count is require greater than zero")
			return
		}
		err := service.IssuedRandomTickets(count)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("issuedRandomTickets success")
	},
}

func init() {
	rootCmd.AddCommand(issuedRandomTicketsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// issuedRandomTicketsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// issuedRandomTicketsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
