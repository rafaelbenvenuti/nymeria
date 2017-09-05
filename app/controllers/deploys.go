package controllers

import (
	"github.com/rafaelbenvenuti/nymeria/app"
	"github.com/rafaelbenvenuti/nymeria/app/models"
	"github.com/revel/revel"
)

// Metada used in responses API responses.
// All the messages and return codes should be documented in the future.
type Meta struct {
	ID      int
	Message string
}

type Deploys struct {
	*revel.Controller
}

func (c Deploys) Create(deploy *models.Deploy) revel.Result {
	// Prepare an object to return as the response in the end.
	response := make(map[string]interface{})

	// Validate if the deploy requested to be created is valid.
	deploy.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Response.Status = 400
		response["meta"] = Meta{ID: 100, Message: "invalid data structure."}
		response["data"] = nil
		return c.RenderJSON(response)
	}

	// Try to store the object in the database.
	op := app.Database.Create(&deploy)
	if op.Error != nil {
		c.Response.Status = 500
		response["meta"] = Meta{ID: 200, Message: "unable to save data internally."}
		response["data"] = deploy
		return c.RenderJSON(response)
	}

	// Return 201 if deploy was successfully created.
	c.Response.Status = 201
	response["meta"] = Meta{ID: 10, Message: "deploy created successfully."}
	response["data"] = deploy
	return c.RenderJSON(response)
}

func (c Deploys) Show() revel.Result {
	return c.Todo()
}

func (c Deploys) Delete() revel.Result {
	return c.Todo()
}

func (c Deploys) List() revel.Result {
	results := app.Database.Find(models.Deploy{})
	return c.RenderJSON(results)
}
