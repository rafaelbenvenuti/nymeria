package models

import (
	"github.com/revel/revel"
	"time"
)

type Status struct {
	ID       uint      `gorm:"primary_key"`
	Status   string    `json:"status"`
	Date     time.Time `json:"date"`
	DeployID uint
}

func (status *Status) Validate(validation *revel.Validation) {
	status.Date = time.Now()
	validation.Required(status.Status)
}
