package UserTypes

import "fmt"

type User struct {
	ID string

	FirstName  string
	MiddleName string
	LastName   string

	Email       string
	Username    string
	CountryCode string
	Phone       string
}

type UserAddress struct {
	UserId string

	FullName    string
	CountryCode string
	Phone       string

	AddressLine1 string
	AddressLine2 string
	AddressLine3 string
	Landmark     string
	Pincode      string
	City         string
	State        string
	Country      string
	AddressType  string // WORK, HOME

	GoogleLocation string
}

func GetFullName(user User) string {
	return fmt.Sprintf("%s %s %s", user.FirstName, user.MiddleName, user.LastName)
}

func GetFullAddress(userAddress UserAddress) string {
	return fmt.Sprintf("%s %s %s %s %s %s %s %s",
		userAddress.AddressLine1,
		userAddress.AddressLine2,
		userAddress.AddressLine3,
		userAddress.Landmark,
		userAddress.City,
		userAddress.State,
		userAddress.Country,
		userAddress.Pincode)
}
