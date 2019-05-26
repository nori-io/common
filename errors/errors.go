package errors

import (
	"fmt"

	"github.com/nori-io/nori-common/meta"
)

type InterfaceAssertError struct {
	Interface meta.Interface
}

func (e InterfaceAssertError) Error() string {
	return fmt.Sprintf("can's assert interface to %s", e.Interface)
}
