package pac

import (
	"github.com/getgauge-contrib/gauge-go/gauge"
	"github.com/openshift-pipelines/release-tests/pkg/pac"
	"github.com/openshift-pipelines/release-tests/pkg/store"
)

var _ = gauge.Step("Set <enable> section under <pipelinesAsCode> to <enable: true|false>", func(enable, pipelinesAsCode string, isEnable string) {
	isEnableBool := isEnable == "false"
	pac.VerifyPipelinesAsCodeEnable(store.Clients(), "openshift-pipelines", isEnableBool)
})

var _ = gauge.Step("Verify the installersets related to PAC are <expectedStatus>", func(expectedStatus string) {
	pac.VerifyInstallerSets(store.Clients(), "openshift-pipelines", expectedStatus)
})

var _ = gauge.Step("Verify that the pods related to PAC are <present|not present> from <namespace> namespace", func(expectedStatus, namespace string) {
	pac.VerifyPACPodsStatus(store.Clients(), "openshift-pipelines", expectedStatus)
})

// var _ = gauge.Step("Verify that the custom resource pipelines-as-code of type <pac> is removed <present|not present>", func(expectedStatus string) {
// 	pac.VerifyPACCustomResource(store.Clients(), "openshift-pipelines", expectedStatus)
// })
