package client

import (
	"math"
	"time"

	. "gopkg.in/check.v1"
)

type UtilSuite struct{}

var _ = Suite(&UtilSuite{})

func (s *UtilSuite) TestAsMilliseconds(c *C) {
	d := time.Duration(30) * time.Millisecond
	c.Check(asMilliseconds(d, math.MaxInt32), Equals, 30)

	// approximately one year
	d = time.Duration(365) * time.Duration(24) * time.Hour
	c.Check(asMilliseconds(d, math.MaxInt32), Equals, math.MaxInt32)

	d = time.Duration(365) * time.Duration(24) * time.Hour
	c.Check(asMilliseconds(d, maxHeartBeat), Equals, maxHeartBeat)
}
