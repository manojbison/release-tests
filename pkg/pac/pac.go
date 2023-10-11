package pac

import (
	"fmt"
	"os/exec"

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

// func VerifyInstallerSets(namespace string, expectedStatus string) {
// 	cmd := exec.Command("oc", "get", "installersets", "-n", namespace)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	if err := cmd.Run(); err != nil {
// 		if expectedStatus == "present" {
// 			gauge.WriteMessage("InstallerSets are not present")
// 		}
// 		return
// 	}

// 	if expectedStatus == "not present" {
// 		gauge.WriteMessage("InstallerSets are present")
// 	}
// }

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
