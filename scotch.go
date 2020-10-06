package scotch

import (
	"strings"
)

var (
	symbolAll                = "*"
	resourceDivider          = ":"
	operationDivider         = ","
	resourceOperationDivider = "="
)

// Scope is scope string
type Scope string

// New Scope
func New(required string) Scope {
	return Scope(required)
}

// Match checks with given scope
func (s *Scope) Satisfy(given Scope) bool {
	if s.String() == given.String() {
		return true
	}

	return given.resources().contain(s.resources()) &&
		given.operations().contain(s.operations())
}

func (s *Scope) String() string {
	return string(*s)
}

func (s *Scope) resources() resources {
	ss := strings.Split(s.String(), resourceOperationDivider)

	resourcesPart := ss[0]
	return newResources(strings.Split(resourcesPart, resourceDivider)...)
}

type resources []string

func newResources(rs ...string) resources {
	return rs
}

func (rs resources) contain(compare resources) bool {
	for i, r := range rs {
		if r == symbolAll {
			// regard all lasting allowed if last element is all
			if len(rs) == i+1 {
				return true
			}
			continue
		}

		// check for out of index
		if len(compare) < i+1 {
			return false
		}

		if strings.HasSuffix(r, symbolAll) {
			prefix := strings.Split(r, symbolAll)[0]
			if strings.HasPrefix(compare[i], prefix) {
				continue
			}
			return false
		}

		if r != compare[i] {
			return false
		}
	}

	return len(rs) == len(compare)
}

func (s *Scope) operations() operations {
	opsPart := strings.Split(s.String(), resourceOperationDivider)
	if len(opsPart) != 2 {
		return newOperations(symbolAll)
	}
	return newOperations(strings.Split(opsPart[1], operationDivider)...)
}

type operations []string

func newOperations(ops ...string) operations {
	return ops
}

func (os operations) all() bool {
	return len(os) == 1 && os[0] == symbolAll
}

func (os operations) contain(compare operations) bool {
	for _, o := range os {
		if o == symbolAll {
			return true
		}

		for _, c := range compare {
			if c == o {
				return true
			}
		}
	}

	return false
}
