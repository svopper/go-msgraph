package msgraph

import (
	"fmt"
	"strings"
)

// Lists represents multiple lists of a Site.
type Lists []List

func (l Lists) String() string {
	var lists = make([]string, len(l))
	for i, list := range l {
		lists[i] = list.Name
	}
	return fmt.Sprintf("Lists(%v)", strings.Join(lists, ", "))
}
