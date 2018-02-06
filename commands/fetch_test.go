package commands_test

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/kiln/commands"
	"net/http/httptest"
	"net/http"
	"fmt"
	"os"
	"path/filepath"
	"gopkg.in/yaml.v2"
	"strings"
)

var _ = Describe("bake", func() {
	var (
		someReleasesDirectory  string
		tmpDir                 string
		server *httptest.Server

		fetch = commands.NewFetch()
		files  = map[string][]byte{
			"release1": []byte(`RELEASE1`),
			"release2": []byte(`RELEASE2`),
		}
	)

	BeforeEach(func() {
		var err error
		tmpDir, err = ioutil.TempDir("", "command-test")
		Expect(err).NotTo(HaveOccurred())

		someReleasesDirectory, err = ioutil.TempDir(tmpDir, "")
		Expect(err).NotTo(HaveOccurred())

	})



	Context("when not all parameters specified", func() {
		It("fetches the releases", func() {
			err := fetch.Execute([]string{
				"--releases-file", "releases.yml",
			})

			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal(`missing required flag "--releases-directory"`))
		})
	})

	Context("when the yaml is invalid", func() {
		BeforeEach(func() {
			err := ioutil.WriteFile("releases.yml",[]byte(`{`), 0755)
			Expect(err).NotTo(HaveOccurred())
		})

		It("fetches the releases", func() {
			err := fetch.Execute([]string{
				"--releases-file", "releases.yml",
				"--releases-directory", someReleasesDirectory,
			})

			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("releases.yml"))

		})
	})

	Context("when the release file does not exist", func() {
		It("fetches the releases", func() {
			err := fetch.Execute([]string{
				"--releases-file", "some-missing-file.yml",
				"--releases-directory", someReleasesDirectory,
			})

			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("some-missing-file.yml"))
		})
	})

	Context("when all parameters specified and yaml is valid, but cannot access URL(s)", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.NotFoundHandler())
			Expect(server.URL).NotTo(BeEmpty())
			err := ioutil.WriteFile("releases.yml",[]byte(fmt.Sprintf(`---
release1: %s/release1
release2: %s/release2`,server.URL,server.URL)),
				0755)
			Expect(err).NotTo(HaveOccurred())
		})
		AfterEach(func() {
			server.Close()
		})

		It("returns an error", func() {
			err := fetch.Execute([]string{
				"--releases-file", "releases.yml",
				"--releases-directory", someReleasesDirectory,
			})
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("release1"))

		})
	})

	newServer := func(files map[string][]byte) (*httptest.Server, []byte) {
		mux := http.NewServeMux()
		for name, contents := range files {
			c := contents
			mux.Handle("/"+name,http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request){
				w.Write(c)
			}))
		}
		server := httptest.NewServer(mux)

		m := make(map[string]string)
		for name := range files {
			m[name] = fmt.Sprintf("%s/%s", server.URL, name)
		}
		b, err := yaml.Marshal(m)
		Expect(err).To(BeNil())
		return server, b
	}

	Context("when all parameters specified, yaml is valid, and URL are valid", func() {
		BeforeEach(func() {
			var yamlBytes []byte
			server, yamlBytes = newServer(files)
			Expect(server.URL).NotTo(BeEmpty())
			Expect(ioutil.WriteFile("releases.yml", yamlBytes,0755)).NotTo(HaveOccurred())
		})
		AfterEach(func() {
			server.Close()
		})

		It("fetches the releases", func() {
			err := fetch.Execute([]string{
				"--releases-file", "releases.yml",
				"--releases-directory", someReleasesDirectory,
			})
			Expect(err).To(BeNil())

			var fis []os.FileInfo
			fis, err = ioutil.ReadDir(someReleasesDirectory)
			Expect(err).To(BeNil())
			Expect(len(fis)).To(Equal(len(files)))
			for _, fi:= range fis {
				b, err := ioutil.ReadFile(filepath.Join(someReleasesDirectory,fi.Name()))
				Expect(err).To(BeNil())
				fileKey := strings.TrimSuffix(fi.Name(), ".tgz")
				Expect(b).To(Equal(files[fileKey]))
			}

		})
	})
})