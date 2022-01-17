package schedule

func GetSchedules() []Schedule {
	return DbGetSchedules()
}
func GetSchedule(id int) Schedule {
	return DbGetSchedule(id)
}
func AddSchedule(schedule Schedule) Schedule {
	return DbAddSchedule(schedule)
}
func UpdateSchedule(schedule Schedule) Schedule {
	return DbUpdateSchedule(schedule)
}
func DeleteSchedule(id int) Schedule {
	return DbDeleteSchedule(id)
}
