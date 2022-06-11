package user

import (
	"github.com/benmatselby/prolificli/client"
	"github.com/spf13/cobra"
)

// NewUserCommand creates a new `user` command
func NewUserCommand(client client.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "User related commands",
	}

	cmd.AddCommand(
		NewMeCommand(client),
	)
	return cmd
}
