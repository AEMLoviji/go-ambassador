package main

import (
	"ambassador/src/database"
	"ambassador/src/models"

	"github.com/bxcodec/faker/v3"
)

func main() {
	// Call it inside application container to successfully establish a connection with DB
	// Example: docker exec -it <name_of_container> sh
	// and then call `go run src/commans/populateUsers.go`
	database.Connect()

	for i := 0; i < 30; i++ {
		ambassador := models.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}
		ambassador.SetPassword("1324")

		database.DB.Create(&ambassador)
	}
}
