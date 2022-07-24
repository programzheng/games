/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/programzheng/games/server/grpc_server"
	"github.com/spf13/cobra"
)

// grpcServerCmd represents the grpcServer command
var grpcServerCmd = &cobra.Command{
	Use:   "grpcServer",
	Short: "run a grpc server",
	Long:  `run a grpc server`,
	Run: func(cmd *cobra.Command, args []string) {
		grpc_server.Run()
	},
}

func init() {
	rootCmd.AddCommand(grpcServerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcServerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcServerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
