//go:generate mockery --name Service

package path

const (
	earthRadius = float64(6371)
)

// Service provides required methods to manipulate ride path.
type Service interface {
	//
}

// serviceImpl implements Service interface.
type serviceImpl struct{}

// NewService returns a new instance of serviceImpl.
func NewService() Service {
	return &serviceImpl{}
}
