package labtest

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"uniqueIndex" valid:"required,email~Email invalid"`
	Phone string `valid:"stringlength(10|10)~Phone invalid"`
	SuId  string `valid:"matches(^[MBD]\\d{7}$)~student_id pattern is not true"`
	Name  string `valid:"required~Name is required"`
}

func TestEmail(t *testing.T) {
	g := gomega.NewWithT(t)

	t.Run(`Email invalid`, func(t *testing.T) {
		user := User{
			Email: "dd",
			Phone: "0619500228",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Email invalid"))

	})
}

func TestPhone(t *testing.T) {
	g := gomega.NewWithT(t)

	t.Run(`Phone invalid`, func(t *testing.T) {
		user := User{
			Email: "dd@gmail.com",
			Phone: "12345",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Phone invalid"))
	})
}

func TestStudent(t *testing.T) {
	g := gomega.NewWithT(t)

	t.Run(`student_id pattern is not true`, func(t *testing.T) {
		user := User{
			Email: "dd@gmail.com",
			Phone: "0619500228",
			SuId:  "A6502456",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("student_id pattern is not true"))

	})
}

func TestValid(t *testing.T) {
	g := gomega.NewWithT(t)

	t.Run(`Email is valid`, func(t *testing.T) {
		user := User{
			Email: "dd@gmail.com",
			Phone: "0619500225",
			SuId:  "B6502456",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(gomega.BeTrue())
		g.Expect(err).To(gomega.BeNil())
	})
}

func TestName(t *testing.T) {
	g := gomega.NewWithT(t)

	t.Run(`Name is required`, func(t *testing.T) {
		user := User{
			Email: "dd@gmail.com",
			Phone: "0619500225",
			SuId:  "B6502456",
			Name:  "",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Name is required"))
	})
}
