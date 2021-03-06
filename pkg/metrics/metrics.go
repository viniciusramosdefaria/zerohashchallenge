package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	Code = "code"
	Name = "name"
)

var (
	Metric = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "team_points",
		Help: "How many games each team won",
	},
		[]string{
			Code,
			Name,
		})
)

func Add(code, name string)  {
	Metric.
		With(prometheus.Labels{Code: code, Name: name}).
		Inc()
}
