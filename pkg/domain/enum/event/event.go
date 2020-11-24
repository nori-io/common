package event

type EventName string

const (
	NoriPluginsInited  EventName = "nori/plugins/inited"
	NoriPluginsStarted EventName = "nori/plugins/started"
	NoriPluginsStopped EventName = "nori/plugins/stopped"
	NoriPlugins        EventName = "nori/plugins/*"
)
