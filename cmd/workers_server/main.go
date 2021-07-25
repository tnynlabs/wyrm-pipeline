package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tnynlabs/wyrm-pipeline/pkg/workers"
	"github.com/tnynlabs/wyrm/pkg/pipelines"
	"github.com/tnynlabs/wyrm/pkg/storage/postgres"
)

func main() {
	if devFlag := os.Getenv("WYRM_DEV"); devFlag == "1" {
		// Load environment variables from .env file
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file (error: %v)", err)
		}
	}

	db, err := postgres.GetFromEnv()
	if err != nil {
		log.Fatalln(err)
	}

	pipelineRepo := postgres.CreatePipelineRepository(db)
	pipelineService := pipelines.CreateService(pipelineRepo)

	workersServer := workers.NewServer(pipelineService)

	log.Println("Running pipeline workers server...")
	err = workers.RunServer(":5053", workersServer)
	if err != nil {
		log.Printf("Failed to run pipeline workers server (%v)\n", err)
	}
}
