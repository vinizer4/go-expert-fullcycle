package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/usecases/report/dto"
)

type ReportUseCaseMock struct {
	mock.Mock
}

func (m *ReportUseCaseMock) Execute(input dto.ReportInput) *dto.ReportOutput {
	args := m.Called(input)
	return args.Get(0).(*dto.ReportOutput)
}
