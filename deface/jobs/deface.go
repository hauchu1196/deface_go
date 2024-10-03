package jobs

import (
	"deface/models"
	"deface/services"
	"log"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func StartJob(db *gorm.DB) {
	c := cron.New()

	c.AddFunc("@every 1h", func() {
		log.Println("Checking websites for defacement...")

		websites, err := models.GetAllWebsites(db)
		if err != nil {
			log.Printf("Error fetching websites from database: %v", err)
			return
		}

		for _, website := range websites {
			err := services.CheckWebsiteChanges(website.URL)
			if err != nil {
				log.Printf("Error checking website %s: %v", website.URL, err)
			}
		}
	})

	c.Start()
}
