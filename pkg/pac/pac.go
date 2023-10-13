package pac

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/getgauge-contrib/gauge-go/gauge"
	"github.com/openshift-pipelines/release-tests/pkg/clients"
)

func VerifyPipelinesAsCodeEnable(cs *clients.Clients, namespace string, enable bool) (string, error) {
	// Construct the JSON payload based on the 'enable' parameter
	enableStatus := "true"
	if enable {
		enableStatus = "false"
	}
	payload := fmt.Sprintf(`{"spec":{"platforms":{"openshift":{"pipelinesAsCode":{"enable": %s}}}}}`, enableStatus)

	cmd := exec.Command("oc", "patch", "tektonconfigs.operator.tekton.dev", "config", "-n", "openshift-pipelines", "--type", "merge", "-p", payload)

	// Run the 'oc' command
	if err := cmd.Run(); err != nil {
		return "", err
	}

	// Return a message indicating the status change
	return fmt.Sprintf("PipelinesAsCode enable status has been set to %s", enableStatus), nil
}

func VerifyInstallerSets(cs *clients.Clients, namespace string, expectedStatus string) {
	cmd := exec.Command("oc", "get", "tektoninstallersets", "-n", namespace, "-o", "custom-columns=NAME:.metadata.name")
	cmdOutput, err := cmd.CombinedOutput()

	if err != nil {
		gauge.WriteMessage(fmt.Sprintf("Failed to get InstallerSets: %v", err))
		return
	}

	installerSets := strings.Split(string(cmdOutput), "\n")
	for _, line := range installerSets {
		// Skip the header line
		if strings.Contains(line, "NAME") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 1 {
			continue
		}

		name := parts[0]

		if strings.HasPrefix(name, "openshiftpipelinesascode-") {
			if expectedStatus == "present" {
				gauge.WriteMessage(fmt.Sprintf("InstallerSet '%s' is present", name))
			} else if expectedStatus == "not present" {
				gauge.WriteMessage(fmt.Sprintf("InstallerSet '%s' is not present", name))
			}
		}
	}
}

// func VerifyPACPods(namespace string, expectedStatus string) {
// 	cmd := exec.Command("oc", "get", "pods", "-n", namespace)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	if err := cmd.Run(); err != nil {
// 		if expectedStatus == "present" {
// 			gauge.WriteMessage("PAC-related pods are not present")
// 		}
// 		return
// 	}

// 	if expectedStatus == "not present" {
// 		gauge.WriteMessage("PAC-related pods are present")
// 	}
// }

// func VerifyPACCustomResource(namespace string, expectedStatus string) {
// 	cmd := exec.Command("oc", "get", "pipelines-as-code", "-n", namespace)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	if err := cmd.Run(); err != nil {
// 		if expectedStatus == "present" {
// 			gauge.WriteMessage("'pipelines-as-code' custom resource is not present")
// 		}
// 		return
// 	}

// 	if expectedStatus == "not present" {
// 		gauge.WriteMessage("'pipelines-as-code' custom resource is present")
// 	}
// }
