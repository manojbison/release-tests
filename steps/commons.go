package steps

import (
	"github.com/getgauge-contrib/gauge-go/gauge"
	"github.com/openshift-pipelines/release-tests/pkg/clients"
	"github.com/openshift-pipelines/release-tests/pkg/helper"
)

func GetNameSpace() string {
	return gauge.GetScenarioStore()["namespace"].(string)

}

func GetClient() *clients.Clients {
	switch c := gauge.GetScenarioStore()["client"].(type) {
	case *clients.Clients:
		return c
	default:
		return nil
	}
}

func GetOperatorClient() *clients.Clients {
	switch c := gauge.GetSuiteStore()["opclient"].(type) {
	case *clients.Clients:
		return c
	default:
		return nil
	}
}

func GetTknBinaryPath() helper.TknRunner {
	switch n := gauge.GetSuiteStore()["tknPath"].(type) {
	case helper.TknRunner:
		return n
	default:
		panic("Error")
	}
}
