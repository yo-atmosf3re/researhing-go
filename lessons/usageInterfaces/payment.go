package usageInterfaces

import (
	"fmt"
	"math/rand"
	"researching-go/pkg/logger"
)

type paymentMethods interface {
	Pay(amount int) int
	Cancel() int
	GetId() int
}

type payment struct {
	name       string
	id         int
	balance    int
	isCanceled bool
}

type paymentInfo struct {
	payments map[int]paymentMethods
}

func newPayment(name string, balance int) *payment {
	if name != "" {
		logger.Ptc("new payment created: ", name)
		return &payment{
			name:       name,
			id:         rand.Intn(15),
			balance:    balance,
			isCanceled: false,
		}
	}
	logger.Ptc("name is empty, exit with occurred error")
	return nil
}

func mewPaymentInfo() *paymentInfo {
	return &paymentInfo{payments: make(map[int]paymentMethods)}
}

func (payment *payment) Pay(amount int) int {
	if payment.balance <= -1000 {
		payment.Cancel()
		logger.Ptc("balance is so small, pay is canceled: ", payment.name)
		return payment.id
	}
	payment.balance -= amount
	logger.Ptc("pay is completed, balance updated")
	return payment.id
}

func (payment *payment) Cancel() int {

	if payment.isCanceled {
		logger.Ptc("operation already is canceled")
		return payment.id
	}
	payment.isCanceled = true
	logger.Ptc("operation is canceled")
	return payment.id
}

func (payment *payment) GetId() int {
	return payment.id
}

func (info *paymentInfo) AddInfo(pt paymentMethods) {
	if p, ok := pt.(*payment); ok {
		logger.Ptc("add payment info: ", p.name)
		info.payments[p.id] = pt
		return
	}
	logger.Ptc("add new info: error has occurred")
	return
}

func (info *paymentInfo) Info(id int) paymentMethods {
	if payment, ok := info.payments[id]; ok {
		logger.Ptc("data of payment by id: ", payment)
		return payment
	}
	logger.Ptc("payment not found")
	return nil
}

func (info *paymentInfo) AllInfo() map[int]paymentMethods {
	if len(info.payments) == 0 {
		logger.Ptc("no payments")
		return nil
	}
	logger.Ptc("all payments data: ", info.payments)
	return info.payments
}

var newPaymentInfo *paymentInfo = mewPaymentInfo()

func paymentTransaction(pt paymentMethods) {
	if p, ok := pt.(*payment); ok {
		logger.Ptc("current payment: ", p)
		logger.Ptc("new payment info is created: ", newPaymentInfo)

		p.Pay(11000)
		newPaymentInfo.AddInfo(p)
		newPaymentInfo.Info(p.id)

		pId := p.Pay(100)
		fmt.Printf("payment %d was denied, maybe trouble with balance", pId)

		return
	}
	logger.Ptc("error of payment, try again later")
	return
}

func PaymentExample() {
	paymentTransaction(newPayment("Bank", 10000))
	paymentTransaction(newPayment("Crypto", 200))
	paymentTransaction(newPayment("PayPal", 5))
	logger.Ptc("payment info, result :", newPaymentInfo)
	ps := convertMapToSlicePaymentInfo(newPaymentInfo)

	r1 := searchPaymentViaSlice(1, ps)
	r2 := searchPaymentViaSlice(2, ps)
	r3 := searchPaymentViaSlice(3, ps)
	r4 := searchPaymentViaSlice(4, ps)
	r5 := searchPaymentViaSlice(5, ps)

	logger.Ptc("founded slice-payment: ", r1)
	logger.Ptc("founded slice-payment: ", r2)
	logger.Ptc("founded slice-payment: ", r3)
	logger.Ptc("founded slice-payment: ", r4)
	logger.Ptc("founded slice-payment: ", r5)
}

func convertMapToSlicePaymentInfo(paymentInfo *paymentInfo) []paymentMethods {
	var paymentInfoSlice []paymentMethods
	for _, p := range paymentInfo.payments {
		paymentInfoSlice = append(paymentInfoSlice, p)
	}
	logger.Ptc("payment info: ", paymentInfoSlice)
	return paymentInfoSlice
}

func searchPaymentViaSlice(id int, ps []paymentMethods) paymentMethods {
	for _, p := range ps {
		if p.GetId() == id {
			return p
		}
	}
	logger.Ptc("slice: payment not found")
	return nil
}
