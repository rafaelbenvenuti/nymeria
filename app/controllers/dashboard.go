package controllers

import (
	"github.com/rafaelbenvenuti/nymeria/app"
	"github.com/rafaelbenvenuti/nymeria/app/models"
	"github.com/revel/revel"
	"time"
)

// Use this struct for dashboard elements.
type DataElement struct {
	Deploy    models.Deploy
	Durations map[string]time.Duration
	Total     time.Duration
}

// Define the Dashboard Controller
type Dashboard struct {
	*revel.Controller
}

func (c Dashboard) Show() revel.Result {
	// Log information about this request data.
	revel.INFO.Println("Request to show the Dashboard received.")

	// Retrieve all Deploys and Statuses from database.
	revel.INFO.Println("Trying to retrieve all Deploy and Statuses data from database.")
	deploys := []models.Deploy{}
	if app.Database.Preload("Statuses").Find(&deploys).Error != nil {
		revel.ERROR.Println("Unable to retrieve data internally.")
		c.Response.Status = 500
		return c.RenderText("Nymeria had a fatal error. Please contact the administrator.")
	}

	// Create map that contains durations for all Statuses for each Deploy.
	dashboardData := []DataElement{}
	for _, deploy := range deploys {
		if len(deploy.Statuses) > 1 {
			durations := make(map[string]time.Duration)
			previousStatus := deploy.Statuses[0]
			for _, currentStatus := range deploy.Statuses[1:] {
				durations[currentStatus.Status] = currentStatus.Date.Sub(previousStatus.Date)
				previousStatus = currentStatus
			}
			dashboardData = append(dashboardData, DataElement{
				Deploy:    deploy,
				Durations: durations,
				Total:     deploy.Statuses[len(deploy.Statuses)-1].Date.Sub(deploy.Statuses[0].Date),
			})
		}
	}

	revel.INFO.Println(dashboardData)

	// Return 200 code and all deploy data.
	revel.INFO.Println("Deploy and Status data successfully retrieved.")
	c.Response.Status = 200
	return c.Render(dashboardData)
}
