package ExportIT

import (
	"github.com/spf13/cobra"
)

func CliNamedParam(fn any) func() (*cobra.Command, error) {
	return func() (*cobra.Command, error) {
		return cliNamedParam(fn)
	}
}

func ExecCli(appName, short, Long string, funcs ...func() (*cobra.Command, error)) error {
	rootCmd := &cobra.Command{
		Use:   appName,
		Short: short,
		Long:  Long,
	}
	for _, fn := range funcs {
		cmd, err := fn()
		if err != nil {
			return err
		}
		rootCmd.AddCommand(cmd)
	}
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
