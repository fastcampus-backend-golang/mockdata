package data

import (
	"fmt"
	"math/rand"
)

func Generate(dataType string) string {
	switch dataType {
	case TYPE_NAME:
		return generateName()
	case TYPE_DATE:
		return generateDate()
	case TYPE_ADDRESS:
		return generateAddress()
	case TYPE_PHONE:
		return generatePhone()
	}

	return ""
}

// generateName generates a random name from data
func generateName() string {
	nameLen := len(Name)

	return Name[rand.Intn(nameLen)]
}

// generateDate generates a random date with format YYYY-MM-DD
func generateDate() string {
	year := rand.Intn(100) + 1950
	month := rand.Intn(12) + 1
	day := rand.Intn(28) + 1

	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

// generateAddress generates a random address from combination of street, random number, and city
func generateAddress() string {
	streetLen := len(Address["street"])
	cityLen := len(Address["city"])

	return fmt.Sprintf("Jl. %s No. %d, %s", Address["street"][rand.Intn(streetLen)], rand.Intn(100), Address["city"][rand.Intn(cityLen)])
}

// generatePhone generates a random phone number with 12 digits limit
func generatePhone() string {
	return fmt.Sprintf("08%d", rand.Intn(10000000000))
}
