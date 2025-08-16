/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mocks // import "sigs.k8s.io/secrets-store-sync-controller/internal/controller/mocks"

import "context"

type FakeReporter struct {
	reportSyncK8SecretCtMetricInvoked    int
	reportSyncK8SecretErrCtMetricInvoked int
	reportSyncK8SecretDurationInvoked    int
}

func NewFakeReporter() *FakeReporter {
	return &FakeReporter{}
}

func (f *FakeReporter) ReportSyncK8SecretCtMetric(_ context.Context, _ string) {
	f.reportSyncK8SecretCtMetricInvoked++
}

func (f *FakeReporter) ReportSyncK8SecretErrCtMetric(_ context.Context, _ string, _ string) {
	f.reportSyncK8SecretErrCtMetricInvoked++
}

func (f *FakeReporter) ReportSyncK8SecretDuration(_ context.Context, _ float64) {
	f.reportSyncK8SecretDurationInvoked++
}

func (f *FakeReporter) ReportSyncK8SecretCtMetricInvoked() int {
	return f.reportSyncK8SecretCtMetricInvoked
}

func (f *FakeReporter) ReportSyncK8SecretErrCtMetricInvoked() int {
	return f.reportSyncK8SecretErrCtMetricInvoked
}

func (f *FakeReporter) ReportSyncK8SecretDurationInvoked() int {
	return f.reportSyncK8SecretDurationInvoked
}
