package step

func GetSteps() []Step {
	return DbGetSteps()
}
func GetStep(id int) Step {
	return DbGetStep(id)
}
func AddStep(step Step) Step {
	return DbAddStep(step)
}
func UpdateStep(step Step) Step {
	return DbUpdateStep(step)
}
func DeleteStep(id int) Step {
	return DbDeleteStep(id)
}
