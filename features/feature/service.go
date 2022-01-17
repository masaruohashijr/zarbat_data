package feature

func GetFeatures() []Feature {
	return DbGetFeatures()
}
func GetFeature(id int) Feature {
	return DbGetFeature(id)
}
func AddFeature(feature Feature) Feature {
	return DbAddFeature(feature)
}
func UpdateFeature(feature Feature) Feature {
	return DbUpdateFeature(feature)
}
func DeleteFeature(id int) Feature {
	return DbDeleteFeature(id)
}
