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

package controller

import (
	"context"
	"runtime"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
)

const (
	scope = "sigs.k8s.io/secrets-store-sync-controller"
)

var (
	providerKey = "provider"
	osTypeKey   = "os_type"
	errorKey    = "error_key"
	runtimeOS   = runtime.GOOS
)

type reporter struct {
	secretSyncTotal    metric.Int64Counter
	secretSyncErrTotal metric.Int64Counter
	secretSyncDuration metric.Float64Histogram
}

type StatsReporter interface {
	ReportSyncK8SecretCtMetric(ctx context.Context, provider string)
	ReportSyncK8SecretErrCtMetric(ctx context.Context, provider string, errType string)
	ReportSyncK8SecretDuration(ctx context.Context, duration float64)
}

func NewStatsReporter() (StatsReporter, error) {
	var err error

	r := &reporter{}
	meter := global.Meter(scope)

	if r.secretSyncTotal, err = meter.Int64Counter("secret_sync_reconciliation", metric.WithDescription("Total number of secret sync call")); err != nil {
		return nil, err
	}
	if r.secretSyncErrTotal, err = meter.Int64Counter("secret_sync_reconciliation_error", metric.WithDescription("Total number of secret sync with failure")); err != nil {
		return nil, err
	}
	if r.secretSyncDuration, err = meter.Float64Histogram("secret_sync_duration", metric.WithDescription("Distribution of how long it took to sync k8s secret")); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *reporter) ReportSyncK8SecretCtMetric(ctx context.Context, provider string) {
	opt := metric.WithAttributes(
		attribute.Key(providerKey).String(provider),
		attribute.Key(osTypeKey).String(runtimeOS),
	)
	r.secretSyncTotal.Add(ctx, 1, opt)
}

func (r *reporter) ReportSyncK8SecretErrCtMetric(ctx context.Context, provider string, errType string) {
	opt := metric.WithAttributes(
		attribute.Key(providerKey).String(provider),
		attribute.Key(osTypeKey).String(runtimeOS),
		attribute.Key(errorKey).String(errType),
	)
	r.secretSyncErrTotal.Add(ctx, 1, opt)
}

func (r *reporter) ReportSyncK8SecretDuration(ctx context.Context, duration float64) {
	opt := metric.WithAttributes(
		attribute.Key(osTypeKey).String(runtimeOS),
	)
	r.secretSyncDuration.Record(ctx, duration, opt)
}
