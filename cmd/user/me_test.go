package user_test

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/benmatselby/prolificli/client"
	"github.com/benmatselby/prolificli/cmd/user"
	"github.com/benmatselby/prolificli/mock_client"
	"github.com/golang/mock/gomock"
)

func TestNewMeCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := mock_client.NewMockAPI(ctrl)

	cmd := user.NewMeCommand(client)

	use := "me"
	short := "Provide details about your account"

	if cmd.Use != use {
		t.Fatalf("expected use: %s; got %s", use, cmd.Use)
	}

	if cmd.Short != short {
		t.Fatalf("expected use: %s; got %s", short, cmd.Short)
	}
}

func TestRenderMe(t *testing.T) {
	tt := []struct {
		name   string
		output string
		err    error
	}{
		{name: "can return a list of sprints", output: `

   Bald Rick


  Email:             baldrick@turnip.co
  Available balance: £10.00
  Balance:           £8.50


`, err: nil},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			c := mock_client.NewMockAPI(ctrl)

			me := client.Me{
				FirstName:        "Bald",
				LastName:         "Rick",
				Email:            "baldrick@turnip.co",
				CurrencyCode:     "GBP",
				AvailableBalance: 1000,
				Balance:          850,
			}

			c.
				EXPECT().
				GetMe().
				Return(&me, nil).
				AnyTimes()

			var b bytes.Buffer
			writer := bufio.NewWriter(&b)

			user.RenderMe(c, writer)
			writer.Flush()

			if strings.ReplaceAll(b.String(), " ", "") != strings.ReplaceAll(tc.output, " ", "") {
				t.Fatalf("expected \n'%s'\ngot\n'%s'", tc.output, b.String())
			}
		})
	}
}
