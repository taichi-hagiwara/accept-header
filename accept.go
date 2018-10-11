package accept

import (
	"strings"

	"github.com/pkg/errors"
)

// Accept is http Accept header.
type Accept map[float32][]ContentType

// ParseAccept parses Accept.
func ParseAccept(v string) (Accept, error) {
	d := make(map[float32][]ContentType)
	s := strings.Split(v, ",")
	for _, l := range s {
		c, err := ParseContentType(l)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parce Accept")
		}
		if _, ok := d[c.Q]; !ok {
			d[c.Q] = make([]ContentType, 0)
		}
		d[c.Q] = append(d[c.Q], c)
	}
	return d, nil
}

// First returns most matched ContentType.
func (a Accept) First(contentTypes ...string) *ContentType {
	var ct *ContentType
	q := float32(0)
ALL_BREAK:
	for _, s := range contentTypes {
		c, err := ParseContentType(s)
		if err != nil {
			for q2, l := range a {
				for _, t := range l {
					if t.Match(&c) && q < q2 {
						ct = &t
						q = q2
						if q == 1 {
							break ALL_BREAK
						}
					}
				}
			}
		}
	}
	return ct
}
