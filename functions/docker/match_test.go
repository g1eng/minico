package docker

import . "gopkg.in/check.v1"

func (s testSuite) TestMatchMatched(c *C) {
	query := "mini"
	target := "minico"
	c.Check(s.provider.match(query, target), Equals, true)
}

func (s testSuite) TestMatchNotMatched(c *C) {
	query := "mono"
	target := "minico"
	c.Check(s.provider.match(query, target), Equals, false)
}
