package controllers

import (
	"time"
	"github.com/rafaelbenvenuti/nymeria/app"
	"github.com/rafaelbenvenuti/nymeria/app/models"
	"github.com/revel/revel"
)

// Use this struct for deploy elements.
type DeployElement struct {
	Deploy         models.Deploy
	Durations      map[string]int64
	Statuses       []string
	DurationsTotal int64
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

	// Genereate Data for all Deploys.
	deployData := []DeployElement{}
	for _, deploy := range deploys {
		if len(deploy.Statuses) > 1 {
			durations := make(map[string]int64)
			previousStatus := deploy.Statuses[0]
			statusesMap := map[string]bool{}
			statusesSlice := []string{}
			for _, currentStatus := range deploy.Statuses[1:] {
			        statusesMap[previousStatus.Status]= true
				durations[previousStatus.Status] = currentStatus.Date.Sub(previousStatus.Date).Nanoseconds()/int64(time.Millisecond)
				previousStatus = currentStatus
			}
			for status, _ := range statusesMap {
					statusesSlice = append(statusesSlice, status)
			}
			deployData = append(deployData, DeployElement{
				Deploy:         deploy,
				Statuses:       statusesSlice,
				Durations:      durations,
				DurationsTotal: deploy.Statuses[len(deploy.Statuses)-1].Date.Sub(deploy.Statuses[0].Date).Nanoseconds()/int64(time.Millisecond),
			})
		}
	}

	// Generate Data for the Dashboard.
	dashboardData := make(map[string][]DeployElement)
	for _, deployElement := range deployData {
		dashboardData[deployElement.Deploy.Component] = append(dashboardData[deployElement.Deploy.Component], deployElement)
	}

	// Return 200 code and all deploy data.
	revel.INFO.Println("Deploy and Status data successfully retrieved.")
	c.Response.Status = 200
	return c.Render(dashboardData)
}
