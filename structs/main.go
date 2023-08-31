package main

func main() {
	someone := person{
		firstName: "Diego",
		lastName:  "Surita",
		contact: contact{
			email:   "suritadev@gmail.com",
			zipCode: 87000123,
		},
	}

	someonePointer := &someone
	someonePointer.updateName("Alex")
	someone.print()
}
