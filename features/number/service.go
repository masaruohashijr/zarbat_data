package number

func GetNumbers() []Number {
	return DbGetNumbers()
}
func GetNumber(id int) Number {
	return DbGetNumber(id)
}
func AddNumber(number Number) Number {
	return DbAddNumber(number)
}
func UpdateNumber(number Number) Number {
	return DbUpdateNumber(number)
}
func DeleteNumber(id int) Number {
	return DbDeleteNumber(id)
}
