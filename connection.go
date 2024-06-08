package sdk

import (
	"github.com/valyala/fastjson"
)

func reader() {
	var p fastjson.Parser

	var event, context, action, deviceId string

	var payload *fastjson.Value

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			close(closeSdk)
			return
		}

		v, err := p.ParseBytes(message)

		if err != nil {
			continue
		}

		event = JsonStringValue(v, CommonEvent)
		context = JsonStringValue(v, CommonContext)
		action = JsonStringValue(v, CommonAction)
		deviceId = JsonStringValue(v, CommonDevice)

		payload = v.Get(CommonPayload)

		switch event {
		case EventDialDown:
			handleEvent(EventDialDown, &DialDownEvent{
				action, context, dialButtonPayload{
					actionPayload: JsonActionPayload(payload),
					Controller:    JsonStringValue(payload, "controller"),
				}, deviceId,
			})
		case EventDialUp:
			handleEvent(EventDialUp, &DialUpEvent{
				action, context, dialButtonPayload{
					actionPayload: JsonActionPayload(payload),
					Controller:    JsonStringValue(payload, "controller"),
				}, deviceId,
			})
		case EventDialRotate:
			handleEvent(EventDialRotate, &DialRotateEvent{
				action, context, dialRotatePayload{
					actionPayload: JsonActionPayload(payload),
					Controller:    JsonStringValue(payload, "controller"),
					Ticks:         payload.GetInt("ticks"),
					Pressed:       payload.GetBool("pressed"),
				}, deviceId,
			})
		case EventPropertyInspectorDidAppear:
			handleEvent(EventPropertyInspectorDidAppear, &PropertyInspectorDidAppearEvent{action, context, deviceId})
		case EventPropertyInspectorDidDisappear:
			handleEvent(EventPropertyInspectorDidDisappear, &PropertyInspectorDidDisappearEvent{
				action, context, deviceId,
			})
		case EventKeyDown:
			handleEvent(EventKeyDown, &KeyDownEvent{
				action, context, JsonKeyPayload(payload), deviceId,
			})
		case EventKeyUp:
			handleEvent(EventKeyUp, &KeyUpEvent{
				action, context, JsonKeyPayload(payload), deviceId,
			})
		case EventWillAppear:
			handleEvent(EventWillAppear, &WillAppearEvent{
				action, context, appearPayload{
					keyPayload: JsonKeyPayload(payload),
					Controller: JsonStringValue(payload, "controller"),
				}, deviceId,
			})
		case EventWillDisappear:
			handleEvent(EventWillDisappear, &WillDisappearEvent{
				action, context, appearPayload{
					keyPayload: JsonKeyPayload(payload),
					Controller: JsonStringValue(payload, "controller"),
				}, deviceId,
			})
		case EventDeviceDidConnect:
			info := v.Get(CommonDeviceInfo)
			handleEvent(EventDeviceDidConnect, &DeviceConnectEvent{
				deviceId, JsonStringValue(info, "name"), JsonSize(info),
			})
		case EventDeviceDidDisconnect:
			handleEvent(EventDeviceDidDisconnect, &DeviceDisconnectEvent{deviceId})
		case EventSendToPlugin:
			handleEvent(EventSendToPlugin, &SendToPluginEvent{action, context, payload, deviceId})
		case EventDidReceiveSettings:
			handleEvent(EventDidReceiveSettings, &ReceiveSettingsEvent{
				action, context, deviceId, JsonActionPayload(payload),
				payload.GetBool("isInMultiAction"),
			})
		case EventDidReceiveGlobalSettings:
			handleEvent(EventDidReceiveGlobalSettings, &GlobalSettingsEvent{payload.Get("settings")})
		case EventApplicationDidLaunch:
			handleEvent(EventApplicationDidLaunch, &ApplicationLaunchEvent{JsonStringValue(payload, "application")})
		case EventApplicationDidTerminate:
			handleEvent(EventApplicationDidTerminate, &ApplicationTerminateEvent{JsonStringValue(payload, "application")})
		}
	}
}
