/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "short description for sub command",
	Long:  "long description for sub command",
	Run: func(cmd *cobra.Command, args []string) {
		subAction()
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}

func subAction() {
	fmt.Println("sub action called")
}
