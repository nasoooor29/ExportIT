package cmd

import (
	"ExportIT/internal"
	"github.com/spf13/cobra"
)

var (
	BatchFilesCmd_dest         string
	BatchFilesCmd_src          string
	BatchFilesCmd_fileExt      string
	BatchFilesCmd_batchSize    int
	BatchFilesCmd_startDateStr string
)

// BatchFilesCmd represents the BatchFiles command
var BatchFilesCmd = &cobra.Command{
	Use:   "BatchFiles",
	Short: "internal.BatchFiles",
	Long: `export: BatchFiles organizes files by given extension into batches and saves them into folders starting from the start date.
mama
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return internal.BatchFiles(
			BatchFilesCmd_dest,
			BatchFilesCmd_src,
			BatchFilesCmd_fileExt,
			BatchFilesCmd_batchSize,
			BatchFilesCmd_startDateStr,
		)
	},
}

func init() {
	rootCmd.AddCommand(BatchFilesCmd)
	BatchFilesCmd.Flags().StringVarP(&BatchFilesCmd_dest, "dest", "d", "", "")
	BatchFilesCmd.Flags().StringVarP(&BatchFilesCmd_src, "src", "s", "", "")
	BatchFilesCmd.Flags().StringVarP(&BatchFilesCmd_fileExt, "fileExt", "f", "", "")
	BatchFilesCmd.Flags().IntVarP(&BatchFilesCmd_batchSize, "batchSize", "b", 0, "")
	BatchFilesCmd.Flags().StringVarP(&BatchFilesCmd_startDateStr, "startDateStr", "t", "", "")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// $1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// $1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
