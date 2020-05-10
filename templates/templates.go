package templates

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/lioneagle/goutil/src/logger"

	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
)

var tplmap map[string]string

var tpls = []string{
	"base", "chart", "header", "page", "routers",
}

func init() {
	tplmap = make(map[string]string)
	box := packr.NewBox(".")

	err := LoadTemplates(box)

	if err != nil {
		err = errors.Wrapf(err, "packr load templates failed")
		logger.Errorf("%+v", err)
		os.Exit(1)
	}
}

func GetTemplate(name string) (string, bool) {
	ret, ok := tplmap[name]
	return ret, ok
}

func LoadTemplates(loader http.FileSystem) error {
	return loadTemplates(tpls, loader)
}

func loadTemplates(files []string, loader http.FileSystem) error {
	for _, v := range files {
		name := v + ".html"
		file, err := loader.Open(name)
		if err != nil {
			return errors.Wrapf(err, "open templates \"%s\" failed", name)
		}

		content, err := ioutil.ReadAll(file)
		if err != nil {
			return errors.Wrapf(err, "read templates \"%s\" failed", name)
		}

		tplmap[v] = string(content)
	}

	return nil
}
