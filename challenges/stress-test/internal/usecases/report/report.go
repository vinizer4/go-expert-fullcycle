package report

import (
	"sort"
	"time"

	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/pkg/httpclient"
	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/usecases/report/dto"
)

type ReportUseCaseInterface interface {
	Execute(input dto.ReportInput) *dto.ReportOutput
}

type ReportUseCase struct{}

func NewReportUseCase() *ReportUseCase {
	return &ReportUseCase{}
}

func (uc *ReportUseCase) Execute(input dto.ReportInput) *dto.ReportOutput {
	statusCount := countStatusOccurrences(input.Results)
	successfulReqs := getSuccessfulRequestsCount(input.Results)

	sortedResults := make([]*httpclient.HttpClientResponse, len(input.Results))
	copy(sortedResults, input.Results)

	sort.Slice(sortedResults, func(i, j int) bool {
		return sortedResults[i].Duration < sortedResults[j].Duration
	})

	return &dto.ReportOutput{
		Duration:           input.Duration,
		StatusCount:        statusCount,
		SuccessfulReqs:     successfulReqs,
		FailedReqs:         uint64(len(input.Results)) - successfulReqs,
		LatencyPercentiles: calculatePercentiles(sortedResults),
	}
}

func countStatusOccurrences(responses []*httpclient.HttpClientResponse) map[int]uint64 {
	statusMap := make(map[int]uint64)

	for _, r := range responses {
		if r.StatusCode == nil {
			continue
		}

		statusMap[*r.StatusCode]++
	}

	return statusMap
}

func getSuccessfulRequestsCount(responses []*httpclient.HttpClientResponse) uint64 {
	var c uint64

	for _, r := range responses {
		if r.Error == nil {
			c++
		}
	}

	return c
}

func calculatePercentiles(sortedResps []*httpclient.HttpClientResponse) map[int]time.Duration {
	sortedRespsLen := len(sortedResps)
	percentiles := make(map[int]time.Duration)

	for _, p := range []int{50, 75, 90, 95, 99} {
		idx := (sortedRespsLen * p) / 100

		if sortedRespsLen == 0 {
			percentiles[p] = 0
			continue
		}

		if sortedRespsLen > 1 {
			idx = idx - 1
		}

		percentiles[p] = sortedResps[idx].Duration
	}

	return percentiles
}
