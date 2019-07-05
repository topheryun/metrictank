package tagquery

import (
	"strings"
)

type expressionMatchNone struct {
	// we keep key, operator, value just to be able to convert the expression back into a string
	expressionCommon
	originalOperator ExpressionOperator
}

func (e *expressionMatchNone) GetKey() string {
	return e.key
}

func (e *expressionMatchNone) GetValue() string {
	return e.value
}

func (e *expressionMatchNone) RequiresNonEmptyValue() bool {
	return true
}

func (e *expressionMatchNone) OperatesOnTag() bool {
	return false
}

func (e *expressionMatchNone) HasRe() bool {
	return false
}
func (e *expressionMatchNone) GetOperator() ExpressionOperator {
	return MATCH_NONE
}

func (e *expressionMatchNone) ValuePasses(value string) bool {
	return false
}

func (e *expressionMatchNone) GetDefaultDecision() FilterDecision {
	return Fail
}

func (e *expressionMatchNone) StringIntoBuilder(builder *strings.Builder) {
	builder.WriteString(e.key)
	e.originalOperator.StringIntoBuilder(builder)
	builder.WriteString(e.value)
}

func (e *expressionMatchNone) GetMetricDefinitionFilter() MetricDefinitionFilter {
	return func(_ string, _ []string) FilterDecision { return Fail }
}