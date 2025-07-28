package main

import "fmt"

func main() {
	days := []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday"}

	fmt.Println(days)
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and value is %v \n", index, day)
	}

	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 5 {
			rogueValue++
			continue
		}

		if rogueValue == 2 {
			goto lco
		}

		// if rogueValue == 5 {
		// 	break
		// }
		fmt.Println("Value is", rogueValue)
		rogueValue++
	}

lco:
	fmt.Print("Jumping at lco")

}
