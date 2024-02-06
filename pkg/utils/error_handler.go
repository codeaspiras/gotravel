package utils

import (
	"gotravel/pkg/errs"
	"gotravel/pkg/stdio"
)

func ErrorHandler(io stdio.IO, err error) {
	if err == nil {
		return
	}

	switch err {
	case errs.ErrCostPerLiterCannotBeEqualOrLowerThanZero, errs.ErrDistanceCannotBeEqualOrLowerThanZero, errs.ErrDistancePerLiterCannotBeEqualOrLowerThanZero:
		io.Echo(err.Error())
	default:
		io.Echo("# Isso gerou um erro:\n- %s", err)
	}

	io.End()
}
