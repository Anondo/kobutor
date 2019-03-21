package service

import (
	"encoding/json"
	"fmt"
	"kobutor/helper"
	"log"
	"net/url"

	"github.com/asaskevich/govalidator"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/spf13/viper"
)

// Disposition ...
const (
	DispositionInline     = "inline"
	DispositionAttachment = "attachment"
	sendGridMaxSize       = 30 * 1000000
)

// Email ...
type Email struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Personalization ...
type Personalization struct {
	To      []Email `json:"to"`
	Subject string  `json:"subject"`
}

// Content ...
type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Attachment ...
type Attachment struct {
	// ContentID   string `json:"content_id,omitempty"`
	Type        string `json:"type,omitempty"`
	FileName    string `json:"filename,omitempty"`
	Content     string `json:"content,omitempty"`
	Disposition string `json:"disposition,omitempty"` // inline or attachment
}

// SendGridMail ...
type SendGridMail struct {
	Personalizations []Personalization `json:"personalizations"`
	From             Email             `json:"from"`
	Content          []Content         `json:"content"`
	Attachments      []*Attachment     `json:"attachments,omitempty"`
}

// SendGridRequest ...
type SendGridRequest struct {
	From        Email         `json:"from"`
	To          Email         `json:"to"`
	Subject     string        `json:"subject"`
	Body        string        `json:"body"`
	Type        string        `json:"type"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

// Send sends the email using the sendgrid api
func (s *SendGridRequest) Send() error {
	sgm := SendGridMail{
		Personalizations: []Personalization{
			{
				To:      []Email{s.To},
				Subject: s.Subject,
			},
		},
		From: s.From,
		Content: []Content{
			{
				Type:  s.Type,
				Value: s.Body,
			},
		},
		Attachments: s.Attachments,
	}

	sgmBytes, err := json.Marshal(sgm)

	if err != nil {
		return err
	}
	request := sendgrid.GetRequest(viper.GetString("sendgrid.key"), viper.GetString("sendgrid.endpoint"),
		viper.GetString("sendgrid.host"))
	request.Method = "POST"
	request.Body = sgmBytes

	resp, err := sendgrid.API(request)

	if err != nil {
		return err
	}

	log.Printf("Response: Status_Code=%d | Body=%s | Headers=%v\n",
		resp.StatusCode, resp.Body, resp.Headers)

	return nil
}

// Validate validates the send grid request payload
func (s *SendGridRequest) Validate() *url.Values {
	errs := &url.Values{}

	if s.From.Email == "" || !govalidator.IsEmail(s.From.Email) {
		errs.Add("from.email", "The from.email field is invalid")
	}

	if s.From.Name == "" {
		errs.Add("from.name", "The from.name field is requied")
	}

	if s.To.Email == "" || !govalidator.IsEmail(s.To.Email) {
		errs.Add("to.email", "The to.email field is invalid")
	}

	if s.To.Name == "" {
		errs.Add("to.name", "The to.name field is requied")
	}

	if s.Subject == "" {
		errs.Add("subject", "The subject field is required")
	}

	if s.Body == "" {
		errs.Add("body", "The body field is required")
	}

	//validate attachments
	aSize := 0
	for i, a := range s.Attachments {
		aSize = aSize + len([]byte(a.Content))

		if a.FileName == "" {
			errs.Add("attachments.filename", fmt.Sprintf("attachment.filename[%d] is invalid", i))
		}
		if a.Content == "" {
			errs.Add("attachments.content", fmt.Sprintf("attachment.content[%d] is invalid", i))
		}
		if a.Type == "" {
			errs.Add("attachments.type", fmt.Sprintf("attachment.type[%d] is invalid", i))
		}
		if a.Disposition != DispositionInline && a.Disposition != DispositionAttachment {
			errs.Add("attachments.disposition", fmt.Sprintf("attachment.disposition[%d] is invalid. Must be inline/attachment", i))
		}

		if !helper.IsBase64Encoded(a.Content) {
			a.Content = helper.EncodeToBase64(a.Content)
		}
	}

	if aSize > sendGridMaxSize {
		errs.Add("attachment.content", "max attachment size for sendgrid is 30MB")
	}

	if errs.Encode() == "" {
		return nil
	}
	return errs
}
