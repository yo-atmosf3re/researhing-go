package lessons

import "researching-go/pkg/logger"

func CollectionsExample() {
	baseWorkWithStaticArray()
}

func baseWorkWithStaticArray() {
	classicalArray := [5]int{1, 2, 3, 4, 5}
	logger.Ptc("it is classical array with fixed length type of int", classicalArray)
	var sameClassicalArray = [5]int{1, 4, 8, 12, 5}
	logger.Ptc("it is same", sameClassicalArray)
	var emptyArray [5]int
	logger.Ptc("it is empty array, inside array will be place 0 value of 5 time", emptyArray)
	var specificArrayValue int = classicalArray[1]
	logger.Ptc("it is specific array value, which equal integer 2, because numerical value position in array is started at 0, as well as JS or others program languages", specificArrayValue)
	sameClassicalArray[2] += 8
	logger.Ptc("we can change specific array value when we same access than value from [] from number of index", sameClassicalArray)
	//tryAccessUndefinedValue := classicalArray[5]
	logger.Ptc("when we try to access index which is not exist, we got error - index out of bounds. Programs will not be compiled. Ide goland will be shown where we doing mistake")
	// array which have fixed length - it's static array, e.g - classicalArray, sameClassicalArray

	arrayForExample := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for currentValue := range arrayForExample {
		if currentValue%2 == 0 {
			logger.Ptc("get even integer", currentValue)
		}
	}

	for index := 0; index < len(arrayForExample); index++ {
		if arrayForExample[index]%2 != 0 {
			logger.Ptc("get odd integer", arrayForExample[index])
		}
	}

	userArray := [3]User{{
		Name:        "Alex",
		Age:         25,
		PhoneNumber: "+79999992277",
		IsClose:     true,
		Rating:      2.0,
	}, {
		Name:        "John",
		Age:         35,
		PhoneNumber: "+39998888822",
		IsClose:     false,
		Rating:      9.9,
	}, {
		Name:        "Dave",
		Age:         18,
		PhoneNumber: "+98887776633",
		IsClose:     true,
		Rating:      4.3,
	}}
	logger.Ptc("create array from struct", userArray)

	for _, currentArray := range userArray {
		if currentArray.Rating+2.0 < 10 {
			currentArray.Rating += 2.0
		}
	}
	logger.Ptc("array after changes. changes is not exist, because we did change with copy array", userArray)

	for index := range userArray {
		if userArray[index].Rating+2.0 < 10 {
			userArray[index].Rating += 2.0
		}
	}
	logger.Ptc("array after changes. changes is exist, because we did change with original array", userArray)

	// for i := 0; i < len(userArray); i++ - classical and full for cycle. inside cycle for need to work with original array for changes apply.
	// for index; user := range userArray - work with range of array, when index - index of array and user - current value of iteration to array, it's copy of original array. need to work original array.
	// fot index := range userArray - work with range of array, when index - index of array. inside cycle for need to work with original array for changes apply.
	// for { } - cycle without arguments, is infinity until we did stop it
}
