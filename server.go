package main

import (
	"gocron/config"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	app.Use(logger.New(
		logger.Config{
			Format: "✨ Latency: ${latency} Time:${time} Status: ${status} Path:${path} Query ${queryParams} \n",
		}))

	// Start the scheduler
	config.PipelineScheduler = gocron.NewScheduler(time.UTC)
	config.PipelineScheduler.StartAsync()

	// Load two existing schedules
	// config.PipelineScheduler.Tag("pipelines", "1").CronWithSeconds("*/5 * * * * *").Do(mytask, "1", "Africa/Johannesburg")
	// config.PipelineScheduler.Tag("pipelines", "2").CronWithSeconds("*/5 * * * * *").Do(mytask, "2", "Europe/London")
	config.PipelineScheduler.Tag("1").CronWithSeconds("*/5 * * * * *").Do(mytask, "1", "Africa/Johannesburg")
	config.PipelineScheduler.Tag("2").CronWithSeconds("*/5 * * * * *").Do(mytask, "2", "Europe/London")
	log.Println("Scheduler:", config.PipelineScheduler.IsRunning(), config.PipelineScheduler.Len())

	app.Post("/update/:nodeid", func(c *fiber.Ctx) error {

		NodeID := string(c.Params("nodeid"))
		Timezone := string(c.Query("timezone"))

		log.Println("Update: ", NodeID, Timezone)

		// Remove by tag to update
		config.PipelineScheduler.RemoveByTag(NodeID)
		log.Println("Scheduler count:", config.PipelineScheduler.Len())

		// Add new schedule
		// config.PipelineScheduler.Tag("pipelines", NodeID).CronWithSeconds("*/5 * * * * *").Do(mytask, NodeID, Timezone)
		config.PipelineScheduler.Tag(NodeID).CronWithSeconds("*/5 * * * * *").Do(mytask, NodeID, Timezone)
		log.Println("Scheduler count:", config.PipelineScheduler.Len())

		return c.Status(http.StatusOK).JSON(fiber.Map{"r": "updated"})
	})

	log.Fatal(app.Listen("0.0.0.0:1234"))
}

func mytask(nodeID string, timezone string) {

	log.Println("Trigger me", nodeID, timezone)

}
