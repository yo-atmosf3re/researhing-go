package lessons

import (
	"researching-go/pkg/logger"
)

func DeferExample() {
	simpleDefer()
}

func simpleDefer() { // LIFO - last in, first out - 1, 2, 5, 6, 4, 3
	logger.Pt(1)
	logger.Pt(2)
	defer logger.Pt(3)
	defer logger.Pt(4)
	logger.Pt(5)
	defer logger.Pt(6)
}
