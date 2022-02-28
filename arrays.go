package main

import "fmt"

func main() {
	months := [12]string{"Jan22", "Feb22", "Mar22", "Apr22", "May22", "Jun21", "Jul21", "Aug21", "Sep21", "Oct21", "Nov21", "Dec21"}
	distanceRunPerMonth := [12]float64{29.89, 79.65, 0, 0, 0, 55.71, 75.94, 68.80, 74.39, 61.57, 3.35, 15.08}
	for i, month := range months {
		fmt.Println(month, distanceRunPerMonth[i])
	}
}
