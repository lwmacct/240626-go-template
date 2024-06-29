// file: cmd.go
package app

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"

	"github.com/spf13/cobra"
)

func getPackageName() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Base(path.Dir(filename))
}
func printFlags() {
	v := &CmdFlags.run
	if vJson, err := json.MarshalIndent(v, "", "  "); err == nil {
		fmt.Printf("%s\n", vJson)
	}
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
		Test string
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
			printFlags()
			Main(CmdFlags.run.Test)
		},
	}
	co := CmdOptions[name].Flags()
	co.StringVar(&CmdFlags.run.Test, "test", "", "测试参数")

	mft := []string{
		"test",
	}
	for _, v := range mft {
		CmdOptions[name].MarkFlagRequired(v)
	}
}
