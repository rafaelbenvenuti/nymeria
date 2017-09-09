package controllers

import (
	"github.com/rafaelbenvenuti/nymeria/app"
	"github.com/rafaelbenvenuti/nymeria/app/models"
	"github.com/revel/revel"
)

// Metada used in responses API responses.
// All the messages and return codes should be documented elsewhere in the future.
type Meta struct {
	ID      int
	Message string
}

// Define the Deploys Controller
type Deploys struct {
	*revel.Controller
}

func (c Deploys) Create() revel.Result {
	// Prepare an object to return as the response in the end.
	response := make(map[string]interface{})

	// Log information about this request data.
	revel.INFO.Println("Request to create new deploy received.")
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	revel.INFO.Println("Deploy and Status data received: ", jsonData)

	// Bind data into the Deploy and Status object.
	revel.INFO.Println("Binding data into Deploy and Status objects.")
	var deploy models.Deploy
	c.Params.BindJSON(&deploy)
	var status models.Status
	c.Params.BindJSON(&status)

	// Validate the Deploy data.
	revel.INFO.Println("Validating that the Deploy data is consistent.")
	deploy.Validate(c.Validation)
	if c.Validation.HasErrors() {
		revel.ERROR.Println("Deploy data structure is invalid.")
		c.Response.Status = 400
		response["meta"] = Meta{ID: 100, Message: "invalid deploy data structure."}
		response["data"] = nil
		return c.RenderJSON(response)
	}

	// Validate the Status data.
	revel.INFO.Println("Validating that the Status data is consistent.")
	status.Validate(c.Validation)
	if c.Validation.HasErrors() {
		revel.ERROR.Println("Status data structure is invalid.")
		c.Response.Status = 400
		response["meta"] = Meta{ID: 101, Message: "invalid status data structure."}
		response["data"] = nil
		return c.RenderJSON(response)
	}

	// Create a new deploy if data about the requested deploy is not found.
	revel.INFO.Println("Finding if the Deploy already exists in the database.")
	if app.Database.Preload("Statuses").Where(&deploy).First(&deploy).Error != nil {
		if app.Database.Preload("Statuses").Where(&deploy).First(&deploy).RecordNotFound() {
			revel.INFO.Println("Deploy does not exists in the database. Creating...")
			if app.Database.Create(&deploy).Error != nil {
				revel.ERROR.Println("Unable to save deploy data internally.")
				c.Response.Status = 500
				response["meta"] = Meta{ID: 200, Message: "unable to save deploy data internally."}
				response["data"] = deploy
				return c.RenderJSON(response)
			}
		} else {
			revel.ERROR.Println("Unable to retrieve data internally.")
			c.Response.Status = 500
			response["meta"] = Meta{ID: 201, Message: "unable to retrieve data internally."}
			response["data"] = nil
			return c.RenderJSON(response)
		}
	}

	// Ensure that the requested deploy does not already have that status.
	revel.INFO.Println("Cheking if the Status data already exists in the database.")
	statusFound := false
	for _, deploy_status := range deploy.Statuses {
		if deploy_status.Status == status.Status {
			statusFound = true
		}
	}

	// Create a new status for this deploy if data about the requested status is nonexistent.
	if statusFound == true {
		revel.INFO.Println("Status already exists in the database.")
		c.Response.Status = 409
		response["meta"] = Meta{ID: 401, Message: "status already exists."}
		response["data"] = deploy
		return c.RenderJSON(response)
	} else {
		revel.INFO.Println("Status does not exists in the database. Updating the Deploy with new Status...")
		deploy.Statuses = append(deploy.Statuses, status)
		if app.Database.Save(&deploy).Error != nil {
			revel.ERROR.Println("Unable to save status data internally.")
			c.Response.Status = 500
			response["meta"] = Meta{ID: 202, Message: "unable to save status data internally."}
			response["data"] = deploy
			return c.RenderJSON(response)
		} else {
		  revel.INFO.Println("Status created successfully.")
		}
	}

	// Return 201 if deploy was successfully created and associated with status.
  revel.INFO.Println("Request to create new data completed successfully.")
	c.Response.Status = 201
	response["meta"] = Meta{ID: 10, Message: "data created successfully."}
	response["data"] = deploy
	return c.RenderJSON(response)

}

func (c Deploys) List() revel.Result {
	// Prepare an object to return as the response in the end.
	response := make(map[string]interface{})

	// Log information about this request data.
	revel.INFO.Println("Request to list all Deploys and Statuses received.")

	// Retrieve all Deploys and Statuses from database.
  revel.INFO.Println("Trying to retrieve all Deploy and Status data from database.")
	deploys := []models.Deploy{}
	if app.Database.Preload("Statuses").Find(&deploys).Error != nil {
    revel.ERROR.Println("Unable to retrieve data internally.")
		c.Response.Status = 500
		response["meta"] = Meta{ID: 201, Message: "unable to retrieve data internally."}
		response["data"] = nil
		return c.RenderJSON(response)
	}

	// Return 200 code and all deploy data.
  revel.INFO.Println("Deploy and Status data successfully retrieved.")
	c.Response.Status = 200
	response["meta"] = Meta{ID: 30, Message: "all data successfully retrieved."}
	response["data"] = deploys
	return c.RenderJSON(response)
}

func (c Deploys) Show(id int) revel.Result {
	// Prepare an object to return as the response in the end.
	response := make(map[string]interface{})

	// Log information about this request data.
	revel.INFO.Println("Request to show Deploy and Statuses with Deploy ID: ", id, "received.")

	// Try to retrieve the object in the database.
  revel.INFO.Println("Retrieve Deploy and Status data for Deploy with ID: ", id)
	var deploy models.Deploy
	if app.Database.Preload("Statuses").First(&deploy, id).Error != nil {
		// Return a 404 if the record is not found.
		if app.Database.Preload("Statuses").First(&deploy, id).RecordNotFound() {
      revel.INFO.Println("Deploy record not found.")
			c.Response.Status = 404
			response["meta"] = Meta{ID: 400, Message: "record not found."}
			response["data"] = nil
			return c.RenderJSON(response)
		} else {
      revel.ERROR.Println("Unable to retrieve data internally.")
			c.Response.Status = 500
			response["meta"] = Meta{ID: 201, Message: "unable to retrieve data internally."}
			response["data"] = nil
			return c.RenderJSON(response)
		}
	}

	// Return 200 if the deploy was successfully retrieved.
  revel.INFO.Println("Deploy and Status data successfully retrieved.")
	c.Response.Status = 200
	response["meta"] = Meta{ID: 20, Message: "data retrieved successfully."}
	response["data"] = deploy
	return c.RenderJSON(response)
}

func (c Deploys) Delete(id int) revel.Result {
	// Prepare an object to return as the response in the end.
	response := make(map[string]interface{})

	// Log information about this request data.
	revel.INFO.Println("Request to delete Deploy and Statuses with Deploy ID: ", id, "received.")

	// Try to retrieve the object in the database.
  revel.INFO.Println("Retrieve Deploy and Status data for Deploy with ID: ", id)
	var deploy models.Deploy
	if app.Database.Preload("Statuses").First(&deploy, id).Error != nil {
		// Return a 404 if the record is not found.
		if app.Database.Preload("Statuses").First(&deploy, id).RecordNotFound() {
      revel.INFO.Println("Deploy record not found.")
			c.Response.Status = 404
			response["meta"] = Meta{ID: 400, Message: "record not found."}
			response["data"] = nil
			return c.RenderJSON(response)
		} else {
      revel.ERROR.Println("Unable to retrieve data internally.")
			c.Response.Status = 500
			response["meta"] = Meta{ID: 201, Message: "unable to retrieve data internally."}
			response["data"] = nil
			return c.RenderJSON(response)
		}
	}

	// Record has been found, so we must remove the record.
  revel.INFO.Println("Removing Deploy and Status data.")
	if app.Database.Delete(&deploy).Error != nil {
    revel.ERROR.Println("Unable to remove data internally.")
		c.Response.Status = 500
		response["meta"] = Meta{ID: 203, Message: "unable to remove data internally."}
		response["data"] = deploy
		return c.RenderJSON(response)
	}

	// Return 200 if the deploy was successfully retrieved.
  revel.INFO.Println("Deploy and Status data successfully removed.")
	c.Response.Status = 200
	response["meta"] = Meta{ID: 21, Message: "data removed successfully."}
	response["data"] = deploy
	return c.RenderJSON(response)
}
