package acceptance

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"net/http"

	yaml "gopkg.in/yaml.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http/httptest"
)

var _ = Describe("kiln fetch", func() {
	var (
		tmpDir                           string
		releasesFile string
		files  = map[string][]byte{
			"release1": []byte(`RELEASE1`),
			"release2": []byte(`RELEASE2`),
		}
		server *httptest.Server
	)

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

	BeforeEach(func() {
		var (
			err error
			filesYaml []byte
		)

		tmpDir, err = ioutil.TempDir("", "kiln-main-test")
		Expect(err).NotTo(HaveOccurred())

		releasesFile = filepath.Join(tmpDir, "releases.yml")

		server, filesYaml = newServer(files)

		Expect(ioutil.WriteFile(releasesFile, filesYaml, 0755)).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		server.Close()
		_ = os.RemoveAll(tmpDir)
	})

	It("downloads releases listed in releases-file into releases-directory", func() {

		command := exec.Command(pathToMain,
		  "fetch",
		  "--releases-file", releasesFile,
		  "--releases-directory", tmpDir,
		)
		output, err := command.CombinedOutput()
		Expect(err).To(BeNil(), string(output))
	})
})
