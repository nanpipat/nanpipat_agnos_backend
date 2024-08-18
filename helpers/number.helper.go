package helpers

import (
	"fmt"
	"strconv"
)

func ConvertByteToGB(byt int64) float64 {
	return float64(byt) / 1073741824
}

func ConvertByteToTB(byt int64) float64 {
	return float64(byt) / (1024 * 1024 * 1024 * 1024)
}

func ConvertGBtoTB(gb float64) float64 {
	return gb / 1024
}

func ConvertFloat64ToInt64(f float64) int64 {
	t := int64(f)
	s := fmt.Sprintf("%.0f", f)
	if i, err := strconv.Atoi(s); err == nil {
		fmt.Println(f, t, i)
		return int64(i)
	} else {
		fmt.Println(f, t, err)
	}

	return t
}
