package commands

import (
	"io/ioutil"
	"fmt"

	"github.com/pivotal-cf/jhanda"
	yaml "gopkg.in/yaml.v2"
	"net/http"
	"os"
	"io"
	"path/filepath"
	"sync"
	"strings"
)

type Fetch struct {
	Options struct {
		ReleasesFile string `short:"rf" long:"releases-file" required:"true" description:"path to the releases file (a YAML map of release name to URL)"`
		ReleasesDirectory string `short:"rd" long:"releases-directory" required:"true" description:"path to the releases directory"`
	}
}

func (f Fetch) Usage() jhanda.Usage {
	return jhanda.Usage{
		Description:      "Downloads the files listed in releases-file into the releases-directory. The releases-directory can then be used by `kiln bake`",
		ShortDescription: "fetches releases",
		Flags:            f.Options,
	}
}


func NewFetch() Fetch {
	return Fetch{}
}

func (f Fetch) Execute(args []string) error {
	args, err := jhanda.Parse(&f.Options, args)
	if err != nil {
		return err
	}


	b, err := ioutil.ReadFile(f.Options.ReleasesFile)
	if err != nil {
		return fmt.Errorf("unable to read '%s': %v", f.Options.ReleasesFile, err)
	}

	var releasesMap map[string]string
	if err := yaml.Unmarshal(b, &releasesMap); err != nil {
		return fmt.Errorf("invalid YAML in '%s': %v", f.Options.ReleasesFile, err)
	}


	errCh := make(chan error, len(releasesMap))
	names := make(chan string, len(releasesMap))
	for name := range releasesMap {
		names <- name
	}
	close(names)

	const numWorkers = 20
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i:=0; i<numWorkers; i++ {
		go func() {
			for name := range names {
				url := releasesMap[name]
				resp, err := http.Get(url)
				if err != nil {
					errCh <- err
					continue
				}
				if resp.StatusCode != http.StatusOK {
					errCh <- fmt.Errorf("[%s] status code %d when accessing %s", name, resp.StatusCode, url)
					continue
				}

				f, err := os.Create(fmt.Sprintf("%s.tgz", filepath.Join(f.Options.ReleasesDirectory, name)))
				if err != nil {
					errCh <- err
					continue
				}
				if _, err := io.Copy(f, resp.Body); err != nil {
					errCh <- err
					continue
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	close(errCh)
	var me multiError
	for err := range errCh {
		me = append(me, err)
	}
	if len(me) != 0 {
		return me
	}
	return nil
}

type multiError []error
func (me multiError) Error() string {
	strs := make([]string, len(me))
	for i, err := range me {
		strs[i] = err.Error()
	}
	return strings.Join(strs, ",")
}