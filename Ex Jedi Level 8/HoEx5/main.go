package main

import (
	"fmt"
	"sort"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []User
type ByLast []User

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	//fmt.Println(users)
	fmt.Println("------------------")

	// your code goes here
	/*sort.Slice(users, func(i, j int) bool { return users[i].Age < users[j].Age })
	fmt.Println("By Age:", users)

	fmt.Println("------------------")

	sort.Slice(users, func(i, j int) bool { return users[i].Last < users[j].Last })
	fmt.Println("By Last:", users)
	*/

	//fmt.Println("------------------")

	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		sort.Strings(v.Sayings)
		for _, val := range v.Sayings {
			fmt.Println("\t", val)
		}
	}
	fmt.Println("------------------")
	sort.Sort(ByAge(users))
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, val := range v.Sayings {
			//sort.Slice(v.Sayings, func(i, j int) bool { return v.Sayings[i] < v.Sayings[j] })
			fmt.Println("\t", val)
		}
	}
	fmt.Println("------------------")
	sort.Sort(ByLast(users))
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, val := range v.Sayings {
			//sort.Slice(v.Sayings, func(i, j int) bool { return v.Sayings[i] < v.Sayings[j] })
			fmt.Println("\t", val)
		}
	}
}
