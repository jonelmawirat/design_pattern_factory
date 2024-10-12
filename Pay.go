package pay

import (
    "errors"
)

type PaymentMethod interface {
    Pay(amount float32) string
}

const (
    Cash        = 1
    DebitCard   = 2
)

func GetPaymentMethod(m int) (PaymentMethod,error) {
    switch m {
        case 1:
            return &CashPM{},nil
        case 2:
            return &CreditCardPM{},nil
        default:
            return nil, errors.New("Not Implemented Yet")
    }
}

type CashPM struct {}
type CreditCardPM struct {}

func (c *CashPM) Pay(amount float32) string {
    return "paid using cash"
}

func (c *CreditCardPM) Pay(amount float32) string {
    return "paid using credit card"
}


