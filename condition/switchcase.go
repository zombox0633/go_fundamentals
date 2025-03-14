package condition

import "fmt"

func IsHexColor(colorName string) string {
	switch colorName {
	case "blue":
		return "#0000FF"
	case "green":
		return "#008000"
	case "pink":
		return "#FFC0CB"
	case "yellow":
		return "#FFFF00"
	default:
		fmt.Println("Unknown color 😾 : ", colorName)
		return "#000000"
	}
}

func ShowSwitchCase() {

	var color string

	for {
		fmt.Print("Enter color: ")
		fmt.Scanln(&color)

		hexColor := IsHexColor(color)

		fmt.Println("hex Color code: ", hexColor)
		// if hexColor == "#0000FF" {
		// 	fmt.Println("Blue is a good color 😺")
		// 	break
		// }

		var answer string
		fmt.Println("Do you want to continue? 😸 (yes/no) : ")
		fmt.Scanln(&answer)

		if answer == "yes" {
			fmt.Println("Goodbye 😿")
			break
		}
	}

}
