package main

type User struct {
	Id         int
	Name       string
	Birth      int
	Created    int
	Updated_at int
}

type Point struct {
	User_id    int
	Points     int
	Max_points int
}

func main() {
	// today := time.Now().Day()

	// createDB("test")

	// insertUser(1, "test 1", 2001, 07, 07)
	// insertUser(2, "test 2", 2001, 07, 07)
	// insertUser(3, "test 3", 2001, 07, 07)
	// insertUser(4, "test 4", 2001, 07, 07)
	// insertUser(5, "test 5", 2001, 07, 07)

	// updateUser(1, "HNQ", 2001, 07, 07)

	// listUser()

	// listUserById(1)

	// for i := 1; i <= 100; i++ {
	// 	insertUser(i, "user "+strconv.Itoa(i), 2001, today, today)
	// }

	// transaction(1, 2000)

}
