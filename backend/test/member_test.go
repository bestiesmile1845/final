package test

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/bestiesmile1845/test/backend/entity"
)

func TestMember(t *testing.T) {
	// g := GomegaWithT(t)
	g := NewGomegaWithT(t)

	t.Run(`Member is positive`,func(t *testing.T) {
		member := entity.Member{
			Name:"Smile",
			Age: 15,
			MemberID: "CM12345678",
		}
		ok,err := govalidator.ValidateStruct(member)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
			

}

func TestMemberWorst(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Member is invalid`, func(t *testing.T) {
		member := entity.Member{
			Name:     "",           // Invalid: Name is required
			Age:      150,          // Invalid: Age is out of the range (0-120)
			MemberID: "B12345678",  // Invalid: MemberID does not start with "CM"
		}

		ok, err := govalidator.ValidateStruct(member)
		g.Expect(ok).NotTo(BeTrue()) // Expect validation to fail
		g.Expect(err).NotTo(BeNil()) // Expect an error

		// Check if the error message matches one of the validation issues
		g.Expect(err.Error()).NotTo(ContainSubstring("Name is required"))
		g.Expect(err.Error()).NotTo(ContainSubstring("Age must be between 0 and 120"))
		g.Expect(err.Error()).NotTo(ContainSubstring("MemberID must start with 'CM' and have 8 digits"))
	})
}
