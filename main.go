package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

var (
	version     string
	versionLong string
)

const (
	KeyVersion = "version"
)

var RootCmd = &cobra.Command{
	Use:   "hclutil",
	Short: "A utilities for HCL.",
	Run: func(c *cobra.Command, args []string) {
		if viper.GetBool(KeyVersion) {
			fmt.Println(version)
			return
		}
		c.Help()
	},
}

func init() {
	pflags := RootCmd.PersistentFlags()
	pflags.BoolP("version", "v", false, "Show version")
	viper.BindPFlag(KeyVersion, pflags.Lookup("version"))
}
