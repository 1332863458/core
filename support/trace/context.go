package trace

import "context"

type key struct{}

var id = key{}

func AppendTracingContext(goCtx context.Context, tracingContext TracingContext) context.Context {
	return context.WithValue(goCtx, id, tracingContext)
}

func ExtractTracingContext(goCtx context.Context) TracingContext {
	tc, _ := goCtx.Value(id).(TracingContext)
	return tc
}

type TracingContext interface {
	// TraceObject() returns underlying tracing implementation
	TraceObject() interface{}
	// SetTags() allows you to set one or more tags to tracing object
	SetTags(tags map[string]interface{}) bool
	// SetTags() allows you to add tag to tracing object
	SetTag(tagKey string, tagValue interface{}) bool
	// LogKV() allows you to log additional details about the entity being traced
	LogKV(kvs map[string]interface{}) bool

	// Inject() injects the current trace context for
	// propagation within `carrier`. The actual type of `carrier` depends on the value of `format`.
	//
	// trace.Tracer defines a common set of `format` values, and each has an expected `carrier` type.
	//
	// Example usage in activity:
	//
	// tc := ctx.GetTracingContext()
	// err = tc.Inject(trace.HTTPHeaders, req)
	Inject(format CarrierFormat, carrier interface{}) error
}



type Config struct {
	Operation string
	Tags      map[string]interface{}
}