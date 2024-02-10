package main

import (
	"fmt"
	"gateway/utils/constants"
	UserTypes "gateway/utils/types"
)

func init() {
	fmt.Println("========= Inside init (main.go:10) ========= ")

	userAlex := UserTypes.User{
		FirstName:  "Alex",
		MiddleName: "Doe",
		LastName:   "Jr.",
	}

	fmt.Printf("%+v\n", userAlex)
	fmt.Printf("%s\n", UserTypes.GetFullName(userAlex))
}

func init() {
	fmt.Println("========= Inside init (main.go:23) ========= ")

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

	fmt.Printf("%s\n", UserTypes.GetFullAddress(alexAddress))
}

func main() {
	fmt.Println(constants.DaysArray[0])
	fmt.Println("========= End ========= ")
}
