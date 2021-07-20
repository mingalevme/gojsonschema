package cmd

import (
	"dmp/gojsonschema/internal"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	RootCmd = &cobra.Command{
		Use:  "gojsonschema",
		Args: cobra.MinimumNArgs(1),
		Run:  RootRunCmd,
	}
	schema string
)

func Execute() {
	_ = RootCmd.Execute()
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&schema, "schema", "s", "schema.json", "schema URL")
}

func RootRunCmd(cmd *cobra.Command, args []string) {
	hasErrors := false
	for _, document := range args {
		result, err := internal.Validate(schema, document)
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				panic(err)
			}
			os.Exit(1)
		}
		if result.Valid() {
			fmt.Printf("%s: Ok\n", document)
			continue
		}
		hasErrors = true
		fmt.Printf("%s: The document is not valid:\n", document)
		for _, desc := range result.Errors() {
			fmt.Printf("  - %s\n", desc)
		}
	}
	if hasErrors {
		os.Exit(1)
	}
}
