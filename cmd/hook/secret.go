package hook

import (
	"fmt"
	"io"

	"github.com/benmatselby/prolificli/client"
	"github.com/spf13/cobra"
)

// ListOptions is the options for the listing secrets command.
type ListSecretOptions struct {
	Args        []string
	WorkspaceID string
}

// NewListSecretCommand creates a new `hook secrets` command to give you details about
// your secrets.
func NewListSecretCommand(commandName string, client client.API, w io.Writer) *cobra.Command {
	var opts ListSecretOptions

	cmd := &cobra.Command{
		Use:   commandName,
		Short: "List your hook secrets",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Args = args
			err := renderSecrets(client, opts, w)
			if err != nil {
				return fmt.Errorf("error: %s", err.Error())
			}

			return nil
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.WorkspaceID, "workspace", "w", "", "Filter secrets by workspace.")

	return cmd
}

// renderSecrets will show all of the secrets in a given workspace.
func renderSecrets(client client.API, opts ListSecretOptions, w io.Writer) error {
	secrets, err := client.GetHookSecrets(opts.WorkspaceID)
	if err != nil {
		return err
	}

	for _, secret := range secrets.Results {
		fmt.Fprintln(w, secret.Value)
	}

	return nil
}
