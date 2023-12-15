package middleware

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/schema-creator/services/sql-service/cmd/app/config"
)

func (m *Middleware) NewApm() *newrelic.Application {
	apm, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.Config.NewRelic.NewRelicAppName),
		newrelic.ConfigLicense(config.Config.NewRelic.NewRelicLicense),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		m.l.Errorf("Error newrelic application: %v", err)
	}
	return apm
}
