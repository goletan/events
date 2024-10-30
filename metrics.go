package events

import (
	"github.com/goletan/observability/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type EventsMetrics struct{}

// Event Metrics: Track event processing metrics
var (
	EventPublished = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "goletan",
			Subsystem: "events",
			Name:      "published_total",
			Help:      "Total number of events published.",
		},
		[]string{"event_type"},
	)
	EventDropped = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "goletan",
			Subsystem: "events",
			Name:      "dropped_total",
			Help:      "Total number of events dropped.",
		},
		[]string{"event_type"},
	)
	EventProcessed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "goletan",
			Subsystem: "events",
			Name:      "processed_total",
			Help:      "Total number of events processed.",
		},
		[]string{"event_type"},
	)
	EventSent = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "goletan",
			Subsystem: "events",
			Name:      "sent_total",
			Help:      "Total number of events sent.",
		},
		[]string{"event_type"},
	)
	EventFailed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "goletan",
			Subsystem: "events",
			Name:      "failed_total",
			Help:      "Total number of events failed.",
		},
		[]string{"event_type"},
	)
	EventDLQ = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "goletan",
			Subsystem: "events",
			Name:      "dlq_total",
			Help:      "Total number of events dead letter queue.",
		},
		[]string{"event_type"},
	)
	SubscriberAdded = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "goletan",
			Subsystem: "events",
			Name:      "subscriber_added_total",
			Help:      "Total number of subscriber added.",
		},
		[]string{"event_type"},
	)
	SubscriberRemoved = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "goletan",
			Subsystem: "events",
			Name:      "subscriber_removed_total",
			Help:      "Total number of subscriber removed.",
		},
		[]string{"event_type"},
	)
)

func InitMetrics() {
	metrics.NewManager().Register(&EventsMetrics{})
}

func (em *EventsMetrics) Register() error {

	if err := prometheus.Register(EventPublished); err != nil {
		return err
	}

	if err := prometheus.Register(EventDropped); err != nil {
		return err
	}

	if err := prometheus.Register(EventProcessed); err != nil {
		return err
	}

	if err := prometheus.Register(EventSent); err != nil {
		return err
	}

	if err := prometheus.Register(EventFailed); err != nil {
		return err
	}

	if err := prometheus.Register(EventDLQ); err != nil {
		return err
	}

	if err := prometheus.Register(SubscriberAdded); err != nil {
		return err
	}

	if err := prometheus.Register(SubscriberRemoved); err != nil {
		return err
	}

	return nil
}

// IncrementEventPublished increments the event published counter.
func IncrementEventPublished(eventType string) {
	EventPublished.WithLabelValues(eventType).Inc()
}

// IncrementEventDropped increments the event dropped counter.
func IncrementEventDropped(eventType string) {
	EventDropped.WithLabelValues(eventType).Inc()
}

// IncrementEventProcessed increments the event processed counter.
func IncrementEventProcessed(eventType string) {
	EventProcessed.WithLabelValues(eventType).Inc()
}

// IncrementEventSent increments the event sent counter.
func IncrementEventSent(eventType string) {
	EventSent.WithLabelValues(eventType).Inc()
}

// IncrementEventFailed increments the event failed counter.
func IncrementEventFailed(eventType string) {
	EventFailed.WithLabelValues(eventType).Inc()
}

// IncrementEventDQL increments the event Dead Letter Queue counter.
func IncrementEventDQL(eventType string) {
	EventDLQ.WithLabelValues(eventType).Inc()
}

// SubscriberAdded increments the number of subscriber added counter.
func IncrementSubscriberAdded(eventType string) {
	SubscriberAdded.WithLabelValues(eventType).Inc()
}

// SubscriberRemoved increments the number of subscriber removed counter.
func IncrementSubscriberRemoved(eventType string) {
	SubscriberRemoved.WithLabelValues(eventType).Inc()
}