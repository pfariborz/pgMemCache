/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// SETCmd represents the SET command
var SETCmd = &cobra.Command{
	Use:   "SET",
	Short: "SET key to store string value",
	Long: `SET key to store string value, if key is already
	present the value will be overwritten. For example: 
	pgMemCache SET key "Hello World!"`,
	Run: func(cmd *cobra.Command, args []string) {
		storeKeyPair(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(SETCmd)
}
