package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	hits *prometheus.CounterVec
)

func New() {
	hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits",
	}, []string{"node_name"})

	prometheus.MustRegister(hits)
}

func CreateRequestHits(status int, r *http.Request) {
	hits.WithLabelValues("node_1").Inc()
}
