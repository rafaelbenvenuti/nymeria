package app

import (
	"time"

	"github.com/revel/revel"
	"github.com/jinzhu/gorm"
	"github.com/rafaelbenvenuti/nymeria/app/models"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string

	// Gorm Database
	Database *gorm.DB
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	revel.OnAppStart(InitDatabase)
}

var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

func InitDatabase() {
	var err error

	driver := revel.Config.StringDefault("db.driver", "sqlite3")
	uri := revel.BasePath + "/" + revel.Config.StringDefault("db.uri", "test.db" )

	revel.INFO.Println("Starting database connection...")

	Database, err = gorm.Open(driver, uri)
	if err != nil {
		revel.ERROR.Println("FATAL", err)
		panic(err)
	}
	revel.INFO.Println("Database connected!")

	Database.LogMode(true)
	Database.AutoMigrate(&models.Deploy{}, &models.Status{})

	seeds := []models.Deploy{
	  models.Deploy {
			Accountable: "dev-team-1",
		  Component: "frontend",
			Version: "v1.0.0",
			Statuses: []models.Status {
			  models.Status {
			    Status: "starting",
				  Date: time.Date(2017, time.September, 1, 12, 0, 0, 0, time.UTC),
			  },
			  models.Status {
			    Status: "ending",
				  Date: time.Date(2017, time.September, 1, 13, 0, 0, 0, time.UTC),
			  },
		  },
	  },
	}

  for _, seed := range seeds {
	  Database.Create(&seed)
	}

}
