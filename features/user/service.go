package user

func GetUsers() []User {
	return DbGetUsers()
}
func GetUser(id int) User {
	return DbGetUser(id)
}
func AddUser(user User) User {
	return DbAddUser(user)
}
func UpdateUser(user User) User {
	return DbUpdateUser(user)
}
func DeleteUser(id int) User {
	return DbDeleteUser(id)
}
