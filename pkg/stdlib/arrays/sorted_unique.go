package arrays

import (
	"context"

	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
)

// SortedUnique sorts all elements in anyArray.
// The function will use the default comparison order for FQL value types.
// Additionally, the values in the result array will be made unique
// @param array (Array) - Target array.
// @returns (Array) - Sorted array.
func SortedUnique(_ context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 1, 1)

	if err != nil {
		return values.None, err
	}

	err = core.ValidateType(args[0], core.ArrayType)

	if err != nil {
		return values.None, err
	}

	arr := args[0].(*values.Array)

	if arr.Length() == 0 {
		return values.NewArray(0), nil
	}

	return ToUniqueArray(arr.Sort()), nil
}
