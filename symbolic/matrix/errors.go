package matrix

import "fmt"

type TypeError struct {
	Object interface{}
}

func (te TypeError) Error() string {
	return fmt.Sprintf(
		"objects of type %T are not currently recognized as matrices. if this is unexpected, then feel free to file an issue!",
		te.Object,
	)
}
