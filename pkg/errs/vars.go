package errs

import "errors"

var (
	// ErrEmptyCtx happens when it's not possible to extract something from value because it's not filled
	ErrEmptyCtx = errors.New("empty context")
	// ErrInvalidCtxValue happens when the context.Context.Value returns something not castable to the what it should be
	ErrInvalidCtxValue = errors.New("invalid context")
	// ErrCostPerLiterCannotBeEqualOrLowerThanZero happens when costPerLiter <= 0
	ErrCostPerLiterCannotBeEqualOrLowerThanZero = errors.New("Se está de graça, essa calculadora não tem serventia alguma...")
	// ErrDistanceCannotBeEqualOrLowerThanZero happens when distance <= 0
	ErrDistanceCannotBeEqualOrLowerThanZero = errors.New("Ora, então você não vai ter custo. Pra quê fazer conta?")
	// ErrDistancePerLiterCannotBeEqualOrLowerThanZero happens when distancePerLiter <= 0
	ErrDistancePerLiterCannotBeEqualOrLowerThanZero = errors.New("Isso é impossível. Seu automóvel é do futuro?")
)
