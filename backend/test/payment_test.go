package test

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/bestiesmile1845/test/backend/entity"
	. "github.com/onsi/gomega"
)

func TestPayment(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`Payment is positive`,func(t *testing.T) {
		payment := entity.Payment{
			PayerName : "smile",
			Date : time.Now() ,
			Amount: 1250,
		}
		ok,err := govalidator.ValidateStruct(payment)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	
}

func TestPaymentWorst(t *testing.T){
	g:= NewGomegaWithT(t)

	t.Run(`payment is negative`,func(t *testing.T){
		payment:=entity.Payment{
			PayerName: "",
			Date: time.Time{},
			Amount: 0,
		}
		ok, err := govalidator.ValidateStruct(payment)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).NotTo(Equal("Payer name is required"))
		g.Expect(err.Error()).NotTo(Equal("Date is required"))
		g.Expect(err.Error()).NotTo(Equal("Amount is required"))
	})
}