package cmd

import (
	"ExportIT/internal/another"
	"github.com/spf13/cobra"
)

var (
	CatCmd_path string
)

// CatCmd represents the Cat command
var CatCmd = &cobra.Command{
	Use:   "Cat",
	Short: "another.Cat",
	Long: `export:
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return another.Cat(
			CatCmd_path,
		)
	},
}

func init() {
	rootCmd.AddCommand(CatCmd)
	CatCmd.Flags().StringVarP(&CatCmd_path, "path", "p", "", "")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// $1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// $1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
