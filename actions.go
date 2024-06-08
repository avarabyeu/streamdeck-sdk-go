package sdk

var (
	actionDownMap = make(map[string]ActionHandlerBase)
	actionUpMap   = make(map[string]ActionHandlerBase)
)

func RegisterActionDown[T receivedPayload](actionName string, f ActionHandler[T]) {
	actionDownMap[actionName] = any(f).(ActionHandlerBase)
}
func RegisterActionUp[T receivedPayload](actionName string, f ActionHandler[T]) {
	actionUpMap[actionName] = any(f).(ActionHandlerBase)
}

func handleActionDown(action, context string, payload keyPayload, deviceId string) {
	if handler, ok := actionDownMap[action]; ok {
		handler(action, context, payload, deviceId)
	}
}
func handleActionUp(action, context string, payload keyPayload, deviceId string) {
	if handler, ok := actionUpMap[action]; ok {
		handler(action, context, payload, deviceId)
	}
}
