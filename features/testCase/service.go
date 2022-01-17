package testCase

func GetTestCases() []TestCase {
	return DbGetTestCases()
}
func GetTestCase(id int) TestCase {
	return DbGetTestCase(id)
}
func AddTestCase(testCase TestCase) TestCase {
	return DbAddTestCase(testCase)
}
func UpdateTestCase(testCase TestCase) TestCase {
	return DbUpdateTestCase(testCase)
}
func DeleteTestCase(id int) TestCase {
	return DbDeleteTestCase(id)
}
