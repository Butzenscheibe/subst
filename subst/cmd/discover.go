package cmd

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/bedag/subst/pkg/config"
	"github.com/spf13/cobra"
)

func newDiscoverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "discover",
		Short: "Discover if plugin is applicable to the given directory",
		Long: heredoc.Doc(`
			Run 'subst discover' to return directories that contain plugin compatible files. Mainly used for automatic plugin discovery by ArgoCD`),
		RunE: discover,
	}

	flags := cmd.Flags()
	addCommonFlags(flags)
	return cmd

}

func discover(cmd *cobra.Command, args []string) error {
	dir, err := rootDirectory(args)
	if err != nil {
		return err
	}

	configuration, err := config.LoadConfiguration(cfgFile, cmd, dir)
	if err != nil {
		return fmt.Errorf("failed loading configuration: %w", err)
	}

	println(configuration)

	return nil
}
