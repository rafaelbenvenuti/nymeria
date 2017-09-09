package controllers

import (
	"github.com/rafaelbenvenuti/nymeria/app"
	"github.com/rafaelbenvenuti/nymeria/app/models"
	"github.com/revel/revel"
)

// Define the Dashboard Controller
type Dashboard struct {
	*revel.Controller
}

func (c Dashboard) Show() revel.Result {
	// Log information about this request data.
	revel.INFO.Println("Request to show the Dashboard received.")

	// Retrieve all Deploys and Statuses from database.
  revel.INFO.Println("Trying to retrieve all Deploy and Status data from database.")
	deploys := []models.Deploy{}
	if app.Database.Preload("Statuses").Find(&deploys).Error != nil {
    revel.ERROR.Println("Unable to retrieve data internally.")
		c.Response.Status = 500
		return c.RenderText("Nymeria had a fatal error. Please contact the administrator.")
	}

	// Return 200 code and all deploy data.
  revel.INFO.Println("Deploy and Status data successfully retrieved.")
	c.Response.Status = 200
	return c.Render(deploys)

	// Calculate the deltas between all deploy statuses.
	//deltas := make(map[string]time.Time)

	//revel.INFO.Println("DELTAS","OK")
	//if len(deploys) > 1 {
	//	previousDeploy := deploys[0]
	//	for index, deploy := range deploys[1:] {
	//		delta := deploy.Date.Sub(previousDeploy.Date)
	//	  previousDeploy = deploy
	//    revel.INFO.Println("DELTA",index,delta)
	//	}
	//}
}
