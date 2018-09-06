package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
	"github.com/hashicorp/hcl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var convCmd = &cobra.Command{
	Use:   "conv",
	Short: "Convert HCL in stdin to another format.",
	Run: func(cmd *cobra.Command, args []string) {
		convert()
	},
}

const (
	FMT_YAML = "yaml"
	FMT_JSON = "json"
)

const (
	KeyConvertFormat = "convert.format"
)

func init() {
	RootCmd.AddCommand(convCmd)

	pflags := convCmd.PersistentFlags()
	pflags.StringP("format", "f", FMT_YAML, fmt.Sprintf("Format to convert. Support %v and %v.", FMT_YAML, FMT_JSON))
	viper.BindPFlag(KeyConvertFormat, pflags.Lookup("format"))

}

func convert() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var v interface{}
	if err := hcl.Unmarshal(b, &v); err != nil {
		log.Fatalf("%v", err)
	}

	var j []byte
	f := viper.GetString(KeyConvertFormat)
	switch f {
	case FMT_YAML:
		j, err = yaml.Marshal(v)
	case FMT_JSON:
		j, err = json.MarshalIndent(v, "", "  ")
	default:
		err = fmt.Errorf("susupported format. %v", f)
	}
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println(string(j))
}
