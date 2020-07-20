package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var  versionCmd = &cobra.Command{
	Use:                        "version",
	//Aliases:                    nil,
	//SuggestFor:                 nil,
	Short:                      "Print the version number of cobrademo",
	Long:                       `All software has versions. This is cobrademo's`,
	//Example:                    "",
	//ValidArgs:                  nil,
	//ValidArgsFunction:          nil,
	//Args:                       nil,
	//ArgAliases:                 nil,
	//BashCompletionFunction:     "",
	//Deprecated:                 "",
	//Hidden:                     false,
	//Annotations:                nil,
	//Version:                    "",
	//PersistentPreRun:           nil,
	//PersistentPreRunE:          nil,
	//PreRun:                     nil,
	//PreRunE:                    nil,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cobrademo version is v1.0")
	},
	//RunE:                       nil,
	//PostRun:                    nil,
	//PostRunE:                   nil,
	//PersistentPostRun:          nil,
	//PersistentPostRunE:         nil,
	//SilenceErrors:              false,
	//SilenceUsage:               false,
	//DisableFlagParsing:         false,
	//DisableAutoGenTag:          false,
	//DisableFlagsInUseLine:      false,
	//DisableSuggestions:         false,
	//SuggestionsMinimumDistance: 0,
	//TraverseChildren:           false,
	//FParseErrWhitelist:         cobra.FParseErrWhitelist{},
}