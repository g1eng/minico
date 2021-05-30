package docker

import (
	. "gopkg.in/check.v1"
)

func (s *testSuite) TestQueryConByNameMatched(c *C) {
	err := s.provider.queryConByName(s.fixtures.con, true, "ok")
	c.Check(err, IsNil)
}

func (s *testSuite) TestQueryConByNameNotMatched(c *C) {
	err := s.provider.queryConByName(s.fixtures.con, true, "nok")
	c.Check(err, NotNil)
}

func (s *testSuite) TestConListCondition(c *C) {
	_, opts := s.provider.getConListCondition(true)
	c.Check(opts.All, Equals, true)
	_, opts = s.provider.getConListCondition(false)
	c.Check(opts.All, Equals, false)
}
