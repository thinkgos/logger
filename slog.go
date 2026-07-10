package logger

import (
	"context"
	"log/slog"
	"time"
)

// Verify at compile time that SlogHandler satisfies the slog.Handler interface.
var _ slog.Handler = (*SlogHandler)(nil)

// SlogHandler implements the slog.Handler interface using a logger.Log
// as the underlying log backend. This allows code that uses the standard
// library's slog package to route log output through logger.
type SlogHandler struct {
	logger  *Log
	prefix  string // group prefix for nested groups
	timeKey string // time field name
	attrs   []slog.Attr
}

// NewSlogHandler creates a new slog.Handler that writes log records to the
// given logger.Log. The handler maps slog levels to logger levels and
// converts slog attributes to logger fields.
func NewSlogHandler(logger *Log) *SlogHandler {
	return &SlogHandler{logger: logger}
}

// Enabled reports whether the handler handles records at the given level.
// It mirrors Logger.should's level and writer checks (without sampling).
func (h *SlogHandler) Enabled(_ context.Context, level slog.Level) bool {
	return h.logger.Enabled(convertSlogLevel(level))
}

// Handle handles the Record. It converts the slog.Record into a logger event
// and writes it using the underlying logger.Log.
func (h *SlogHandler) Handle(ctx context.Context, record slog.Record) error {
	zlevel := convertSlogLevel(record.Level)
	event := h.logger.OnLevel(zlevel)
	if event == nil {
		return nil
	}

	// Propagate slog context to the logger event so that hooks
	// relying on Event.Context() (e.g. tracing) can access it.
	if ctx != nil {
		event = event.WithContext(ctx)
	}

	// Add pre-attached attrs from WithAttrs
	for _, a := range h.attrs {
		event = appendSlogAttr(event, a, h.prefix)
	}

	// Add attrs from the record itself
	record.Attrs(func(a slog.Attr) bool {
		event = appendSlogAttr(event, a, h.prefix)
		return true
	})
	if !record.Time.IsZero() {
		event.Time(h.timeKey, record.Time)
	}

	event.Msg(record.Message)
	return nil
}

// WithAttrs returns a new Handler with the given attributes pre-attached.
// These attributes will be included in every subsequent log record.
func (h *SlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	h2 := h.clone()
	h2.attrs = append(h2.attrs, attrs...)
	return h2
}

// WithGroup returns a new Handler with the given group name. All subsequent
// attributes will be nested under this group name in the output.
func (h *SlogHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	h2 := h.clone()
	if h2.prefix != "" {
		h2.prefix = h2.prefix + "." + name
	} else {
		h2.prefix = name
	}
	return h2
}

func (h *SlogHandler) clone() *SlogHandler {
	h2 := &SlogHandler{
		logger:  h.logger.With(),
		prefix:  h.prefix,
		timeKey: h.timeKey,
	}
	if len(h.attrs) > 0 {
		h2.attrs = make([]slog.Attr, len(h.attrs))
		copy(h2.attrs, h.attrs)
	}
	return h2
}

// convertSlogLevel maps slog Levels to zap Levels.
// Note that there is some room between slog levels while zap levels are continuous, so we can't 1:1 map them.
// See also https://go.googlesource.com/proposal/+/master/design/56345-structured-logging.md?pli=1#levels
func convertSlogLevel(l slog.Level) Level {
	switch {
	case l >= slog.LevelError:
		return ErrorLevel
	case l >= slog.LevelWarn:
		return WarnLevel
	case l >= slog.LevelInfo:
		return InfoLevel
	default:
		return DebugLevel
	}
}

// joinPrefix concatenates a prefix and key with a dot separator.
// It avoids allocations when either prefix or key is empty.
func joinPrefix(prefix, key string) string {
	if prefix == "" {
		return key
	}
	if key == "" {
		return prefix
	}
	return prefix + "." + key
}

// appendSlogAttr appends a single slog.Attr to the logger event, handling
// type-specific encoding to avoid reflection where possible.
func appendSlogAttr(event *Event, attr slog.Attr, prefix string) *Event {
	if event == nil {
		return event
	}

	// Resolve the attribute to handle LogValuer types.
	// This handles slog.KindLogValuer implicitly by unwrapping
	// any values that implement slog.LogValuer to their resolved form.
	attr.Value = attr.Value.Resolve()

	// For group kinds, handle grouping before key concatenation
	if attr.Value.Kind() == slog.KindGroup {
		attrs := attr.Value.Group()
		if len(attrs) == 0 {
			return event
		}
		groupPrefix := joinPrefix(prefix, attr.Key)
		for _, ga := range attrs {
			event = appendSlogAttr(event, ga, groupPrefix)
		}
		return event
	}

	// Skip empty keys for non-group attributes
	if attr.Key == "" {
		return event
	}

	key := joinPrefix(prefix, attr.Key)
	val := attr.Value

	switch val.Kind() {
	case slog.KindString:
		event = event.String(key, val.String())
	case slog.KindInt64:
		event = event.Int64(key, val.Int64())
	case slog.KindUint64:
		event = event.Uint64(key, val.Uint64())
	case slog.KindFloat64:
		event = event.Float64(key, val.Float64())
	case slog.KindBool:
		event = event.Bool(key, val.Bool())
	case slog.KindDuration:
		event = event.Duration(key, val.Duration())
	case slog.KindTime:
		event = event.Time(key, val.Time())
	case slog.KindAny:
		v := val.Any()
		switch cv := v.(type) {
		case error:
			event = event.NamedError(key, cv)
		case time.Duration:
			event = event.Duration(key, cv)
		case time.Time:
			event = event.Time(key, cv)
		case []byte:
			event = event.ByteString(key, cv)
		default:
			event = event.Any(key, v)
		}
	default:
		event = event.Any(key, val.Any())
	}

	return event
}
