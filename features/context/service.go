package context

func GetContexts() []Context {
	return DbGetContexts()
}
func GetContext(id int) Context {
	return DbGetContext(id)
}
func GetContextsByEnvironmentId(environmentId int) []Context {
	return DbGetContextsByEnv(environmentId)
}
func AddContext(context Context) Context {
	return DbAddContext(context)
}
func UpdateContext(context Context) Context {
	return DbUpdateContext(context)
}
func DeleteContext(id int) Context {
	return DbDeleteContext(id)
}
