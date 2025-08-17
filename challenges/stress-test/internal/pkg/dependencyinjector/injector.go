package dependencyinjector

import (
	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/infra/cli"
	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/infra/cli/commands"
	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/pkg/httpclient"
	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/usecases/report"
	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/usecases/stress"
)

type DependencyInjectorInterface interface {
	Inject() (*Dependencies, error)
}

type DependencyInjector struct{}

type Dependencies struct {
	CLI cli.CLIInterface
}

func NewDependencyInjector() *DependencyInjector {
	return &DependencyInjector{}
}

func (d *DependencyInjector) Inject() (*Dependencies, error) {
	httpClient := httpclient.NewHttpClient()
	stressTestUseCase := stress.NewStressTestUseCase(httpClient)
	reportUseCase := report.NewReportUseCase()
	stressTestCmd := commands.NewStressTestCmd(stressTestUseCase, reportUseCase)

	cli := cli.NewCLI(stressTestCmd.Build())

	return &Dependencies{
		CLI: cli,
	}, nil
}
