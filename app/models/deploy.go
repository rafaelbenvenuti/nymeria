package models

import (
	"time"
	"github.com/revel/revel"
)

type Deploy struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Component   string `json:"component"`
	Version     string `json:"version"`
	Accountable string `json:"accountable"`
	Status      string `json:"status"`
	Duration    int    `json:"duration"`
}

func (deploy *Deploy) Validate(v *revel.Validation) {
	v.Required(deploy.Component)
	v.Required(deploy.Version)
	v.Required(deploy.Accountable)
	v.Required(deploy.Status)
	v.Required(deploy.Duration)
}
