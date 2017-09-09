package models

import (
	"github.com/revel/revel"
)

type Deploy struct {
	ID          uint     `gorm:"primary_key"`
	Component   string   `json:"component"`
	Version     string   `json:"version"`
	Accountable string   `json:"accountable"`
	Statuses    []Status `json:"statuses"`
}

func (deploy *Deploy) Validate(validation *revel.Validation) {
	validation.Required(deploy.Component)
	validation.Required(deploy.Version)
	validation.Required(deploy.Accountable)
}
