package docker

import "regexp"

// match is a name filter for regex name filtering in cosh.
// It is sometimes called `fuzzy matcher`
func (p *Provider) match(filter, name string) bool {
	return regexp.MustCompile(".*" + filter + ".*").MatchString(name)
}
