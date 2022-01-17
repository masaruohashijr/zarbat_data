package run

func GetRuns() []Run {
	return DbGetRuns()
}
func GetRun(id int) Run {
	return DbGetRun(id)
}
func AddRun(run Run) Run {
	return DbAddRun(run)
}
func UpdateRun(run Run) Run {
	return DbUpdateRun(run)
}
func DeleteRun(id int) Run {
	return DbDeleteRun(id)
}
