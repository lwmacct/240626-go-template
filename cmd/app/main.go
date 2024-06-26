package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/lwmacct/240626-go-template/plugin/app"
	"github.com/spf13/cobra"
)

var (
	version   string = "Unknown"
	commit    string = "Unknown"
	buildTime string = "Unknown"
	developer string = "Unknown"

	rootCmd = &cobra.Command{
		Use:               "App",
		Short:             "App is a command-line tool",
		CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
	}

	verCmd = &cobra.Command{
		Use:   "version",
		Short: "打印版本信息",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("AppVersion:   %s\n", version)
			fmt.Printf("Go Version:   %s\n", runtime.Version())
			fmt.Printf("Git Commit:   %s\n", commit)
			fmt.Printf("Build Time:   %s\n", buildTime)
			fmt.Printf("Developer :   %s\n", developer)
		},
	}
)

func init() {
	rootCmd.AddCommand(verCmd)
	var cmds []*cobra.Command
	cmds = append(cmds, app.CmdMenu())

	for _, cmd := range cmds {
		rootCmd.AddCommand(cmd)
	}

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
