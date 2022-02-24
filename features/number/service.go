package number

func GetNumbers() []PhoneNumber {
	return DbGetNumbers()
}
func GetNumber(id int) PhoneNumber {
	return DbGetNumber(id)
}
func AddNumber(number PhoneNumber) PhoneNumber {
	return DbAddNumber(number)
}
func UpdateNumber(number PhoneNumber) PhoneNumber {
	return DbUpdateNumber(number)
}
func DeleteNumber(id int) PhoneNumber {
	return DbDeleteNumber(id)
}
