package service

import (
	"request-mapper/api/repository"
	er "request-mapper/error"
)

type RequestMapperService interface {
	MapRequest(inputJSON map[string]any, requestMap map[string]any) error
}

type requestMapperService struct {
	repository repository.RequestMapperRepository
}

func NewRequestMapperService(repository repository.RequestMapperRepository) RequestMapperService {
	return &requestMapperService{
		repository: repository,
	}
}

func (s *requestMapperService) MapRequest(inputJSON map[string]any, requestMap map[string]any) error {
	// validate input
	if inputJSON == nil || requestMap == nil {
		return er.GenerateErrorCodeAndMessage(400, "input JSON or requestMap cannot be nil")
	}

	// validate customer object
	// if customer, ok := inputJSON["customer"].(map[string]any); !ok || len(customer) == 0 {
	// 	return nil, er.GenerateErrorCodeAndMessage(400, "customer object cannot be empty")
	// }

	err := s.repository.MapRequest(inputJSON, requestMap)
	if err != nil {
		return err
	}

	return nil
}
