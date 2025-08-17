package stress

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/pkg/httpclient"
	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/tests/mocks"
	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/tests/utils"
	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/usecases/stress/dto"
)

type StressTestUseCaseTestSuite struct {
	suite.Suite
	HTTPClientMock    *mocks.HttpClientMock
	StressTestUseCase StressTestUseCaseInterface
}

func (s *StressTestUseCaseTestSuite) SetupTest() {
	s.HTTPClientMock = new(mocks.HttpClientMock)
	s.StressTestUseCase = NewStressTestUseCase(s.HTTPClientMock)
}

func (s *StressTestUseCaseTestSuite) cleanMocks() {
	s.HTTPClientMock.ExpectedCalls = nil
}

func TestStressTestUseCase(t *testing.T) {
	suite.Run(t, new(StressTestUseCaseTestSuite))
}

func (s *StressTestUseCaseTestSuite) TestExecute() {
	s.Run("Should execute stress test with received params", func() {
		defer s.cleanMocks()

		input := dto.StressTestInput{
			Concurrency: 1,
			Requests:    10,
			URL:         "http://localhost:8080",
		}

		s.HTTPClientMock.On("Get", input.URL).Return(&httpclient.HttpClientResponse{
			StatusCode: utils.IntPtr(200),
			Duration:   time.Duration(rand.Intn(500)) * time.Millisecond,
			Error:      nil,
		})

		r, err := s.StressTestUseCase.Execute(input)

		s.NoError(err)
		s.NotNil(r)
		s.Len(r.Results, 10)
		s.HTTPClientMock.AssertNumberOfCalls(s.T(), "Get", 10)
	})

	s.Run("Should return error when input is invalid", func() {
		defer s.cleanMocks()

		input := dto.StressTestInput{
			Concurrency: 0,
			Requests:    0,
			URL:         "",
		}

		_, err := s.StressTestUseCase.Execute(input)

		s.Error(err)
		s.ErrorContains(err, "Field validation for 'URL' failed on the 'required' tag")
	})
}
