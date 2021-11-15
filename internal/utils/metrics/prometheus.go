package metrics

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	hits *prometheus.CounterVec
)

func New() {
	hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits",
	}, []string{"node1"})

	prometheus.MustRegister(hits)
}

func CreateRequestHits(status int, r *http.Request) {
	hits.WithLabelValues(strconv.Itoa(status), strings.Split(r.URL.Path, "/")[2], r.Method).Inc()
}
