package utils

import (
	"encoding/json"
	"github.com/Dbinggo/HireSphere/server/common/log/zlog"
)

// StructToMap
//
//	@Description: struct to map
//	@param value
//	@return map[string]interface{}
func StructToMap(value interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	resJson, err := json.Marshal(value)
	if err != nil {
		zlog.Errorf("Json Marshal failed ,msg: %s", err.Error())
		return nil
	}
	err = json.Unmarshal(resJson, &m)
	if err != nil {
		zlog.Errorf("Json Unmarshal failed,msg : %s", err.Error())
		return nil
	}
	return m
}

// StructToJson
//
//	@Description: struct to json
//	@param value
//	@return string
//	@return error
func StructToJson(value interface{}) (string, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(data), err
}

// JsonToStruct
//
//	@Description: json to struct
//	@param str
//	@param value
//	@return error
func JsonToStruct(str string, value interface{}) error {
	return json.Unmarshal([]byte(str), value)
}
