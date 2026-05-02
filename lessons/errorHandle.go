package lessons

import (
	"errors"
	"math/rand"
	"researching-go/pkg/logger"
)

type client struct {
	name string
	usd  int
}

func pay(client *client, usd int) error {
	if client.usd <= 0 {
		return errors.New("insufficient funds")
	}

	client.usd -= usd
	return nil
}

func newClient(name string, usd int) *client {
	return &client{
		name: name,
		usd:  usd,
	}
}

func firstExample() {
	client := newClient("client1", 100)
	logger.Ptc("new client: ", client)
	if err := pay(client, 20); err != nil {
		logger.Ptc("operation is denied", err)
		return
	}
	logger.Ptc("operation accomplished")
	logger.Ptc("new client: ", client)
	if err := pay(client, 80); err != nil {
		logger.Ptc("operation is denied", err)
		return
	}
	logger.Ptc("operation accomplished")
	logger.Ptc("new client: ", client)
	if err := pay(client, 80); err != nil {
		logger.Ptc("operation is denied", err)
		return
	}
	logger.Ptc("operation accomplished")
}

type car struct {
	Armor int
}

func (car *car) onGas() (int, error) {
	if car.Armor <= 0 {
		return 0, errors.New("car is broken")
	}
	logger.Ptc("car is slow speed up")
	car.Armor -= 10
	return rand.Intn(200), nil
}

func secondExample() {
	c := car{
		Armor: 30,
	}

	for {
		logger.Ptc("car is started speed up: ", c)
		speed, err := c.onGas()
		if err != nil {
			logger.Ptc("car is stopped, because", err.Error())
			break
		}
		logger.Ptc("total speed and car after: ", speed, c)
	}

}

func ErrorHandleExample() {
	secondExample()
}
