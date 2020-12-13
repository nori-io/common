package plugin

import "github.com/nori-io/common/v4/pkg/domain/event"

type Notifiable interface {
	Subscribe(emitter event.EventEmitter)
}
