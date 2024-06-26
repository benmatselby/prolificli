package hook

import (
	"io"

	"github.com/benmatselby/prolificli/client"
	"github.com/spf13/cobra"
)

// NewHookCommand creates a new `hook` command
func NewHookCommand(client client.API, w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hook",
		Short: "Manage and view your hook subscriptions",
		Long: `Manage your hook subscriptions.

A hook subscription registers your interest to be notified of events happening
in the the Prolific Platform.`,
	}

	cmd.AddCommand(
		NewListCommand("list", client, w),
		NewEventTypeCommand("event-types", client, w),
		NewListSecretCommand("secrets", client, w),
		NewEventListCommand("events", client, w),
	)

	return cmd
}
