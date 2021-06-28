//go:generate mockery --name Service

package ride

const (
	buffSize = 50
)

// Service provides required methods to manipulate ride.
type Service interface {
	//
}

// serviceImpl implements Service interface.
type serviceImpl struct{}

// NewService returns a new instance of serviceImpl.
func NewService() Service {
	return &serviceImpl{}
}
