package study

import (
	"github.com/benmatselby/prolificli/client"
	"github.com/spf13/cobra"
)

// NewStudyCommand creates a new `study` command
func NewStudyCommand(client client.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "study",
		Short: "Study related commands",
	}

	cmd.AddCommand(
		NewListCommand("list", client),
	)
	return cmd
}
