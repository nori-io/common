package plugin

import "github.com/nori-io/common/v3/pkg/domain/event"

type Notifiable interface {
	Subscribe(emitter event.EventEmitter)
}
