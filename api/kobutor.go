package api

import (
	"kobutor/helper"
	"net/http"

	"kobutor/service"

	"github.com/thedevsaddam/renderer"
)

// SendMail sends mail
func SendMail(w http.ResponseWriter, r *http.Request) {
	sr := service.SendGridRequest{}

	if err := helper.ParseBody(r.Body, &sr); err != nil {
		renderer.New().JSON(w, http.StatusBadRequest, renderer.M{
			"message": "Invalid request data",
			"error":   err,
		})
		return

	}

	if err := sr.Validate(); err != nil {
		renderer.New().JSON(w, http.StatusBadRequest, renderer.M{
			"message": "Validation error",
			"error":   err,
		})
		return
	}

	if err := sr.Send(); err != nil {
		renderer.New().JSON(w, http.StatusInternalServerError, renderer.M{
			"message": "Failed to send email",
			"error":   err,
		})
		return
	}

	renderer.New().JSON(w, http.StatusOK, renderer.M{
		"message": "Email Sent",
	})
	return

}
