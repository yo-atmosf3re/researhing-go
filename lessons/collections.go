package lessons

import (
	"researching-go/pkg/logger"
)

func CollectionsExample() {
	workWithMap()
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

	userArray := [3]user{{
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

func workWithSlice() {
	userArray := []user{}
	logger.Ptc("slice is created when we didnt type length of array inside square brackets", userArray, "- it's empty slice struct of User")
	// slice have ability to increase itself capacity when we put some value inside it. when length of slice (len) is increased, this slice increase itself capacity is 2 times from previous length - current length 4, capacity 5, +2 to length, capacity is increased *2 => capacity is 10 and length is 6.
	logger.Ptc("len, cap before: ", len(userArray), cap(userArray))

	userArray = append(userArray, user{
		Name:        "Alex",
		Age:         18,
		PhoneNumber: "+79998887766",
		IsClose:     true,
		Rating:      1.1,
	})
	logger.Ptc("append - added value of specific struct in this slice", userArray)
	logger.Ptc("len, cap after: ", len(userArray), cap(userArray))

	intSlice := make([]int, 0, 0)
	// make - it's function for slice, map, channels. for slice - first argument it's type of slice, second argument it's length of slice, third argument - capacity of slice.

	for i := 1; i < 10; i++ {
		intSlice = append(intSlice, i*i*i)
	}
	logger.Ptc("slice of int: ", len(intSlice), cap(intSlice), intSlice)
}

func workWithMap() {
	weather := map[string]int{
		"1 January": -30,
		"2 January": -28,
		"3 January": -15,
		"4 January": -12,
		"5 January": -23,
	}
	logger.Ptc("it's map, which declared with map-operator. inside square bracket pointed type of key, after bracket pointed type of value. inside braces pointed body of map.", weather)
	logger.Ptc("specific value of map from square brackets :", weather["1 January"])
	logger.Ptc("if we tried to get value which is not exist from nonexisted key that we got default value for specific type of value", weather["6 January"], "- key of January 6")

	sixthJanuaryValue, isExist := weather["6 January"]
	logger.Ptc("to determine exist value in map, we have ability to check second variables when we create new variable for read map-value. second variable have bool type, if first variable is not exist in map, than second variable will be false.", sixthJanuaryValue, isExist, "- existing value of map")

	weather["6 January"] = -15
	logger.Ptc("we can to add value with new key or change existing value from key which pointed inside square bracket. if key is exist, that old value will be replace, else be added like new value")

	logger.Ptc("for-cycle for map, when key - key of map, value - value of map. order of output is not same as declared.")
	for key, value := range weather { // "value := range weather" - it's copy, it's same as slice, array and e.g.
		logger.Ptc("key, value :", key, value)
	}

	for key, _ := range weather {
		weather[key] += 1
	}
	logger.Ptc("changed weather map: ", weather)

	for key := range weather {
		weather[key] += 10
	}
	logger.Ptc("changed weather map again: ", weather)

	names := make(map[string]string, 10)
	names["Alex"] = "Worker"
	names["John"] = "Programmer"
	names["Dave"] = "Gamer"
	logger.Ptc("declaring map with make-function, when first argument - classical map, second argument - capacity of map. init map is empty.")
}
