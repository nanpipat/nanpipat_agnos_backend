package helpers

import (
	"encoding/base64"
	"encoding/json"

	"github.com/teamups-dev/food-delivery-api/core"
	"github.com/teamups-dev/food-delivery-api/core/utils"
	"github.com/thedevsaddam/gojsonq/v2"
)

func SetJSONValue(raw *json.RawMessage, field string, value interface{}) (*json.RawMessage, error) {
	var mapper core.Map
	err := json.Unmarshal(*raw, &mapper)
	if err != nil {
		return nil, err
	}
	mapper[field] = value

	b, err := json.Marshal(mapper)
	if err != nil {
		return nil, err
	}

	var newRawMessage json.RawMessage
	err = utils.Copy(&newRawMessage, b)
	return &newRawMessage, err
}

func GetJSONValue(raw *json.RawMessage, field string) interface{} {
	if raw == nil {
		return nil
	}

	return gojsonq.New().FromString(string(*raw)).Find(field)
}

func JSONToBase64NoPadding(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return base64.RawStdEncoding.EncodeToString(b)
}

func JSONToBase64URLNoPadding(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return base64.RawURLEncoding.EncodeToString(b)
}

func StructToJSON(value interface{}) (*json.RawMessage, error) {
	mapper, err := json.Marshal(value)
	if err != nil {

		return nil, err
	}
	var newRawMessage json.RawMessage
	err = utils.Copy(&newRawMessage, mapper)
	return &newRawMessage, err
}

func StructToMap(i interface{}) (map[string]interface{}, error) {
	// convert struct to JSON string
	b, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	// convert JSON string to map
	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
