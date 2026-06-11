/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/cannon/pkg/lib"
	"github.com/cannon/pkg/core/query"

)

// listpodsCmd represents the listpods command
var listpodsCmd = &cobra.Command{
	Use:   "listpods",
	Short: "This command lists all the pods in the current namespace",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		var namespace string
		if len(args) > 0 {
			namespace = args[0]
		} else {
			namespace, _ = lib.GetCurrentNamespace();
		}

		kubeConfig, err := lib.GetKubeConfig()
		if err != nil {
			fmt.Println("Error occured: ", err);
		}

		if err = query.NewDynamicQuery().ListPods(kubeConfig, namespace); err != nil {
			fmt.Println("Error occured: ", err);
		}

	},
}

func init() {
	rootCmd.AddCommand(listpodsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listpodsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listpodsCmd.Flags().BoolP("namespace", "n", false, "Namespace for the pods")
}
