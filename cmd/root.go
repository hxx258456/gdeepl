/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/hxx258456/gdeepl/internal/pkg/client"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gdeepl",
	Short: "gdeepl",
	Long:  `gdeepl translate's command line tools`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		req := &client.Request{
			Text:       cmd.Flag("text").Value.String(),
			TargetLang: cmd.Flag("target_lang").Value.String(),
			SourceLang: cmd.Flag("source_lang").Value.String(),
		}
		if err := client.Post(req); err != nil {
			return err
		}
		return nil
	},
	Example: `gdeepl -T "hello"`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gdeepl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(versionCmd)

	rootCmd.Flags().StringP("source_lang", "s", "auto", "源语言")
	rootCmd.Flags().StringP("target_lang", "t", "ZH", "目标语言")
	rootCmd.Flags().StringP("text", "T", "hello world!", "翻译内容")
	rootCmd.MarkFlagRequired("text")
}
