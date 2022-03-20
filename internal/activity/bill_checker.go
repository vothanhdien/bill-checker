package activity

import (
	"strconv"
	"strings"
)

type CheckerStub struct {
	M map[string]int
}

func (c *CheckerStub) Check(cusCode string) (bool, error) {
	// response from input
	// <<anything>>_<<has bill after n attempts>>
	// foo_10: has bills after 10 attempts
	split := strings.Split(cusCode, "_")
	if len(split) != 2 {
		return true, nil
	}

	prefix := split[0]
	if count, ok := c.M[prefix]; ok {
		if count <= 1 {
			return true, nil
		}

		c.M[prefix] = count - 1
		return false, nil
	}

	// init
	t, _ := strconv.Atoi(split[1])
	c.M[prefix] = t
	return false, nil
}
