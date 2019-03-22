package cmd

import (
	"encoding/json"
	"kobutor/service"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	SendCmd = &cobra.Command{
		Use:   "send",
		Short: "send emails from the terminal",
		Run:   sendEmail,
	}
)

func init() {
	SendCmd.Flags().StringP("from", "f", "", "the email address of the sender")
	SendCmd.Flags().StringP("to", "t", "", "the email address of the receiver")
	SendCmd.Flags().StringP("subject", "s", "", "the subject of the email")
	SendCmd.Flags().StringP("body", "b", "", "the body of the email")
	SendCmd.Flags().StringP("type", "p", "", "the type of the content body")
	viper.BindPFlags(SendCmd.Flags())
}

func sendEmail(cmd *cobra.Command, args []string) {
	from := viper.GetString("from")
	to := viper.GetString("to")
	subject := viper.GetString("subject")
	body := viper.GetString("body")
	tpe := viper.GetString("type")

	if from == "" || to == "" || subject == "" || body == "" || tpe == "" {
		log.Println("The flags: -f(from), -t(to), -s(subject), -b(body), --type  must be provided")
		return
	}

	fe := service.Email{}
	te := service.Email{}

	if err := json.Unmarshal([]byte(from), &fe); err != nil {
		log.Fatal("Could not decode the from flag")
	}
	if err := json.Unmarshal([]byte(to), &te); err != nil {
		log.Fatal("Could not decode the from flag")
	}

	sr := service.SendGridRequest{
		From:    fe,
		To:      te,
		Subject: subject,
		Body:    body,
		Type:    tpe,
	}

	if err := sr.Validate(); err != nil {
		log.Fatal(err)
	}

	if err := sr.Send(); err != nil {
		log.Fatal(err)
	}
}
