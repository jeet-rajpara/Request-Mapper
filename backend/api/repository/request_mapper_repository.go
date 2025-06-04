package repository

import (
	"strings"
)

type RequestMapperRepository interface {
	MapRequest(inputJSON map[string]any, requestMap map[string]any) error
}

type requestMapperRepository struct{}

func NewRequestMapperRepository() RequestMapperRepository {
	return &requestMapperRepository{}
}

// getValueFromInputJSON retrieves a value from the input JSON using a dot-notation path
func getValueFromInputJSON(inputJSON map[string]any, requestMapFieldValueString string) (any, error) {

	keys := strings.Split(requestMapFieldValueString, ".")

	current := inputJSON

	for i, key := range keys {
		if i == len(keys)-1 {
			value, exists := current[key]
			if !exists {
				return requestMapFieldValueString, nil
			}

			if value == nil {
				return "", nil
			}

			return value, nil
		}

		next, ok := current[key].(map[string]any)
		if !ok {
			return requestMapFieldValueString, nil
		}
		current = next
	}

	return "", nil
}

func (r *requestMapperRepository) MapRequest(inputJSON map[string]any, requestMap map[string]any) error {

	for requestMapFieldKey, requestMapFieldValue := range requestMap {

		requestMapFieldValueString, ok := requestMapFieldValue.(string)
		if !ok {
			isMap, ok := requestMapFieldValue.(map[string]any)
			if ok {
				err := r.MapRequest(inputJSON, isMap)
				if err != nil {
					return err
				}
			}
			continue
		}

		value, err := getValueFromInputJSON(inputJSON, requestMapFieldValueString)
		if err != nil {
			return err
		}

		requestMap[requestMapFieldKey] = value
	}

	return nil
}
