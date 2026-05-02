package lessons

import "researching-go/pkg/logger"

// panic - it's specific error, e.g. error of compute, non-handle some case

func PanicExample() {
	defer func() {
		p := recover()
		if p != nil {
			logger.Ptc("panic is happened: ", p)
		}
	}()

	a := 1
	b := 0
	logger.Ptc("haha", a/b)
}
