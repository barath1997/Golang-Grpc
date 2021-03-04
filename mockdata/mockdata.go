package mockdata

import (
	"errors"
	"go-grpc-example/models"

	//"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// variables decalaration and initialization with mock data
var users = make([]models.User, 5)
var city []string = []string{"chennai", "mumbai", "delhi", "kolkata", "banglore"}
var names []string = []string{"saravanan", "kishore", "shekhar", "DaS", "sebastien"}
var height []float64 = []float64{5.1, 4.3, 6.3, 4.6, 5.6}

// GenerateData - function is used to generate mock data
func GenerateData() {

	// assigns value to users slice
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			users[i-1].Married = true
		} else {
			users[i-1].Married = false
		}

		users[i-1].ID = i
		users[i-1].City = city[i-1]
		users[i-1].Fname = names[i-1]
		users[i-1].Height = height[i-1]
		users[i-1].Phone = 986537286 + uint64(i)
	}

}

// GetData function returns user/users data
func GetData(userID []int) ([]models.User, error) {

	// variblae of type []models.User is created
	var data []models.User

	// user data is picked according to the id and assigned to data variable
	for _, id := range userID {
		for _, user := range users {
			if user.ID == id {
				data = append(data, user)
			}
		}
	}

	// error check if no data with the id is available
	if len(data) == 0 {
		log.Info().Msg("no data found for the specified id")
		return nil, errors.New("no data found for the specified id")
	}

	// user data is returned
	return data, nil

}
