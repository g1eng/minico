package docker

import . "gopkg.in/check.v1"

func (s testSuite) TestQueryImageMatched(c *C) {
	err := s.provider.queryImageByName(s.fixtures.images, true, "dev")
	c.Check(err, IsNil)
}
func (s testSuite) TestQueryImageNotMatched(c *C) {
	err := s.provider.queryImageByName(s.fixtures.images, true, "that")
	c.Check(err, NotNil)
}
