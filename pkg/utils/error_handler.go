package utils

import (
	"gotravel/pkg/errs"
	"gotravel/pkg/stdio"
)

func ErrorHandler(err error) {
	if err == nil {
		return
	}

	switch err {
	case errs.ErrCostPerLiterCannotBeEqualOrLowerThanZero, errs.ErrDistanceCannotBeEqualOrLowerThanZero, errs.ErrDistancePerLiterCannotBeEqualOrLowerThanZero:
		stdio.Echo(err.Error())
	default:
		stdio.Echo("# Isso gerou um erro:\n- %s", err)
	}

	End()
}
