package typemgmt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func init() {
	lang.DefineMethod("datetime", cmdDateTime, types.Any, types.Any)

	defaults.AppendProfile(`	
		autocomplete: set datetime {
			[{
				"FlagsDesc": {
					"--value": "(optional) Input data/time string to be parsed",
					"--in": "Formatting rules of input data/time string",
					"--out": "Formatting rules of output date/time string"
				},
				"AllowMultiple": true,
				"AnyValue": true
			}]
		}
	`)
}

func cmdDateTime(p *lang.Process) error {
	flags, _, err := p.Parameters.ParseFlags(&parameters.Arguments{
		Flags: map[string]string{
			"--in":    types.String,
			"--out":   types.String,
			"--value": types.String,
		},
		AllowAdditional: false,
	})
	if err != nil {
		return err
	}

	if flags["--value"] == "" && flags["--in"] != "{now}" {
		if err := p.ErrIfNotAMethod(); err != nil {
			return err
		}

		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		flags["--value"] = string(utils.CrLfTrim(b))
	}

	in := flags["--in"]
	out := flags["--out"]

	if in == "" {
		return errors.New("no datatime format provided for input value: `--in` required")
	}

	if out == "" {
		return errors.New("no datatime format provided for output value: `--out` required")
	}

	var datetime time.Time

	// Parse --in

	switch {
	case strings.HasPrefix(in, "{go}"):
		datetime, err = time.Parse(in[4:], flags["--value"])
		if err != nil {
			return err
		}

	case strings.HasPrefix(in, "{py}"):
		/*datetime, err = time.Parse(in[4:], flags["--value"])
		if err != nil {
			return err
		}*/
		return errors.New("TODO! This feature hasn't yet been developed")

	case strings.HasPrefix(in, "{now}"):
		datetime = time.Now()

	default:
		return errors.New("unknown or invalid input parser formatter, expecting `{go}`, `{py}` or `{now}`")
	}

	// Write --out

	switch {
	case strings.HasPrefix(out, "{go}"):
		_, err = p.Stdout.Write([]byte(datetime.Format(out[4:])))
		return err

	case strings.HasPrefix(out, "{py}"):
		s, err := pyParse(out[4:], datetime)
		if err != nil {
			return err
		}
		_, err = p.Stdout.Write([]byte(s))
		return err

	case strings.HasPrefix(out, "{unix}"):
		_, err = p.Stdout.Write([]byte(strconv.FormatInt(datetime.Unix(), 10)))
		return err

	default:
		return errors.New("unknown or invalid output parser formatter, expecting `{go}`, `{py}`, or `{unix}`")
	}

}

func pyParse(format string, t time.Time) (string, error) {
	var (
		isDirective bool
		s           string
	)

	for _, c := range format {
		if !isDirective {
			if c == '%' {
				isDirective = true
			} else {
				s += string(c)
			}
			continue
		}

		isDirective = false

		switch c {
		case 'a':
			s += t.Format("Mon")
		case 'A':
			s += t.Format("Monday")
		case 'w':
			switch t.Format("Monday") {
			case "Sunday":
				s += "0"
			case "Monday":
				s += "1"
			case "Tuesday":
				s += "2"
			case "Wednesday":
				s += "3"
			case "Thursday":
				s += "4"
			case "Friday":
				s += "5"
			case "Saturday":
				s += "6"
			}
		case 'd':
			s += t.Format("02")
		case 'b':
			s += t.Format("Jan")
		case 'B':
			s += t.Format("January")
		case 'm':
			s += t.Format("01")
		case 'y':
			s += t.Format("06")
		case 'Y':
			s += t.Format("2006")
		case 'H':
			s += t.Format("15")
		case 'I':
			s += t.Format("03")
		case 'p':
			s += t.Format("PM")
		case 'M':
			s += t.Format("04")
		case 'S':
			s += t.Format("05")
		case 'f':
			s += t.Format("000000")
		case 'z':
			s += t.Format("-0700")
		case 'Z':
			s += t.Format("MST")
		case 'j':
			s += t.Format("002")
		case 'U':
			return "", errors.New("`%U` is currently unsupported")
		case 'W':
			return "", errors.New("`%W` is currently unsupported")
		case 'c':
			s += t.Format("Mon Jan 03 15:04:05 2006")
		case 'x':
			s += t.Format("01/02/06")
		case 'X':
			s += t.Format("15:04:05")
		case '%':
			s += "%"
		default:
			return "", fmt.Errorf("`%%%s` is not a valid Python datetime directive", string(c))
		}
	}

	return s, nil
}
