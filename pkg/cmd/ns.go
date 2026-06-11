/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/cannon/pkg/core/mutate"
)

// nsCmd represents the ns command
var nsCmd = &cobra.Command{
	Use:   "ns",
	Short: "This command switch the current namespace to given value",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		namespace := "default" 
		if len(args) > 0  {
			namespace = args[0]
		} 

		err := mutate.SetNamespace(namespace)
		if err != nil {
			fmt.Println("Error occured: ", err);
		}
	},
}

func init() {
	rootCmd.AddCommand(nsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
