// Copyright Splunk, Inc.
// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package components

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/countconnector"
	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector"
	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/pulsarexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/httpforwarder"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/dockerobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecstaskobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbyattrsprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/logstransformprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/routingprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanmetricsprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudfoundryreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/collectdreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/journaldreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8seventsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sobjectsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkametricsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mysqlreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/postgresqlreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sapmreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/signalfxreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/simpleprometheusreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkhecreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sshcheckreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/syslogreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tcplogreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/udplogreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/wavefrontreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowseventlogreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsperfcountersreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/connector/forwardconnector"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/debugexporter"
	"go.opentelemetry.io/collector/exporter/loggingexporter"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/exporter/otlphttpexporter"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/ballastextension"
	"go.opentelemetry.io/collector/extension/zpagesextension"
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/batchprocessor"
	"go.opentelemetry.io/collector/processor/memorylimiterprocessor"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"
	"go.uber.org/multierr"

	"github.com/signalfx/splunk-otel-collector/internal/exporter/httpsinkexporter"
	"github.com/signalfx/splunk-otel-collector/internal/receiver/databricksreceiver"
	"github.com/signalfx/splunk-otel-collector/internal/receiver/discoveryreceiver"
	"github.com/signalfx/splunk-otel-collector/internal/receiver/lightprometheusreceiver"
	"github.com/signalfx/splunk-otel-collector/internal/receiver/scriptedinputsreceiver"
	"github.com/signalfx/splunk-otel-collector/internal/receiver/signalfxgatewayprometheusremotewritereceiver"
	"github.com/signalfx/splunk-otel-collector/pkg/extension/smartagentextension"
	"github.com/signalfx/splunk-otel-collector/pkg/processor/timestampprocessor"
	"github.com/signalfx/splunk-otel-collector/pkg/receiver/smartagentreceiver"
)

func Get() (otelcol.Factories, error) {
	var errs []error
	extensions, err := extension.MakeFactoryMap(
		ballastextension.NewFactory(),
		basicauthextension.NewFactory(),
		ecsobserver.NewFactory(),
		ecstaskobserver.NewFactory(),
		dockerobserver.NewFactory(),
		filestorage.NewFactory(),
		healthcheckextension.NewFactory(),
		hostobserver.NewFactory(),
		httpforwarder.NewFactory(),
		k8sobserver.NewFactory(),
		pprofextension.NewFactory(),
		smartagentextension.NewFactory(),
		zpagesextension.NewFactory(),
	)
	if err != nil {
		errs = append(errs, err)
	}

	receivers, err := receiver.MakeFactoryMap(
		azureeventhubreceiver.NewFactory(),
		carbonreceiver.NewFactory(),
		cloudfoundryreceiver.NewFactory(),
		collectdreceiver.NewFactory(),
		databricksreceiver.NewFactory(),
		lightprometheusreceiver.NewFactory(),
		discoveryreceiver.NewFactory(),
		fluentforwardreceiver.NewFactory(),
		filelogreceiver.NewFactory(),
		hostmetricsreceiver.NewFactory(),
		jaegerreceiver.NewFactory(),
		jmxreceiver.NewFactory(),
		journaldreceiver.NewFactory(),
		k8sclusterreceiver.NewFactory(),
		k8seventsreceiver.NewFactory(),
		k8sobjectsreceiver.NewFactory(),
		kafkametricsreceiver.NewFactory(),
		kafkareceiver.NewFactory(),
		kubeletstatsreceiver.NewFactory(),
		mongodbatlasreceiver.NewFactory(),
		mongodbreceiver.NewFactory(),
		mysqlreceiver.NewFactory(),
		oracledbreceiver.NewFactory(),
		otlpreceiver.NewFactory(),
		postgresqlreceiver.NewFactory(),
		prometheusreceiver.NewFactory(),
		receivercreator.NewFactory(),
		redisreceiver.NewFactory(),
		sapmreceiver.NewFactory(),
		scriptedinputsreceiver.NewFactory(),
		signalfxreceiver.NewFactory(),
		signalfxgatewayprometheusremotewritereceiver.NewFactory(),
		simpleprometheusreceiver.NewFactory(),
		smartagentreceiver.NewFactory(),
		solacereceiver.NewFactory(),
		splunkhecreceiver.NewFactory(),
		sqlqueryreceiver.NewFactory(),
		sshcheckreceiver.NewFactory(),
		statsdreceiver.NewFactory(),
		syslogreceiver.NewFactory(),
		tcplogreceiver.NewFactory(),
		udplogreceiver.NewFactory(),
		vcenterreceiver.NewFactory(),
		wavefrontreceiver.NewFactory(),
		windowsperfcountersreceiver.NewFactory(),
		windowseventlogreceiver.NewFactory(),
		zipkinreceiver.NewFactory(),
	)
	if err != nil {
		errs = append(errs, err)
	}

	exporters, err := exporter.MakeFactoryMap(
		awss3exporter.NewFactory(),
		debugexporter.NewFactory(),
		fileexporter.NewFactory(),
		kafkaexporter.NewFactory(),
		loadbalancingexporter.NewFactory(),
		loggingexporter.NewFactory(),
		otlpexporter.NewFactory(),
		otlphttpexporter.NewFactory(),
		sapmexporter.NewFactory(),
		signalfxexporter.NewFactory(),
		splunkhecexporter.NewFactory(),
		httpsinkexporter.NewFactory(),
		pulsarexporter.NewFactory(),
	)
	if err != nil {
		errs = append(errs, err)
	}

	processors, err := processor.MakeFactoryMap(
		attributesprocessor.NewFactory(),
		batchprocessor.NewFactory(),
		cumulativetodeltaprocessor.NewFactory(),
		filterprocessor.NewFactory(),
		groupbyattrsprocessor.NewFactory(),
		k8sattributesprocessor.NewFactory(),
		logstransformprocessor.NewFactory(),
		memorylimiterprocessor.NewFactory(),
		metricstransformprocessor.NewFactory(),
		probabilisticsamplerprocessor.NewFactory(),
		resourcedetectionprocessor.NewFactory(),
		resourceprocessor.NewFactory(),
		routingprocessor.NewFactory(),
		spanmetricsprocessor.NewFactory(),
		spanprocessor.NewFactory(),
		tailsamplingprocessor.NewFactory(),
		timestampprocessor.NewFactory(),
		transformprocessor.NewFactory(),
	)
	if err != nil {
		errs = append(errs, err)
	}

	connectors, err := connector.MakeFactoryMap(
		countconnector.NewFactory(),
		forwardconnector.NewFactory(),
		routingconnector.NewFactory(),
		spanmetricsconnector.NewFactory(),
	)
	if err != nil {
		errs = append(errs, err)
	}

	factories := otelcol.Factories{
		Extensions: extensions,
		Receivers:  receivers,
		Processors: processors,
		Exporters:  exporters,
		Connectors: connectors,
	}

	return factories, multierr.Combine(errs...)
}
