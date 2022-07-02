package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"os"
	"strings"
)

const policy = "./policies/example.rego"
const query = "data.site.allow = true"

func main() {

	policyContent, err := os.ReadFile(policy)
	if err != nil {
		panic(err)
	}

	r := rego.New(
		rego.Query(query),
		rego.Module("example.rego", string(policyContent)),
		rego.Input(map[string]any{
			"method":  "GET",
			"ownerId": "OID1234",
			"path":    []string{"opinions", "OPID12345"},
		}),
		rego.Unknowns([]string{"data.opinions"}),
	)

	resp, err := r.Partial(context.Background())
	if err != nil {
		panic(err)
	}

	body, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	_, err = ParsePartialRawResult(body)
	if err != nil {
		panic(err)
	}
}

var (
	AccessDeniedError                 = errors.New("access denied")
	InvalidPolicyQueryTermError       = errors.New("invalid policy term")
	InvalidPolicyQueryExpressionError = errors.New("invalid policy expression")
)

func ParsePartialRawResult(result []byte) (any, error) {
	var r rego.PartialQueries
	if err := json.Unmarshal(result, &r); err != nil {
		return nil, err
	}

	if err := IsAccessDenied(r); err != nil {
		return nil, err
	}

	if IsUnconditional(r.Queries) {
		return nil, nil
	}

	for _, body := range r.Queries {
		for _, expr := range body {
			parsed, err := ParseExpression(expr)
			if err != nil {
				return nil, err
			}
			fmt.Println(parsed)
		}
	}

	return nil, nil
}

// IsUnconditional checks if the given queries contains an unconditional query.
// Unconditional queries could be threaten like a super admin access
// https://www.openpolicyagent.org/docs/latest/rest-api/#unconditional-results-from-partial-evaluation
func IsUnconditional(queries []ast.Body) bool {
	for _, q := range queries {
		if len(q) == 0 {
			return true
		}
	}
	return false
}

// IsAccessDenied checks if the query response states access denied error
// https://www.openpolicyagent.org/docs/latest/rest-api/#unconditional-results-from-partial-evaluation
func IsAccessDenied(q rego.PartialQueries) error {
	if q.Queries == nil {
		return AccessDeniedError
	}
	return nil
}

type ParsedExpression struct {
	FieldName string
	Value     any
	Operator  string
}

func ParseExpression(e *ast.Expr) (ParsedExpression, error) {
	pe := ParsedExpression{
		Operator: e.Operator().String(),
	}

	if !e.IsCall() {
		return ParsedExpression{}, InvalidPolicyQueryExpressionError
	}

	if len(e.Operands()) != 2 {
		return ParsedExpression{}, InvalidPolicyQueryExpressionError
	}

	for _, term := range e.Operands() {
		// todo refactor in own method
		if ast.IsConstant(term.Value) {
			val, err := ast.JSON(term.Value)
			if err != nil {
				return ParsedExpression{}, err
			}
			pe.Value = val
			continue
		}

		fieldName, err := ParseTerm(term.String())
		if err != nil {
			return ParsedExpression{}, err
		}
		pe.FieldName = fieldName
	}

	if pe.Value == nil {
		return ParsedExpression{}, InvalidPolicyQueryExpressionError
	}
	return pe, nil
}

func ParseTerm(term string) (string, error) {
	split := strings.Split(term, ".")
	splitLen := len(split)

	// TODO: nachdeken wie das aussehen soll: strict nur 3 erlauben data.tablename.property, wenn ja sollen table names gespeichert werden?
	if splitLen < 3 {
		return "", InvalidPolicyQueryTermError
	}
	return split[splitLen-1], nil
}
