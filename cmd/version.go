/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hxx258456/gdeepl/internal/pkg/metadata"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "version",
	Long:    `Prints gdeepl version`,
	Example: `gdeepl version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(metadata.GetVersion())
	},
}
