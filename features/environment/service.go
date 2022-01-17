package environment

func GetEnvironments() []Environment {
	return DbGetEnvironments()
}
func GetEnvironment(id int) Environment {
	return DbGetEnvironment(id)
}
func AddEnvironment(environment Environment) Environment {
	return DbAddEnvironment(environment)
}
func UpdateEnvironment(environment Environment) Environment {
	return DbUpdateEnvironment(environment)
}
func DeleteEnvironment(id int) Environment {
	return DbDeleteEnvironment(id)
}
