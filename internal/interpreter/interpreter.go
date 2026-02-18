// Package interpreter defines the core Interpreter interface for the World language.
//
// An Interpreter executes World source code and returns an error if execution fails.
// Different implementations may interpret code using different strategies,
// such as direct string interpretation, tokenization, or abstract syntax trees.
//
// This abstraction allows World to support multiple interpreter implementations
// while maintaining a consistent execution contract.
package interpreter

type Interpreter interface {
	Run(code string) error
}
