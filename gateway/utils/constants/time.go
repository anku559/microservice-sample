package constants

import "fmt"

var DaysArray = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

var MonthsArray = [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

func init() {
	fmt.Println("========= Inside init (/utils/constants/time.go) ========= ")
	for _, day := range DaysArray {
		fmt.Println(day)
	}
}
