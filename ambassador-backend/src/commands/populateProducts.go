package main

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func main() {
	// Call it inside application container to successfully establish a connection with DB
	// Example: docker exec -it <name_of_container> sh
	// and then call `go run src/commans/populateUsers.go`
	database.Connect()

	for i := 0; i < 30; i++ {
		produt := models.Product{
			Title:       faker.Username(),
			Description: faker.Username(),
			Image:       faker.URL(),
			Price:       float64(rand.Intn(90) + 10),
		}

		database.DB.Create(&produt)
	}
}
