package commands

import (
	"net/http"

	"github.com/gookit/color"
	"github.com/spf13/cobra"

	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/usecases/report"
	report_dto "github.com/vinizer4/go-expert-fullcycle/stress_test/internal/usecases/report/dto"
	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/usecases/stress"
	stress_dto "github.com/vinizer4/go-expert-fullcycle/stress_test/internal/usecases/stress/dto"
)

var blue = color.FgBlue.Render
var bold = color.OpBold.Render

type StressTestCmdInterface interface {
	Build() *cobra.Command
}

type StressTestCmd struct {
	StressTestUseCase stress.StressTestUseCaseInterface
	ReportUseCase     report.ReportUseCaseInterface
}

func NewStressTestCmd(
	s stress.StressTestUseCaseInterface,
	r report.ReportUseCaseInterface,
) *StressTestCmd {
	return &StressTestCmd{
		StressTestUseCase: s,
		ReportUseCase:     r,
	}
}

func (s *StressTestCmd) Build() *cobra.Command {
	cmd := &cobra.Command{
		Short: "Stress test a given URL",
		Long:  "Executes a stress test on a given URL with a given number of requests and concurrency.",
		RunE:  s.run,
	}

	cmd.Flags().String("url", "", "service URL to test")
	cmd.Flags().Uint64("requests", 0, "number of requests to perform")
	cmd.Flags().Uint64("concurrency", 0, "number of simultaneous requests to make at a time")
	cmd.MarkFlagsRequiredTogether("url", "requests", "concurrency")

	return cmd
}

func (s *StressTestCmd) run(cmd *cobra.Command, args []string) error {
	url, _ := cmd.Flags().GetString("url")
	requests, _ := cmd.Flags().GetUint64("requests")
	concurrency, _ := cmd.Flags().GetUint64("concurrency")

	input := stress_dto.StressTestInput{
		URL:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}

	s.printHeader(cmd, input)

	stressOut, err := s.StressTestUseCase.Execute(input)
	if err != nil {
		return err
	}

	reportInput := report_dto.ReportInput{
		Duration: stressOut.Duration,
		Results:  stressOut.Results,
	}

	reportOut := s.ReportUseCase.Execute(reportInput)
	s.printReport(cmd, reportOut)

	return nil
}

func (s *StressTestCmd) printHeader(cmd *cobra.Command, input stress_dto.StressTestInput) {
	cmd.Println(bold(blue("===== API STRESS TEST =====")))

	cmd.Printf("🔗 URL: %s\n", input.URL)
	cmd.Printf("🔁 Requests: %d\n", input.Requests)
	cmd.Printf("🚀 Concurrency: %d\n\n", input.Concurrency)
}

func (s *StressTestCmd) printReport(cmd *cobra.Command, output *report_dto.ReportOutput) {
	cmd.Println(bold(blue("===== TEST RESULTS =====")))
	cmd.Printf("⌛ Test duration: %.3fs\n", output.Duration.Seconds())
	cmd.Printf("👍 Successfull requests: %d\n", output.SuccessfulReqs)
	cmd.Printf("😢 Failed requests: %d\n", output.FailedReqs)

	cmd.Printf("\n🧮 Status count:\n")
	for status, count := range output.StatusCount {
		cmd.Printf("\t%s %s: %d\n", bold(status), bold(http.StatusText(status)), count)
	}

	cmd.Printf("\n🕒 Percentiles:\n")
	for percentile, duration := range output.LatencyPercentiles {
		cmd.Printf("\t%s: %.3fs\n", bold("P", percentile), duration.Seconds())
	}
}
