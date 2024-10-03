package main

import (
	"deface/config"
	"deface/jobs"
	"deface/models"
	"deface/routes"
)

func main() {
	config.InitDB()
	models.MigrateDB(config.DB)
	jobs.StartJob(config.DB)
	r := routes.SetupRouter(config.DB)
	r.Run(":8080")
}
