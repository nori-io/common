package plugin

import "github.com/nori-io/common/v3/pkg/domain/event"

type Eventable interface {
	Subscribe(emitter event.EventEmitter)
}
