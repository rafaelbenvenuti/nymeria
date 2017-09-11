package app

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rafaelbenvenuti/nymeria/app/models"
	"github.com/revel/revel"
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
	uri := revel.BasePath + "/" + revel.Config.StringDefault("db.uri", "test.db")

	seedDatabase := false
	if _, err := os.Stat(revel.BasePath + "/" + revel.Config.StringDefault("db.uri", "test.db")); os.IsNotExist(err) {
		seedDatabase = true
	}

	revel.INFO.Println("Starting database connection...")

	Database, err = gorm.Open(driver, uri)
	if err != nil {
		revel.ERROR.Println("FATAL", err)
		panic(err)
	}
	revel.INFO.Println("Database connected!")

	Database.LogMode(true)
	Database.AutoMigrate(&models.Deploy{}, &models.Status{})

	if seedDatabase {
		SeedDatabase()
	}

}

func SeedDatabase() {
	seeds := []models.Deploy{
		models.Deploy{
			Accountable: "dev-team-1",
			Component:   "frontend",
			Version:     "v1.0.0",
			Statuses: []models.Status{
				models.Status{
					Status: "start",
					Date:   time.Date(2017, time.September, 1, 12, 0, 0, 0, time.Local),
				},
				models.Status{
					Status: "build",
					Date:   time.Date(2017, time.September, 1, 12, 5, 0, 0, time.Local),
				},
				models.Status{
					Status: "test",
					Date:   time.Date(2017, time.September, 1, 12, 15, 0, 0, time.Local),
				},
				models.Status{
					Status: "deliver",
					Date:   time.Date(2017, time.September, 1, 12, 30, 0, 0, time.Local),
				},
				models.Status{
					Status: "end",
					Date:   time.Date(2017, time.September, 1, 13, 0, 0, 0, time.Local),
				},
			},
		},
		models.Deploy{
			Accountable: "dev-team-1",
			Component:   "frontend",
			Version:     "v1.0.1",
			Statuses: []models.Status{
				models.Status{
					Status: "start",
					Date:   time.Date(2017, time.September, 1, 18, 0, 0, 0, time.Local),
				},
				models.Status{
					Status: "build",
					Date:   time.Date(2017, time.September, 1, 18, 5, 0, 0, time.Local),
				},
				models.Status{
					Status: "test",
					Date:   time.Date(2017, time.September, 1, 18, 15, 0, 0, time.Local),
				},
				models.Status{
					Status: "deliver",
					Date:   time.Date(2017, time.September, 1, 18, 45, 0, 0, time.Local),
				},
				models.Status{
					Status: "end",
					Date:   time.Date(2017, time.September, 1, 19, 10, 0, 0, time.Local),
				},
			},
		},
		models.Deploy{
			Accountable: "dev-team-1",
			Component:   "frontend",
			Version:     "v1.0.2",
			Statuses: []models.Status{
				models.Status{
					Status: "start",
					Date:   time.Date(2017, time.September, 2, 9, 0, 0, 0, time.Local),
				},
				models.Status{
					Status: "build",
					Date:   time.Date(2017, time.September, 2, 9, 5, 0, 0, time.Local),
				},
				models.Status{
					Status: "test",
					Date:   time.Date(2017, time.September, 2, 9, 15, 0, 0, time.Local),
				},
				models.Status{
					Status: "deliver",
					Date:   time.Date(2017, time.September, 2, 9, 50, 0, 0, time.Local),
				},
				models.Status{
					Status: "end",
					Date:   time.Date(2017, time.September, 2, 10, 20, 0, 0, time.Local),
				},
			},
		},
		models.Deploy{
			Accountable: "dev-team-1",
			Component:   "frontend",
			Version:     "v1.0.3",
			Statuses: []models.Status{
				models.Status{
					Status: "start",
					Date:   time.Date(2017, time.September, 2, 13, 0, 0, 0, time.Local),
				},
				models.Status{
					Status: "build",
					Date:   time.Date(2017, time.September, 2, 13, 5, 0, 0, time.Local),
				},
				models.Status{
					Status: "test",
					Date:   time.Date(2017, time.September, 2, 13, 15, 0, 0, time.Local),
				},
				models.Status{
					Status: "deliver",
					Date:   time.Date(2017, time.September, 2, 13, 45, 0, 0, time.Local),
				},
				models.Status{
					Status: "end",
					Date:   time.Date(2017, time.September, 2, 14, 30, 0, 0, time.Local),
				},
			},
		},
	}

	for _, seed := range seeds {
		Database.Create(&seed)
	}
}
