package dto

import (
	"time"

	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/pkg/httpclient"
)

type ReportInput struct {
	Duration time.Duration                    `validate:"required"`
	Results  []*httpclient.HttpClientResponse `validate:"required"`
}

type ReportOutput struct {
	Duration           time.Duration
	StatusCount        map[int]uint64
	SuccessfulReqs     uint64
	FailedReqs         uint64
	LatencyPercentiles map[int]time.Duration
}
