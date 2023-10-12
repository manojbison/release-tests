PIPELINES-20
# Pipelines As Code tests

## Enable/Disable PAC: PIPELINES-20-TC01
Tags: pac, sanity, to-do
Component: PAC
Level: Integration
Type: Functional
Importance: Critical

This scenario tests enable/disable of pipelines as code from tektonconfig custom resource

Steps:
  * Set "enable" section under "pipelinesAsCode" to "false"
  * Set "enable" section under "pipelinesAsCode" to "true"