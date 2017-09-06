package models

import (
	"time"
	"regexp"
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
	v.MinSize(deploy.Component ,4)
	v.MaxSize(deploy.Component, 80)
	v.Required(deploy.Version)
	v.Required(deploy.Accountable)
	v.MinSize(deploy.Accountable ,4)
	v.MaxSize(deploy.Accountable, 80)
	v.Required(deploy.Status)
	v.Match(deploy.Status, regexp.MustCompile("success|failed"))
	v.Required(deploy.Duration)
}
