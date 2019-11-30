package there

import "github.com/gobuffalo/here"

var cache = here.New()

func Dir(p string) (here.Info, error) {
	return cache.Dir(p)
}

func Package(p string) (here.Info, error) {
	return cache.Package(p)
}

func Current() (here.Info, error) {
	return cache.Current()
}
