package containerusers

import "fmt"

func SaysImUsername(users string) {
	toText("My Username:", users)
}

func SaysMyAgeUse(age string) {
	toText("My Age", age)
}

func toText(txt1 string, txt2 string)  {
	fmt.Println(txt1, txt2)
}