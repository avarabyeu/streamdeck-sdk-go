package sdk

import "github.com/valyala/fastjson"

func JsonStringValue(v *fastjson.Value, keys ...string) string {
	value := v.Get(keys...)

	if value == nil {
		return ""
	}

	strB, err := value.StringBytes()

	if err != nil {
		return ""
	}

	return string(strB)
}

func JsonCoordinates(v *fastjson.Value) Coordinates {
	return Coordinates{
		Row:    v.GetInt("coordinates", "row"),
		Column: v.GetInt("coordinates", "column"),
	}
}

func JsonSize(v *fastjson.Value) Size {
	return Size{
		Rows:    v.GetInt("size", "rows"),
		Columns: v.GetInt("size", "columns"),
	}
}

func JsonActionPayload(v *fastjson.Value) actionPayload {
	return actionPayload{
		Settings:    v.Get("settings"),
		Coordinates: JsonCoordinates(v),
	}
}

func JsonKeyPayload(v *fastjson.Value) KeyPayload {
	return KeyPayload{
		actionPayload:   JsonActionPayload(v),
		State:           v.GetInt("state"),
		IsInMultiAction: v.GetBool("isInMultiAction"),
	}
}
