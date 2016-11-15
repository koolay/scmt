package parse

import (
	"encoding/json"
	"fmt"
	"strings"
)

/*

var inputJsonStr = `
{
    "data": {
        "info": {
            "age": 18,
            "fee": 12.01,
            "meta": [ "test1" , "test2"],
            "sums": [ 11 , 12],
            "parents": [{"fld1" : "val1"} , {"fld2" : "val2"}]
        },
        "greet":"hello world"
    }
}
`
*/
// convert apiResponse to swagger json schema
func ConvertResponseContentToJsonSchema(inputStr string) []byte {

	if strings.HasPrefix(inputStr, "[") && strings.HasSuffix(inputStr, "]") {

		fmt.Println("-----------------------------------------")
		inputArr := []interface{}{}
		err := json.Unmarshal([]byte(inputStr), &inputArr)
		if err != nil {
			panic(err)
		}

		items := map[string]interface{}{}
		outputObj := map[string]interface{}{"type": "array", "items": items}
		parseArray(inputArr, items)
		if result, err := json.Marshal(outputObj); err == nil {
			return result
		} else {
			panic(err)
		}

	} else if strings.HasPrefix(inputStr, "{") && strings.HasSuffix(inputStr, "}") {

		inputObj := map[string]interface{}{}
		err := json.Unmarshal([]byte(inputStr), &inputObj)
		if err != nil {
			panic(err)
		}

		schemaObj := map[string]interface{}{}
		outputObj := map[string]interface{}{"type": "object", "properties": schemaObj}
		parseMap(inputObj, schemaObj)
		if result, err := json.Marshal(outputObj); err == nil {
			return result

		} else {
			panic(err)
		}

	} else {
		return nil
	}
}

// parse object
func parseMap(aMap map[string]interface{}, out map[string]interface{}) {

	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			subSchema := map[string]interface{}{}
			out[key] = map[string]interface{}{"type": "object", "properties": subSchema}
			parseMap(val.(map[string]interface{}), subSchema)
		case []interface{}:
			items := map[string]interface{}{}
			out[key] = map[string]interface{}{"type": "array", "items": items}
			parseArray(val.([]interface{}), items)
		default:
			typeStr := getType(concreteVal)
			out[key] = map[string]string{"type": typeStr}
		}
	}
}

// parse array
func parseArray(anArray []interface{}, itemMap map[string]interface{}) {

	if len(anArray) > 0 {
		elem := anArray[0]
		switch concreteVal := elem.(type) {
		case map[string]interface{}:
			itemMap["type"] = "object"
			propsMap := map[string]interface{}{}
			itemMap["properties"] = propsMap
			parseMap(elem.(map[string]interface{}), propsMap)
		case []interface{}:
			tempItems := map[string]interface{}{}
			itemMap = map[string]interface{}{"type": "array", "items": tempItems}
			parseArray(elem.([]interface{}), tempItems)
		case string:
			itemMap["type"] = "string"
		case int, int64, float64:
			itemMap["type"] = "number"
		default:
			panic(fmt.Sprintf("unknown type: %s", concreteVal))
		}
	}
}

func getType(val interface{}) string {

	switch val.(type) {
	case string:
		return "string"
	case int, int32, int64, float64, float32:
		return "number"
	case bool:
		return "boolean"
	default:
		return "unknown"
	}

}
