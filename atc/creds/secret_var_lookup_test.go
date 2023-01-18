package creds_test

import (
	"github.com/concourse/concourse/atc/creds"
	"github.com/concourse/concourse/atc/creds/dummy"
	"github.com/concourse/concourse/vars"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("VariableLookupFromSecrets", func() {
	var (
		variables vars.Variables
	)

	BeforeEach(func() {
		secrets := dummy.NewSecretsFactory([]dummy.VarFlag{
			{
				Name:  "a",
				Value: "foo",
			},
		}).NewSecrets()
		variables = creds.NewVariables(secrets, "team", "pipeline", true)
	})

	Describe("Get", func() {
		It("retrieves a static var", func() {
			result, found, err := variables.Get(vars.Reference{Path: "a"})
			Expect(err).NotTo(HaveOccurred())
			Expect(found).To(BeTrue())

			Expect(result).To(Equal("foo"))
		})

		Context("when a path is missing", func() {
			It("errors", func() {
				_, found, _ := variables.Get(vars.Reference{Path: "b"})
				Expect(found).To(BeFalse())
			})
		})
	})
})
