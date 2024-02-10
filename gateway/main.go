package main

import (
	"fmt"
	UserTypes "github/anku559/microservice-sample/gateway/utils/types"
)

func main() {

	userAlex := UserTypes.User{
		FirstName:  "Alex",
		MiddleName: "Doe",
		LastName:   "Jr.",
	}

	alexAddress := UserTypes.UserAddress{
		AddressLine1: "House No. 123",
		AddressLine2: "Street ABC",
		AddressLine3: "XYZ Road",
		Landmark:     "Near QWERTY Tower",
		City:         "Ujjain",
		State:        "Madhya Pradesh",
		Country:      "India",
		Pincode:      "456001",
	}

	fmt.Printf("%+v\n", userAlex)
	fmt.Printf("%s\n", UserTypes.GetFullName(userAlex))
	fmt.Printf("%s\n", UserTypes.GetFullAddress(alexAddress))
}
