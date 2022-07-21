/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"

	"github.com/gonsole/annotation"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	directory := generateCmd.Flags().StringP("directory", "d", "", "Handler or Controller directory path.")
	output := generateCmd.Flags().StringP("output", "o", "", "Output file path; routes.go")

	generateCmd.MarkFlagRequired("directory")
	generateCmd.MarkFlagRequired("output")

	annotation := new(annotation.App)

	generateCmd.Run = func(cmd *cobra.Command, args []string) {
		annotation.Config.Directory = *directory
		annotation.Config.Output = *output
		annotation.ParseDirectory()
		annotation.Generate()
	}

}
