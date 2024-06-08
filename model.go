package sdk

import "github.com/valyala/fastjson"

// Sent on socket open to identify the plugin/register it
type openMessage struct {
	Event string `json:"event"`
	UUID  string `json:"uuid"`
}

// Sent to the Stream Deck software as a generic event
type sentEvent struct {
	Event   string      `json:"event"`
	Context string      `json:"context,omitempty"`
	Payload interface{} `json:"payload"`
}

// Coordinates on the device
type Coordinates struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

type Size struct {
	Rows    int `json:"rows"`
	Columns int `json:"columns"`
}

// Payloads for sent events

type openUrlPayload struct {
	URL string `json:"url"`
}

type logMessagePayload struct {
	Message string `json:"message"`
}

type setTitlePayload struct {
	Title  string `json:"title"`
	Target int    `json:"target"`
}

type setStatePayload struct {
	State int `json:"state"`
}

type setImagePayload struct {
	Image  string `json:"image"`
	Target int    `json:"target"`
}

type switchToProfilePayload struct {
	Profile string `json:"profile"`
	Page    int    `json:"page"`
}

// Payloads for received events

type actionPayload struct {
	Settings    *fastjson.Value
	Coordinates Coordinates
}

type KeyPayload struct {
	actionPayload
	State           int
	IsInMultiAction bool
}

type DialButtonPayload struct {
	actionPayload
	Controller string
}

type DialRotatePayload struct {
	actionPayload
	Controller string
	// Ticks contains a value representative of the rotation that triggered the event.
	// For a dial, positive value signifies clockwise rotation, and a negative value signifies anticlockwise rotation.
	// The lowest position is set as 0, and highest position is set as 192.
	Ticks   int
	Pressed bool
}

type AppearPayload struct {
	KeyPayload
	Controller string
}

type ReceivedPayload interface {
	*fastjson.Value | AppearPayload | DialRotatePayload | DialButtonPayload | KeyPayload | actionPayload
}

// Events

type KeyDownEvent struct {
	Action   string
	Context  string
	Payload  KeyPayload
	DeviceId string
}

type KeyUpEvent struct {
	Action   string
	Context  string
	Payload  KeyPayload
	DeviceId string
}

type DialDownEvent struct {
	Action   string
	Context  string
	Payload  DialButtonPayload
	DeviceId string
}

type DialUpEvent struct {
	Action   string
	Context  string
	Payload  DialButtonPayload
	DeviceId string
}
type DialRotateEvent struct {
	Action   string
	Context  string
	Payload  DialRotatePayload
	DeviceId string
}

type WillAppearEvent struct {
	Action   string
	Context  string
	Payload  AppearPayload
	DeviceId string
}

type WillDisappearEvent struct {
	Action   string
	Context  string
	Payload  AppearPayload
	DeviceId string
}

type PropertyInspectorDidAppearEvent struct {
	Action   string
	Context  string
	DeviceId string
}

type PropertyInspectorDidDisappearEvent struct {
	Action   string
	Context  string
	DeviceId string
}

type DeviceConnectEvent struct {
	DeviceId string
	Name     string
	Size     Size
}

type DeviceDisconnectEvent struct {
	DeviceId string
}

type SendToPluginEvent struct {
	Action   string
	Context  string
	Payload  *fastjson.Value
	DeviceId string
}

type ReceiveSettingsEvent struct {
	Action   string
	Context  string
	DeviceId string
	actionPayload
	IsInMultiAction bool
}

type GlobalSettingsEvent struct {
	Settings *fastjson.Value
}

type ApplicationLaunchEvent struct {
	Application string
}

type ApplicationTerminateEvent struct {
	Application string
}
