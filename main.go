package main

func main() {
	createDB("test")

	insertUser(1, "test 1", 2001, 07, 07)
	insertUser(2, "test 2", 2001, 07, 07)
	insertUser(3, "test 3", 2001, 07, 07)
	insertUser(4, "test 4", 2001, 07, 07)
	insertUser(5, "test 5", 2001, 07, 07)

	updateUser(1, "HNQ", 2001, 07, 07)

	listUser()

	listUserById(1)
}
