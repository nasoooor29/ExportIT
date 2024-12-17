package export

func GetCustomCommandTemplate() string {
	return `package cmd

import (

	"github.com/spf13/cobra"
	"$PKG_PATH"
)

var (
$CLI_VARS
)
// $FUNC_NAMECmd represents the $FUNC_NAME command
var $FUNC_NAMECmd = &cobra.Command{
	Use:   "$FUNC_NAME",
	Short: "$PKG_NAME.$FUNC_NAME",
	Long: $FUNC_COMMENT,
	RunE: func(cmd *cobra.Command, args []string) error {
		return $PKG_NAME.$FUNC_NAME(
$FUNC_PARAMS
		)
	},
}

func init() {
	rootCmd.AddCommand($FUNC_NAMECmd)
$CLI_PARAMS

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// $1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// $1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
`
}
