# Logstore Helm Test

This folder contains a collection of go tests that test if a Logstore canary is running correctly. It's primary use it to test that the helm chart is working correctly by using metrics from the Logstore canary. In the helm chart, the template for this test is only available if you are running both the Logstore canary and have self monitoring enabled (as the Logstore canary's logs need to be in Logstore for it to work). However, the tests in this folder can be run against any running Logstore canary using `go test`.

## Instructions

Run `go test .` from this directory, or use the Docker image published at `acme/logstore-helm-test`.
