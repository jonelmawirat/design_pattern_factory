package pay

import (
    "testing"
    "strings"
)

func TestCreatePaymentMethodCash(t *testing.T){
    payment,err := GetPaymentMethod(Cash)
    if err != nil {
        t.Fatal("A payment method of cash must exist")
    }
    
    msg := payment.Pay(10.30)
    if !strings.Contains(msg,"paid using cash"){
        t.Errorf("The cash payment method message wasnt correct")
    }
    t.Log("LOG:",msg)

}


func TestGetPaymentMethodDebitCard(t *testing.T){
    payment,err := GetPaymentMethod(DebitCard)
    if err != nil {
        t.Fatal("A payment method of debit card must exist")
    }

    msg := payment.Pay(22.30)

    if !strings.Contains(msg,"paid using credit card") {
        t.Errorf("The debit card method message wasnt correct")
    }
    t.Log("LOG:",msg)
}

func TestGetPaymentMethodNonExisting(t *testing.T){
    _,err := GetPaymentMethod(20)
    if err == nil {
        t.Error("A payment method with ID 20 must return an error")
    }
}
