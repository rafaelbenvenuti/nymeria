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

func (c Deploys) Show(id int) revel.Result {
	// Prepare an object to return as the response in the end.
	response := make(map[string]interface{})

	// Try to retrieve the object in the database.
	var deploy models.Deploy
	op := app.Database.First(&deploy, id)
	if op.Error != nil {
		// Return a 404 if the record is not found.
		if op.RecordNotFound() {
			c.Response.Status = 404
			response["meta"] = Meta{ID: 400, Message: "record not found."}
			response["data"] = nil
			return c.RenderJSON(response)

			// Return an error if data can't be accessed in the database.
		} else {
			c.Response.Status = 500
			response["meta"] = Meta{ID: 201, Message: "unable to retrieve data internally."}
			response["data"] = nil
			return c.RenderJSON(response)
		}

	}

	// Return 200 if the deploy was successfully retrieved.
	c.Response.Status = 200
	response["meta"] = Meta{ID: 20, Message: "deploy retrieved successfully."}
	response["data"] = deploy
	return c.RenderJSON(response)
}

func (c Deploys) Delete(id int) revel.Result {
	// Prepare an object to return as the response in the end.
	response := make(map[string]interface{})

	// Try to retrieve the object in the database.
	var deploy models.Deploy
	op := app.Database.First(&deploy, id)
	if op.Error != nil {
		// Return a 404 if the record is not found.
		if op.RecordNotFound() {
			c.Response.Status = 404
			response["meta"] = Meta{ID: 400, Message: "record not found."}
			response["data"] = nil
			return c.RenderJSON(response)

			// Return an error if data can't be accessed in the database.
		} else {
			c.Response.Status = 500
			response["meta"] = Meta{ID: 201, Message: "unable to retrieve data internally."}
			response["data"] = nil
			return c.RenderJSON(response)
		}
	}

	// Record has been found, so we must remove the record.
	op = app.Database.Delete(&deploy)
	if op.Error != nil {
		c.Response.Status = 500
		response["meta"] = Meta{ID: 202, Message: "unable to remove data internally."}
		response["data"] = deploy
		return c.RenderJSON(response)
	}

	// Return 200 if the deploy was successfully retrieved.
	c.Response.Status = 200
	response["meta"] = Meta{ID: 21, Message: "deploy removed successfully."}
	response["data"] = deploy
	return c.RenderJSON(response)
}

func (c Deploys) List() revel.Result {
	// Prepare an object to return as the response in the end.
	response := make(map[string]interface{})

	// Retrieve all deploys from database.
	deploys := []models.Deploy{}
	op := app.Database.Find(&deploys)
	if op.Error != nil {
		c.Response.Status = 500
		response["meta"] = Meta{ID: 201, Message: "unable to retrieve data internally."}
		response["data"] = nil
		return c.RenderJSON(response)
	}

	// Return 200 code and all deploy data.
	c.Response.Status = 200
	response["meta"] = Meta{ID: 30, Message: "all data successfully retrieved."}
	response["data"] = deploys
	return c.RenderJSON(response)
}
