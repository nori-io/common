package event

type EventName string

const (
	// core events
	NoriPluginsInited  EventName = "nori/plugins/inited"
	NoriPluginsStarted EventName = "nori/plugins/started"
	NoriPluginsStopped EventName = "nori/plugins/stopped"
	// patterns
	NoriPlugins EventName = "nori/plugins/*"
)
