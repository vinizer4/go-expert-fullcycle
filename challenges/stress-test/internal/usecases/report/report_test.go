package report

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/pkg/httpclient"
	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/tests/utils"
	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/usecases/report/dto"
)

type ReportUseCaseTestSuite struct {
	suite.Suite
	ReportUseCase ReportUseCaseInterface
}

func (s *ReportUseCaseTestSuite) SetupTest() {
	s.ReportUseCase = NewReportUseCase()
}

func TestReportUseCase(t *testing.T) {
	suite.Run(t, new(ReportUseCaseTestSuite))
}

func (s *ReportUseCaseTestSuite) TestExecute() {
	s.Run("Should calculate report with received params", func() {
		input := dto.ReportInput{
			Duration: time.Duration(1000) * time.Millisecond,
			Results: []*httpclient.HttpClientResponse{
				{
					StatusCode: utils.IntPtr(200),
					Duration:   time.Duration(100) * time.Millisecond,
					Error:      nil,
				},
				{
					StatusCode: utils.IntPtr(404),
					Duration:   time.Duration(200) * time.Millisecond,
					Error:      nil,
				},
				{
					StatusCode: utils.IntPtr(500),
					Duration:   time.Duration(300) * time.Millisecond,
					Error:      nil,
				},
				{
					StatusCode: utils.IntPtr(200),
					Duration:   time.Duration(400) * time.Millisecond,
					Error:      nil,
				},
				{
					StatusCode: nil,
					Duration:   time.Duration(500) * time.Millisecond,
					Error:      errors.New("any-error"),
				},
			},
		}

		r := s.ReportUseCase.Execute(input)

		s.NotNil(r)
		s.Equal(time.Duration(1000)*time.Millisecond, r.Duration)
		s.Equal(uint64(2), r.StatusCount[200])
		s.Equal(uint64(1), r.StatusCount[404])
		s.Equal(uint64(1), r.StatusCount[500])
		s.Equal(uint64(4), r.SuccessfulReqs)
		s.Equal(uint64(1), r.FailedReqs)
		s.Len(r.LatencyPercentiles, 5)
	})

	s.Run("Should calculate report with empty responses list", func() {
		input := dto.ReportInput{
			Duration: time.Duration(1000) * time.Millisecond,
			Results:  []*httpclient.HttpClientResponse{},
		}

		r := s.ReportUseCase.Execute(input)

		s.NotNil(r)
		s.Equal(time.Duration(1000)*time.Millisecond, r.Duration)
		s.Equal(uint64(0), r.SuccessfulReqs)
		s.Equal(uint64(0), r.FailedReqs)
		s.Len(r.StatusCount, 0)
		s.Len(r.LatencyPercentiles, 5)
	})
}
