package metrics

import (
	"strconv"

	"github.com/pingcap/errors"
	"github.com/pingcap/parser/terror"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// PanicCounter measures the count of panics.
	PanicCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: ModuleWeirProxy,
			Subsystem: LabelServer,
			Name:      "panic_total",
			Help:      "Counter of panic.",
		}, []string{LblType})

	QueryTotalCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: ModuleWeirProxy,
			Subsystem: LabelServer,
			Name:      "query_total",
			Help:      "Counter of queries.",
		}, []string{LblType, LblResult})

	ExecuteErrorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: ModuleWeirProxy,
			Subsystem: LabelServer,
			Name:      "execute_error_total",
			Help:      "Counter of execute errors.",
		}, []string{LblType})

	CriticalErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: ModuleWeirProxy,
			Subsystem: LabelServer,
			Name:      "critical_error_total",
			Help:      "Counter of critical errors.",
		})

	ConnGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: ModuleWeirProxy,
			Subsystem: LabelServer,
			Name:      "connections",
			Help:      "Number of connections.",
		})

	EventStart        = "start"
	EventGracefulDown = "graceful_shutdown"
	// Eventkill occurs when the server.Kill() function is called.
	EventKill          = "kill"
	EventClose         = "close"
	ServerEventCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: ModuleWeirProxy,
			Subsystem: LabelServer,
			Name:      "event_total",
			Help:      "Counter of weirproxy-server event.",
		}, []string{LblType})

	GetTokenDurationHistogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: ModuleWeirProxy,
			Subsystem: LabelServer,
			Name:      "get_token_duration_seconds",
			Help:      "Duration (us) for getting token, it should be small until concurrency limit is reached.",
			Buckets:   prometheus.ExponentialBuckets(1, 2, 30), // 1us ~ 528s
		})

	HandShakeErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: ModuleWeirProxy,
			Subsystem: LabelServer,
			Name:      "handshake_error_total",
			Help:      "Counter of hand shake error.",
		},
	)
)

// ExecuteErrorToLabel converts an execute error to label.
func ExecuteErrorToLabel(err error) string {
	err = errors.Cause(err)
	switch x := err.(type) {
	case *terror.Error:
		return x.Class().String() + ":" + strconv.Itoa(int(x.Code()))
	default:
		return "unknown"
	}
}
