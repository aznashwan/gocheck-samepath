package samepath_test

import (
	"runtime"
	"testing"

	gc "launchpad.net/gocheck"	

	sp "github.com/aznashwan/gocheck-samepath"
)

func Test(t *testing.T) {
	gc.TestingT(t)
}


type SamePathLinuxSuite struct {}
var _ = gc.Suite(&SamePathLinuxSuite{})

func (s * SamePathLinuxSuite) SetUpSuite(c *gc.C) {
	if runtime.GOOS != "linux" {
		c.Skip("Skipped Linux-intented SamePath tests.")
	}
}

func (s *SamePathLinuxSuite) TestSamePathLinuxBasic(c *gc.C) {
	c.Assert("/usr", sp.SamePath, "/usr")
	c.Assert("/usr\\share", sp.SamePath, "/usr/share")
	c.Assert("/usr/Share", gc.Not(sp.SamePath), "/usr/share")
}

func (s *SamePathLinuxSuite) TestSamePathLinuxSymlinks(c *gc.C) {
	c.Assert("/bin/rnano", sp.SamePath, "/bin/nano")
	c.Assert("/bin/rnano", gc.Not(sp.SamePath), "/bin/echo")
	c.Assert("/bin/echo", gc.Not(sp.SamePath), "/bin/rnano")
}


type SamePathWindowsSuite struct{}
var _ = gc.Suite(&SamePathWindowsSuite{})

func (s *SamePathWindowsSuite) SetUpSuite(c *gc.C) {
	if runtime.GOOS != "windows" {
		c.Skip("Skipped Windows-intented SamePath tests.")
	}
}

func (s *SamePathWindowsSuite) TestSamePathWindowsBasic(c *gc.C) {
	c.Assert("C:\\Users" sp.SamePath, "C:\\Users")
	c.Assert("C:\\Go\\src", sp.SamePath, "C:/Go/src")
	c.Assert("C:/Go/src", gc.Not(sp.SamePath), "C:/Go/pkg")
	c.Assert("C:/UseRs/hAroLd", sp.SamePath, "C:/Users/Harold")
}

func (s *SamePathSuite) TestSamePathShortenedPaths(c *gc.C) {
	c.Assert("C:/PROGRA~1", sp.SamePath, "C:/Program Files")
	c.Assert("C:/Program Files", sp.SamePath, "C:/PROGRA~1")
	c.Assert("C:/PROGRA~1", gc.Not(sp.SamePath), "C:/Users")
}