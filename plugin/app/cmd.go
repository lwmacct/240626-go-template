// file: cmd.go
package app

import (
	"path"
	"runtime"

	"github.com/spf13/cobra"
)

func getPackageName() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Base(path.Dir(filename))
}

func CmdMenu() *cobra.Command {
	// 创建新的 rootCmd
	var rootCmd = &cobra.Command{
		Use:   getPackageName(),
		Short: "",
	}
	for _, cmd := range CmdOptions {
		rootCmd.AddCommand(cmd)
	}
	return rootCmd
}

var CmdOptions = make(map[string]*cobra.Command)
var CmdFlags = struct {
	run struct {
		arg string
	}

	DebugStr string
	DebugMap map[string]string
}{}

func init() {
	name := "run"
	CmdOptions[name] = &cobra.Command{
		Use:   name,
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
	co := CmdOptions[name].Flags()
	co.StringVar(&CmdFlags.run.arg, "arg", "", "测试参数")

	mft := []string{
		"arg",
	}
	for _, v := range mft {
		CmdOptions[name].MarkFlagRequired(v)
	}
}
