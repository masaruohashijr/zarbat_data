package parameter

func GetParameters() []Parameter {
	return DbGetParameters()
}
func GetParameter(id int) Parameter {
	return DbGetParameter(id)
}
func AddParameter(parameter Parameter) Parameter {
	return DbAddParameter(parameter)
}
func UpdateParameter(parameter Parameter) Parameter {
	return DbUpdateParameter(parameter)
}
func DeleteParameter(id int) Parameter {
	return DbDeleteParameter(id)
}
