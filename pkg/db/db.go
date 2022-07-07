package db

import "github.com/foruscommunity/go-rest-api-example/pkg/models"

// Fake mock data for testing

func MockData() []models.User {
	var users []models.User

	users = append(users, models.User{Id: 0, Name: "Gustavo", Username: "gustavo", Bio: "forus.app/gustavo\n#ThisIsForus"})
	users = append(users, models.User{Id: 1, Name: "Eric", Username: "eric", Bio: "github.com/erictanaka1-zz"})
	users = append(users, models.User{Id: 2, Name: "Elon", Username: "elon", Bio: "Coins"})
	users = append(users, models.User{Id: 3, Name: "Steve", Username: "stevej", Bio: "Will not be like..."})

	return users
}
