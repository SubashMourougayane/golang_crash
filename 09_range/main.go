package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, id := range ids {
		fmt.Printf("ID : %d\n", id)
	}

	sum := 0

	for i := 0; i < len(ids); i++ {
		sum += ids[i]
	}

	for _, id := range ids {
		sum += id
	}

	fmt.Println(sum)

	emails := map[string]string{
		"bob":  "bob@yopmail.com",
		"brad": "brad@yopmail.com",
	}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

}
