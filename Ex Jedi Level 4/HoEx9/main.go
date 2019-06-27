package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      []string{`Women`, `Martinis`, `Shaken, not stirred`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	x["gold_finger"] = []string{`Gold Finger`, "Gold", "Golden Gun"}

	for i, v := range x {
		fmt.Println("Record: ", i)
		for j, z := range v {
			fmt.Printf("\t")
			fmt.Println("Index:", j, "Value:", z)
		}
	}

}
