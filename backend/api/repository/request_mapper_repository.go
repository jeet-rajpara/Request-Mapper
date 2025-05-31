package repository

import (
	er "request-mapper/error"
	"strings"
)

type RequestMapperRepository interface {
	MapRequest(inputJSON map[string]interface{}, requestMap map[string]string) (map[string]interface{}, error)
}

type requestMapperRepository struct{}

func NewRequestMapperRepository() RequestMapperRepository {
	return &requestMapperRepository{}
}

func (r *requestMapperRepository) MapRequest(inputJSON map[string]interface{}, requestMap map[string]string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for requestMapFieldKey, requestMapFieldValue := range requestMap {
		// if requestMapFieldValue is empty then set empty string value for that key in the result map
		if requestMapFieldValue == "" {
			setValueInResponse(result, requestMapFieldKey, "")
			continue
		}

		value, err := getValueFromInputJSON(inputJSON, requestMapFieldValue)
		if err != nil {
			return nil, err
		}
		setValueInResponse(result, requestMapFieldKey, value)
	}

	return result, nil
}

// getValueFromInputJSON retrieves a value from the input JSON using a dot-notation path
func getValueFromInputJSON(inputJSON map[string]interface{}, requestMapFieldValue string) (string, error) {
	keys := strings.Split(requestMapFieldValue, ".")

	current := inputJSON

	for i, key := range keys {
		if i == len(keys)-1 {
			value, exists := current[key]
			if !exists {
				return "", er.GenerateErrorCodeAndMessage(400, "requestMapFieldValue not found in inputJSON: "+requestMapFieldValue)
			}

			if value == nil {
				return "", nil
			}

			strValue, ok := value.(string)
			if !ok {
				return "", er.GenerateErrorCodeAndMessage(400, "value is not a string: "+requestMapFieldValue)
			}
			return strValue, nil
		}

		next, ok := current[key].(map[string]interface{})
		if !ok {
			return "", er.GenerateErrorCodeAndMessage(400, "invalid inputJSON")
		}
		current = next
	}

	return "", nil
}

// setValueInResponse sets a value in the result map
func setValueInResponse(result map[string]interface{}, requestMapFieldKey string, value interface{}) {
	result[requestMapFieldKey] = value
}
