package main

import "fmt"

func main() {
	diego := person{
		firstName: "Diego",
		lastName:  "Surita",
		contact: contact{
			email:   "suritadev@gmail.com",
			zipCode: 87000123,
		},
	}

	fmt.Printf("%+v", diego)
}
