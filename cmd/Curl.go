package cmd

import (
	"ExportIT/internal/another"
	"github.com/spf13/cobra"
)

var (
	CurlCmd_method      string
	CurlCmd_link        string
	CurlCmd_payloadPath string
)

// CurlCmd represents the Curl command
var CurlCmd = &cobra.Command{
	Use:   "Curl",
	Short: "another.Curl",
	Long: `export: get response from website
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return another.Curl(
			CurlCmd_method,
			CurlCmd_link,
			CurlCmd_payloadPath,
		)
	},
}

func init() {
	rootCmd.AddCommand(CurlCmd)
	CurlCmd.Flags().StringVarP(&CurlCmd_method, "method", "m", "", "")
	CurlCmd.Flags().StringVarP(&CurlCmd_link, "link", "l", "", "")
	CurlCmd.Flags().StringVarP(&CurlCmd_payloadPath, "payloadPath", "p", "", "")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// $1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// $1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
