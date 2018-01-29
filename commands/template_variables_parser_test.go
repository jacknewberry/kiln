package commands_test

import (
	"io/ioutil"
	"os"

	"github.com/pivotal-cf/kiln/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TemplateVariablesParser", func() {
	Describe("Execute", func() {
		var (
			parser commands.TemplateVariablesParser
			path   string
		)

		BeforeEach(func() {
			parser = commands.NewTemplateVariableParser()

			contents := `---
key-1:
  key-2:
  - value-1
  - value-2
key-3: value-3
`

			file, err := ioutil.TempFile("", "variables")
			Expect(err).NotTo(HaveOccurred())

			path = file.Name()

			_, err = file.WriteString(contents)
			Expect(err).NotTo(HaveOccurred())

			err = file.Close()
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			err := os.Remove(path)
			Expect(err).NotTo(HaveOccurred())
		})

		It("parses template variables from a collection of files", func() {
			variables, err := parser.Execute([]string{path}, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(variables).To(Equal(map[string]interface{}{
				"key-1": map[interface{}]interface{}{
					"key-2": []interface{}{
						"value-1",
						"value-2",
					},
				},
				"key-3": "value-3",
			}))
		})

		It("parses template variables from command-line arguments", func() {
			variables, err := parser.Execute(nil, []string{
				"key-1=value-1",
				"key-2=value-2",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(variables).To(Equal(map[string]interface{}{
				"key-1": "value-1",
				"key-2": "value-2",
			}))
		})

		Context("failure cases", func() {
			Context("when the variable file cannot be read", func() {
				It("returns an error", func() {
					_, err := parser.Execute([]string{"missing.yml"}, nil)
					Expect(err).To(MatchError(ContainSubstring("open missing.yml: no such file or directory")))
				})
			})

			Context("when the variable file contents cannot be unmarshalled", func() {
				It("returns an error", func() {
					err := ioutil.WriteFile(path, []byte("\t\t\t"), 0644)
					Expect(err).NotTo(HaveOccurred())

					_, err = parser.Execute([]string{path}, nil)
					Expect(err).To(MatchError("yaml: found character that cannot start any token"))
				})
			})

			Context("when the command-line variables are malformed", func() {
				It("returns an error", func() {
					_, err := parser.Execute(nil, []string{"garbage"})
					Expect(err).To(MatchError("could not parse variable \"garbage\": expected variable in \"key=value\" form"))
				})
			})
		})
	})
})