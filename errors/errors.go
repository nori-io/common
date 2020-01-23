package errors

import (
	"fmt"

	"github.com/nori-io/nori-common/v2/meta"
)

type InterfaceAssertError struct {
	Interface meta.Interface
}

func (e InterfaceAssertError) Error() string {
	return fmt.Sprintf("can't assert interface to %s", e.Interface)
}
