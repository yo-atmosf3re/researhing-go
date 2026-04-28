package lessons

import "researching-go/pkg/logger"

func PointersExample() {
	accessPointer()
}

func accessPointer() {
	pointer := new(10)
	var nilPointer *int

	getPointerValue(pointer)
	setPointerValue(pointer, 15)
	getPointerValue(pointer)

	getPointerValue(nilPointer)
	setPointerValue(nilPointer, 8)

}

func getPointerValue(pointer *int) {
	if pointer != nil {
		logger.Pt("address :", pointer, ", pointer value: ", *pointer)
	} else {
		logger.Pt("get: nil-pointer detected")
	}
}

func setPointerValue(pointer *int, value int) {
	if pointer != nil {
		*pointer = value
	} else {
		logger.Pt("set: nil-pointer detected")
	}
}
