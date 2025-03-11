package dataStructure

func GetAnArrayWithLen10(initValue int) [10]int {
	// var array [length]int  // The size of array must be a constant
	var array [10]int
	for i := 0; i < 10; i++ {
		array[i] = initValue
	}
	return array
}
