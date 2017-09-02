package controllers

import (
	"github.com/revel/revel"
	"github.com/rafaelbenvenuti/nymeria/app"
	"github.com/rafaelbenvenuti/nymeria/app/models"
)

type Deploys struct {
	*revel.Controller
}

func (c Deploys) List() revel.Result {

	deploy := &models.Deploy{"C1", "v1.3.2", "Someone", "OK"}

	app.Database.Create(deploy)

	return c.RenderJSON(deploy)
}
