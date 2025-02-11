package conversions

import (
	"errors"
	"fmt"
	"math"
)

// SafeIntConversion convierte de manera segura de int64 a int32, verificando si id excede el rango de int.
func SafeIntConversion(id int64) (int, error) {
	if id > int64(math.MaxInt) || id < int64(math.MinInt) {
		return 0, errors.New(fmt.Sprintf("el valor %d excede el rango de int", id))
	}
	return int(id), nil
}
