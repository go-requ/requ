package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var variableFilePath string

var rootCmd = &cobra.Command{
	Use:   "requ [HTTP file path]",
	Short: "A convenient cli for IntelliJ HTTP client file",
	Run:   main,
	Args:  checkArgument,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// TODO: Add Variable related option here.
	// rootCmd.PersistentFlags().StringVar(&httpFile, "httpFile", "", "Http file path (ex: test.http or test.rest)")
	// rootCmd.MarkPersistentFlagRequired("httpFile")
	rootCmd.PersistentFlags().StringVarP(&variableFilePath, "variable", "v", "", "Http Varable file path")
}

func checkArgument(cmd *cobra.Command, args []string) error {
	// When argument is empty
	if len(args) < 1 {
		return errors.New("requires HTTP file path. (ex: test.http, test.rest)")
	}

	// When extension is not '.http' or '.rest'
	var filePath = args[0]
	if !strings.HasSuffix(filePath, ".http") && !strings.HasSuffix(filePath, ".rest") {
		return errors.New("requires HTTP file. extension must be .http or .rest")
	}

	// When file is not exist
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New("invalid HTTP file path. file does not exist")
	}

	if variableFilePath != "" {
		if !strings.HasSuffix(variableFilePath, ".json") {
			return errors.New("invalid HTTP variable file. extension must be .json")
		}

		// When Variable file is not exist
		if _, err := os.Stat(variableFilePath); os.IsNotExist(err) {
			return errors.New("invalid HTTP variable file path. file does not exist")
		}
	}

	return nil
}
