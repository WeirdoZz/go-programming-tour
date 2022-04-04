package cmd

import "github.com/spf13/cobra"

/*
放置根命令
*/

var rootCmd = &cobra.Command{}

// Execute 执行根命令
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	//向根命令中注册添加子命令
	rootCmd.AddCommand(wordCmd, timeCmd, sqlCmd)
}
