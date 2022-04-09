/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package get

import (
	"fmt"
	"github.com/spf13/cobra"
)

// podsCmd represents the pods command
var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Pods is called\n")
		str, err := cmd.Flags().GetString("n")
		if err != nil {
			fmt.Printf("error happens \n")
			return
		}
		fmt.Printf("params n %s\n", str)
	},
}

func init() {
	getCmd.AddCommand(podsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	podsCmd.PersistentFlags().String("n", "", "namespace")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// podsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
