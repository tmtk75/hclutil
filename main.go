package main

import (
	_ "encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "gopkg.in/yaml.v2"
)

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

var cfgFile string // given by a flag

const Key = "welcome.message"

var RootCmd = &cobra.Command{
	Use:   "hclutil",
	Short: "A utilities for HCL.",
	//Run: func(cmd *cobra.Command, args []string) {
	//	_main()
	//},
}

func init() {
	cobra.OnInitialize(initConfig)
	pflags := RootCmd.PersistentFlags()
	pflags.StringVarP(&cfgFile, "config", "c", "", "config file (default is ./.config.yaml)")

	//
	opts := []struct {
		long, short, def, desc string // option
		key, env               string // env var
	}{
		{long: "msg", short: "m",
			def: "Hello, world!", desc: "Welcome message",
			key: Key, env: "SPREADSHEET_ID"},
	}

	for _, e := range opts {
		pflags.StringP(e.long, e.short, e.def, e.desc)
		viper.BindPFlag(e.key, pflags.Lookup(e.long))
		viper.BindEnv(e.key, e.env)
	}
}

func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("config") // name of config file (without extension)
	//viper.AddConfigPath("$HOME")          // adding home directory as first search path
	viper.AddConfigPath(".")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
