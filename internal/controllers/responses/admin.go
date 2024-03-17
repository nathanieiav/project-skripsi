package responses

import (
	"project-skbackend/internal/models/base"
	"project-skbackend/packages/consttypes"
	"project-skbackend/packages/customs"
)

type (
	Admin struct {
		base.Model

		User User `json:"user"`

		FirstName   string            `json:"first_name"`
		LastName    string            `json:"last_name"`
		Gender      consttypes.Gender `json:"gender"`
		DateOfBirth customs.CDT_DATE  `json:"date_of_birth"`
	}
)
