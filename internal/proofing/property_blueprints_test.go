package proofing_test

import (
	"errors"

	"github.com/pivotal-cf/kiln/internal/proofing"
	yaml "gopkg.in/yaml.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PropertyBlueprints", func() {
	var productTemplate proofing.ProductTemplate

	BeforeEach(func() {
		var err error
		productTemplate, err = proofing.Parse("fixtures/property_blueprints.yml")
		Expect(err).NotTo(HaveOccurred())
	})

	It("parses the different types", func() {
		Expect(productTemplate.PropertyBlueprints[0]).To(BeAssignableToTypeOf(proofing.SimplePropertyBlueprint{}))
		Expect(productTemplate.PropertyBlueprints[1]).To(BeAssignableToTypeOf(proofing.SelectorPropertyBlueprint{}))
		Expect(productTemplate.PropertyBlueprints[2]).To(BeAssignableToTypeOf(proofing.CollectionPropertyBlueprint{}))
	})

	Context("failure cases", func() {
		Context("when the YAML cannot be unmarshalled", func() {
			It("returns an error", func() {
				propertyBlueprints := proofing.PropertyBlueprints([]proofing.PropertyBlueprint{})

				err := propertyBlueprints.UnmarshalYAML(func(v interface{}) error {
					return errors.New("unmarshal failed")
				})
				Expect(err).To(MatchError("unmarshal failed"))
			})
		})

		Context("when the YAML contains unknown fields", func() {
			It("returns an error", func() {
				propertyBlueprints := proofing.PropertyBlueprints([]proofing.PropertyBlueprint{})

				err := propertyBlueprints.UnmarshalYAML(func(v interface{}) error {
					return yaml.Unmarshal([]byte(`[foo: bar]`), v)
				})
				Expect(err).To(MatchError(ContainSubstring("field foo not found")))
			})
		})
	})
})