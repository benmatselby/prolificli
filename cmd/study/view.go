package study

import (
	"fmt"
	"os"
	"strings"

	"github.com/benmatselby/prolificli/client"
	studyui "github.com/benmatselby/prolificli/ui/study"
	"github.com/spf13/cobra"
)

// NewViewCommand creates a new `study view` command to give you details about
// your studies.
func NewViewCommand(client client.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view",
		Short: "Provide details about your study, requires a Study ID",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) < 1 {
				fmt.Println("Requests a Study ID")
				os.Exit(1)
			}

			study, err := client.GetStudy(args[0])
			if err != nil {
				fmt.Printf("Error: %s", strings.ReplaceAll(err.Error(), "\n", ""))
				os.Exit(1)
			}

			fmt.Fprintln(os.Stdout, studyui.RenderStudy(client, *study))
		},
	}

	return cmd
}
