package accept

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// ContentType is a pair of MIME type and Q value.
type ContentType struct {
	Type    string
	Subtype string
	Q       float32
}

// MIME returns MIME type of ContentType.
func (c *ContentType) MIME() string {
	return fmt.Sprintf("%s/%s", c.Type, c.Subtype)
}

func (c *ContentType) String() string {
	if c.Q >= 1 {
		return c.MIME()
	}
	return fmt.Sprintf("%s;q=%v", c.MIME(), c.Q)
}

// Match returns matching result bitween two ContentTypes.
func (c *ContentType) Match(d *ContentType) bool {
	if c.Type != "*" && d.Type != "*" && c.Type != d.Type {
		return false
	}
	if c.Subtype != "*" && d.Subtype != "*" && c.Subtype != d.Subtype {
		return false
	}
	return true
}

// ParseContentType parses content type.
func ParseContentType(v string) (ContentType, error) {
	s := strings.Split(v, ";")
	c := ContentType{Q: 1}

	for i, j := range s {
		if i == 0 {
			s := strings.Split(j, "/")
			if len(s) != 2 {
				return c, errors.Errorf("invalid ContentType format: %s", v)
			}
			c.Type = s[0]
			c.Subtype = s[1]
		} else {
			if strings.HasPrefix(j, "q=") {
				q, err := strconv.ParseFloat(strings.TrimPrefix(j, "q="), 32)
				if err != nil || math.IsNaN(q) {
					return c, errors.Errorf("invalid ContentType format: %s", v)
				}
				if q < 0 {
					q = 0
				}
				if q > 1 {
					q = 1
				}
				c.Q = float32(q)
			}
		}
	}

	return c, nil
}
