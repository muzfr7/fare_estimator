//go:generate mockery --name Service

package fare

import (
	pathUsecase "github.com/muzfr7/fare_estimator/app/usecases/path"
)

const (
	buffSize = 50
)

// Service provides required methods to manipulate ride fare.
type Service interface {
	//
}

// serviceImpl implements Service interface.
type serviceImpl struct {
	pathSVC pathUsecase.Service
}

// NewService returns a new instance of serviceImpl.
func NewService(pathSVC pathUsecase.Service) Service {
	return &serviceImpl{
		pathSVC: pathSVC,
	}
}
