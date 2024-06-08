package sdk

// keyDownEventHandler is an event handler for KeyDownEvent events.
type keyDownEventHandler func(*KeyDownEvent)

// Type returns the event type for KeyDownEvent events.
func (eh keyDownEventHandler) Type() string {
	return EventKeyDown
}

// New returns a new instance of KeyDownEvent.
func (eh keyDownEventHandler) New() interface{} {
	return &KeyDownEvent{}
}

// Handle is the handler for KeyDownEvent events.
func (eh keyDownEventHandler) Handle(i interface{}) {
	if t, ok := i.(*KeyDownEvent); ok {
		eh(t)
	}
}

// keyDownEventHandler is an event handler for KeyUpEvent events.
type keyUpEventHandler func(*KeyUpEvent)

// Type returns the event type for KeyUpEvent events.
func (eh keyUpEventHandler) Type() string {
	return EventKeyUp
}

// New returns a new instance of KeyUpEvent.
func (eh keyUpEventHandler) New() interface{} {
	return &KeyUpEvent{}
}

// Handle is the handler for KeyUpEvent events.
func (eh keyUpEventHandler) Handle(i interface{}) {
	if t, ok := i.(*KeyUpEvent); ok {
		eh(t)
	}
}

// willAppearEventHandler is an event handler for WillAppearEvent events.
type willAppearEventHandler func(*WillAppearEvent)

// Type returns the event type for WillAppearEvent events.
func (eh willAppearEventHandler) Type() string {
	return EventWillAppear
}

// New returns a new instance of WillAppearEvent.
func (eh willAppearEventHandler) New() interface{} {
	return &WillAppearEvent{}
}

// Handle is the handler for WillAppearEvent events.
func (eh willAppearEventHandler) Handle(i interface{}) {
	if t, ok := i.(*WillAppearEvent); ok {
		eh(t)
	}
}

// willDisappearEventHandler is an event handler for WillDisappearEvent events.
type willDisappearEventHandler func(*WillDisappearEvent)

// Type returns the event type for WillDisappearEvent events.
func (eh willDisappearEventHandler) Type() string {
	return EventWillDisappear
}

// New returns a new instance of WillDisappearEvent.
func (eh willDisappearEventHandler) New() interface{} {
	return &WillDisappearEvent{}
}

// Handle is the handler for WillDisappearEvent events.
func (eh willDisappearEventHandler) Handle(i interface{}) {
	if t, ok := i.(*WillDisappearEvent); ok {
		eh(t)
	}
}

// deviceConnectEventHandler is an event handler for DeviceConnectEvent events.
type deviceConnectEventHandler func(*DeviceConnectEvent)

// Type returns the event type for DeviceConnectEvent events.
func (eh deviceConnectEventHandler) Type() string {
	return EventDeviceDidConnect
}

// New returns a new instance of DeviceConnectEvent.
func (eh deviceConnectEventHandler) New() interface{} {
	return &DeviceConnectEvent{}
}

// Handle is the handler for DeviceConnectEvent events.
func (eh deviceConnectEventHandler) Handle(i interface{}) {
	if t, ok := i.(*DeviceConnectEvent); ok {
		eh(t)
	}
}

// deviceDisconnectEventHandler is an event handler for DeviceDisconnectEvent events.
type deviceDisconnectEventHandler func(*DeviceDisconnectEvent)

// Type returns the event type for DeviceDisconnectEvent events.
func (eh deviceDisconnectEventHandler) Type() string {
	return EventDeviceDidDisconnect
}

// New returns a new instance of DeviceDisconnectEvent.
func (eh deviceDisconnectEventHandler) New() interface{} {
	return &DeviceDisconnectEvent{}
}

// Handle is the handler for DeviceDisconnectEvent events.
func (eh deviceDisconnectEventHandler) Handle(i interface{}) {
	if t, ok := i.(*DeviceDisconnectEvent); ok {
		eh(t)
	}
}

// sendToPluginEventHandler is an event handler for SendToPluginEvent events.
type sendToPluginEventHandler func(*SendToPluginEvent)

// Type returns the event type for SendToPluginEvent events.
func (eh sendToPluginEventHandler) Type() string {
	return EventSendToPlugin
}

// New returns a new instance of SendToPluginEvent.
func (eh sendToPluginEventHandler) New() interface{} {
	return &SendToPluginEvent{}
}

// Handle is the handler for SendToPluginEvent events.
func (eh sendToPluginEventHandler) Handle(i interface{}) {
	if t, ok := i.(*SendToPluginEvent); ok {
		eh(t)
	}
}

// globalSettingsEventHandler is an event handler for GlobalSettingsEvent events.
type globalSettingsEventHandler func(*GlobalSettingsEvent)

// Type returns the event type for GlobalSettingsEvent events.
func (eh globalSettingsEventHandler) Type() string {
	return EventDidReceiveGlobalSettings
}

// New returns a new instance of GlobalSettingsEvent.
func (eh globalSettingsEventHandler) New() interface{} {
	return &GlobalSettingsEvent{}
}

// Handle is the handler for GlobalSettingsEvent events.
func (eh globalSettingsEventHandler) Handle(i interface{}) {
	if t, ok := i.(*GlobalSettingsEvent); ok {
		eh(t)
	}
}

// receiveSettingsEventHandler is an event handler for ReceiveSettingsEvent events.
type receiveSettingsEventHandler func(*ReceiveSettingsEvent)

// Type returns the event type for ReceiveSettingsEvent events.
func (eh receiveSettingsEventHandler) Type() string {
	return EventDidReceiveSettings
}

// New returns a new instance of ReceiveSettingsEvent.
func (eh receiveSettingsEventHandler) New() interface{} {
	return &ReceiveSettingsEvent{}
}

// Handle is the handler for ReceiveSettingsEvent events.
func (eh receiveSettingsEventHandler) Handle(i interface{}) {
	if t, ok := i.(*ReceiveSettingsEvent); ok {
		eh(t)
	}
}

// applicationLaunchEventHandler is an event handler for ApplicationLaunchEvent events.
type applicationLaunchEventHandler func(*ApplicationLaunchEvent)

// Type returns the event type for ApplicationLaunchEvent events.
func (eh applicationLaunchEventHandler) Type() string {
	return EventApplicationDidLaunch
}

// New returns a new instance of ApplicationLaunchEvent.
func (eh applicationLaunchEventHandler) New() interface{} {
	return &ApplicationLaunchEvent{}
}

// Handle is the handler for ApplicationLaunchEvent events.
func (eh applicationLaunchEventHandler) Handle(i interface{}) {
	if t, ok := i.(*ApplicationLaunchEvent); ok {
		eh(t)
	}
}

// applicationTerminateEventHandler is an event handler for ApplicationTerminateEvent events.
type applicationTerminateEventHandler func(*ApplicationTerminateEvent)

// Type returns the event type for ApplicationTerminateEvent events.
func (eh applicationTerminateEventHandler) Type() string {
	return EventApplicationDidTerminate
}

// New returns a new instance of ApplicationTerminateEvent.
func (eh applicationTerminateEventHandler) New() interface{} {
	return &ApplicationTerminateEvent{}
}

// Handle is the handler for ApplicationTerminateEvent events.
func (eh applicationTerminateEventHandler) Handle(i interface{}) {
	if t, ok := i.(*ApplicationTerminateEvent); ok {
		eh(t)
	}
}

// propertyInspectorDidAppearEventHandler is an event handler for PropertyInspectorDidAppearEvent events.
type propertyInspectorDidAppearEventHandler func(*PropertyInspectorDidAppearEvent)

// Type returns the event type for PropertyInspectorDidAppearEvent events.
func (eh propertyInspectorDidAppearEventHandler) Type() string {
	return EventPropertyInspectorDidAppear
}

// New returns a new instance of PropertyInspectorDidAppearEvent.
func (eh propertyInspectorDidAppearEventHandler) New() interface{} {
	return &PropertyInspectorDidAppearEvent{}
}

// Handle is the handler for PropertyInspectorDidAppearEvent events.
func (eh propertyInspectorDidAppearEventHandler) Handle(i interface{}) {
	if t, ok := i.(*PropertyInspectorDidAppearEvent); ok {
		eh(t)
	}
}

// propertyInspectorDidDisappearEventHandler is an event handler for PropertyInspectorDidDisappearEvent events.
type propertyInspectorDidDisappearEventHandler func(*PropertyInspectorDidDisappearEvent)

// Type returns the event type for PropertyInspectorDidDisappearEvent events.
func (eh propertyInspectorDidDisappearEventHandler) Type() string {
	return EventPropertyInspectorDidDisappear
}

// New returns a new instance of PropertyInspectorDidDisappearEvent.
func (eh propertyInspectorDidDisappearEventHandler) New() interface{} {
	return &PropertyInspectorDidDisappearEvent{}
}

// Handle is the handler for PropertyInspectorDidDisappearEvent events.
func (eh propertyInspectorDidDisappearEventHandler) Handle(i interface{}) {
	if t, ok := i.(*PropertyInspectorDidDisappearEvent); ok {
		eh(t)
	}
}

// dialUpEventHandler is an event handler for DialUpEvent events.
type dialUpEventHandler func(*DialUpEvent)

// Type returns the event type for DialUpEvent events.
func (eh dialUpEventHandler) Type() string {
	return EventDialUp
}

// New returns a new instance of DialUpEvent.
func (eh dialUpEventHandler) New() interface{} {
	return &DialUpEvent{}
}

// Handle is the handler for DialUpEvent events.
func (eh dialUpEventHandler) Handle(i interface{}) {
	if t, ok := i.(*DialUpEvent); ok {
		eh(t)
	}
}

// dialDownEventHandler is an event handler for DialDownEvent events.
type dialDownEventHandler func(*DialDownEvent)

// Type returns the event type for DialDownEvent events.
func (eh dialDownEventHandler) Type() string {
	return EventDialDown
}

// New returns a new instance of DialDownEvent.
func (eh dialDownEventHandler) New() interface{} {
	return &DialDownEvent{}
}

// Handle is the handler for DialDownEvent events.
func (eh dialDownEventHandler) Handle(i interface{}) {
	if t, ok := i.(*DialDownEvent); ok {
		eh(t)
	}
}

// dialRotateEventHandler is an event handler for DialRotateEvent events.
type dialRotateEventHandler func(*DialRotateEvent)

// Type returns the event type for DialRotateEvent events.
func (eh dialRotateEventHandler) Type() string {
	return EventDialRotate
}

// New returns a new instance of DialRotateEvent.
func (eh dialRotateEventHandler) New() interface{} {
	return &DialRotateEvent{}
}

// Handle is the handler for DialRotateEvent events.
func (eh dialRotateEventHandler) Handle(i interface{}) {
	if t, ok := i.(*DialRotateEvent); ok {
		eh(t)
	}
}

func handlerForInterface(handler interface{}) EventHandler {
	switch v := handler.(type) {
	case func(interface{}):
		return interfaceEventHandler(v)
	case func(*KeyDownEvent):
		return keyDownEventHandler(v)
	case func(*KeyUpEvent):
		return keyUpEventHandler(v)
	case func(*WillAppearEvent):
		return willAppearEventHandler(v)
	case func(*WillDisappearEvent):
		return willDisappearEventHandler(v)
	case func(*DeviceConnectEvent):
		return deviceConnectEventHandler(v)
	case func(*DeviceDisconnectEvent):
		return deviceDisconnectEventHandler(v)
	case func(*SendToPluginEvent):
		return sendToPluginEventHandler(v)
	case func(*GlobalSettingsEvent):
		return globalSettingsEventHandler(v)
	case func(*ReceiveSettingsEvent):
		return receiveSettingsEventHandler(v)
	case func(*ApplicationLaunchEvent):
		return applicationLaunchEventHandler(v)
	case func(*ApplicationTerminateEvent):
		return applicationTerminateEventHandler(v)
	case func(*PropertyInspectorDidAppearEvent):
		return propertyInspectorDidAppearEventHandler(v)
	case func(*PropertyInspectorDidDisappearEvent):
		return propertyInspectorDidDisappearEventHandler(v)
	case func(*DialUpEvent):
		return dialUpEventHandler(v)
	case func(*DialDownEvent):
		return dialDownEventHandler(v)
	case func(*DialRotateEvent):
		return dialRotateEventHandler(v)
	}

	return nil
}
