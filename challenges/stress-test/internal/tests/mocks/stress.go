package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/usecases/stress/dto"
)

type StressTestUseCaseMock struct {
	mock.Mock
}

func (m *StressTestUseCaseMock) Execute(input dto.StressTestInput) (*dto.StressTestOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*dto.StressTestOutput), args.Error(1)
}
