package stress

import (
	"fmt"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/color"

	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/pkg/httpclient"
	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/usecases/stress/dto"
)

type StressTestUseCaseInterface interface {
	Execute(input dto.StressTestInput) (*dto.StressTestOutput, error)
}

type StressTestUseCase struct {
	HTTPClient httpclient.HttpClientInterface
	validator  *validator.Validate
	spinner    *spinner.Spinner
}

type Job struct {
	Endpoint string
}

func NewStressTestUseCase(c httpclient.HttpClientInterface) *StressTestUseCase {
	green := color.FgGreen.Render

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Color("blue")
	s.Suffix = " Executing stress test..."
	s.FinalMSG = fmt.Sprintf("%s Stress test finished!\n\n", green("✓"))

	return &StressTestUseCase{
		HTTPClient: c,
		validator:  validator.New(validator.WithRequiredStructEnabled()),
		spinner:    s,
	}
}

func (uc *StressTestUseCase) Execute(input dto.StressTestInput) (*dto.StressTestOutput, error) {
	if err := uc.validateInput(input); err != nil {
		return nil, err
	}

	uc.spinner.Start()

	reqsCount := int(input.Requests)

	jobs := make(chan Job, reqsCount)
	results := make(chan *httpclient.HttpClientResponse, reqsCount)
	var wg sync.WaitGroup

	for i := 0; i < int(input.Concurrency); i++ {
		go uc.runJobs(jobs, results, &wg)
	}

	wg.Add(reqsCount)
	start := time.Now()

	for range reqsCount {
		jobs <- Job{
			Endpoint: input.URL,
		}
	}

	close(jobs)
	wg.Wait()

	uc.spinner.Stop()

	output := &dto.StressTestOutput{
		Duration: time.Since(start),
	}

	for i := 0; i < reqsCount; i++ {
		res := <-results
		output.Results = append(output.Results, res)
	}

	return output, nil
}

func (uc *StressTestUseCase) validateInput(input dto.StressTestInput) error {
	return uc.validator.Struct(input)
}

func (uc *StressTestUseCase) runJobs(
	jobs <-chan Job,
	results chan<- *httpclient.HttpClientResponse,
	wg *sync.WaitGroup,
) {
	for job := range jobs {
		res := uc.HTTPClient.Get(job.Endpoint)
		results <- res

		wg.Done()
	}
}
