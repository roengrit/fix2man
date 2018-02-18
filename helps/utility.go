package helps

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//ValidateDate _
func ValidateDate(reqDate string) (time.Time, error) {
	var errDate error
	retDate := time.Now()
	sp := strings.Split(reqDate, "-")
	if len(sp) == 3 {
		retDate, errDate = time.Parse("2006-02-01", sp[2]+"-"+sp[0]+"-"+sp[1])
	}
	return retDate, errDate
}

//ThCommaSep _
func ThCommaSep(in float64) (out string) {
	out = fmt.Sprintf("%s", RenderFloat("#,###.##", in))
	return
}

//RowOrderInc _
func RowOrderInc(in int) (out string) {
	out = strconv.Itoa(in + 1)
	return
}

//TextThCommaSep _
func TextThCommaSep(in string) (out string) {
	val, _ := strconv.ParseFloat(in, 64)
	out = fmt.Sprintf("%s", RenderFloat("#,###.##", val))
	return
}

var renderFloatPrecisionMultipliers = [10]float64{
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000,
}

var renderFloatPrecisionRounders = [10]float64{
	0.5,
	0.05,
	0.005,
	0.0005,
	0.00005,
	0.000005,
	0.0000005,
	0.00000005,
	0.000000005,
	0.0000000005,
}

//RenderFloat _
func RenderFloat(format string, n float64) string {
	// Special cases:
	// NaN = "NaN"
	// +Inf = "+Infinity"
	// -Inf = "-Infinity"
	if math.IsNaN(n) {
		return "NaN"
	}
	if n > math.MaxFloat64 {
		return "Infinity"
	}
	if n < -math.MaxFloat64 {
		return "-Infinity"
	}

	// default format
	precision := 2
	decimalStr := "."
	thousandStr := ","
	positiveStr := ""
	negativeStr := "-"

	if len(format) > 0 {
		// If there is an explicit format directive,
		// then default values are these:
		precision = 9
		thousandStr = ""

		// collect indices of meaningful formatting directives
		formatDirectiveChars := []rune(format)
		formatDirectiveIndices := make([]int, 0)
		for i, char := range formatDirectiveChars {
			if char != '#' && char != '0' {
				formatDirectiveIndices = append(formatDirectiveIndices, i)
			}
		}

		if len(formatDirectiveIndices) > 0 {
			// Directive at index 0:
			// Must be a '+'
			// Raise an error if not the case
			// index: 0123456789
			// +0.000,000
			// +000,000.0
			// +0000.00
			// +0000
			if formatDirectiveIndices[0] == 0 {
				if formatDirectiveChars[formatDirectiveIndices[0]] != '+' {
					panic("RenderFloat(): invalid positive sign directive")
				}
				positiveStr = "+"
				formatDirectiveIndices = formatDirectiveIndices[1:]
			}

			// Two directives:
			// First is thousands separator
			// Raise an error if not followed by 3-digit
			// 0123456789
			// 0.000,000
			// 000,000.00
			if len(formatDirectiveIndices) == 2 {
				if (formatDirectiveIndices[1] - formatDirectiveIndices[0]) != 4 {
					panic("RenderFloat(): thousands separator directive must be followed by 3 digit-specifiers")
				}
				thousandStr = string(formatDirectiveChars[formatDirectiveIndices[0]])
				formatDirectiveIndices = formatDirectiveIndices[1:]
			}

			// One directive:
			// Directive is decimal separator
			// The number of digit-specifier following the separator indicates wanted precision
			// 0123456789
			// 0.00
			// 000,0000
			if len(formatDirectiveIndices) == 1 {
				decimalStr = string(formatDirectiveChars[formatDirectiveIndices[0]])
				precision = len(formatDirectiveChars) - formatDirectiveIndices[0] - 1
			}
		}
	}

	// generate sign part
	var signStr string
	if n >= 0.000000001 {
		signStr = positiveStr
	} else if n <= -0.000000001 {
		signStr = negativeStr
		n = -n
	} else {
		signStr = ""
		n = 0.0
	}

	// split number into integer and fractional parts
	intf, fracf := math.Modf(n + renderFloatPrecisionRounders[precision])

	// generate integer part string
	intStr := strconv.Itoa(int(intf))

	// add thousand separator if required
	if len(thousandStr) > 0 {
		for i := len(intStr); i > 3; {
			i -= 3
			intStr = intStr[:i] + thousandStr + intStr[i:]
		}
	}

	// no fractional part, we can leave now
	if precision == 0 {
		return signStr + intStr
	}

	// generate fractional part
	fracStr := strconv.Itoa(int(fracf * renderFloatPrecisionMultipliers[precision]))
	// may need padding
	if len(fracStr) < precision {
		fracStr = "000000000000000"[:precision-len(fracStr)] + fracStr
	}

	return signStr + intStr + decimalStr + fracStr
}

//RenderInteger _
func RenderInteger(format string, n int) string {
	return RenderFloat(format, float64(n))
}

//File64Encode _
func File64Encode(path string) (string, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buff), nil
}

//RemoveContents _
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

//RemoveContentsExcludeFile _
func RemoveContentsExcludeFile(dir string, exFilename string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		if exFilename != name {
			err = os.RemoveAll(filepath.Join(dir, name))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
