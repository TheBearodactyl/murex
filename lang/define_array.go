package lang

import (
	"context"
	"fmt"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils"
)

// ArrayTemplate is a template function for reading arrays from marshalled data
func ArrayTemplate(ctx context.Context, marshal func(any) ([]byte, error), unmarshal func([]byte, any) error, read stdio.Io, callback func([]byte)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	if len(utils.CrLfTrim(b)) == 0 {
		return nil
	}

	var v interface{}
	err = unmarshal(b, &v)
	if err != nil {
		return err
	}

	return ArrayDataTemplate(ctx, marshal, unmarshal, v, callback)
}

func ArrayDataTemplate(ctx context.Context, marshal func(any) ([]byte, error), unmarshal func([]byte, any) error, data any, callback func([]byte)) error {
	switch v := data.(type) {
	case string:
		return readArrayByString(v, callback)

	case []string:
		return readArrayBySliceString(ctx, v, callback)

	case []interface{}:
		return readArrayBySliceInterface(ctx, marshal, v, callback)

	case map[string]string:
		return readArrayByMapStrStr(ctx, v, callback)

	case map[string]interface{}:
		return readArrayByMapStrIface(ctx, marshal, v, callback)

	case map[interface{}]string:
		return readArrayByMapIfaceStr(ctx, v, callback)

	case map[interface{}]interface{}:
		return readArrayByMapIfaceIface(ctx, marshal, v, callback)

	default:
		jBytes, err := marshal(v)
		if err != nil {

			return err
		}

		callback(jBytes)

		return nil
	}
}

func readArrayByString(v string, callback func([]byte)) error {
	callback([]byte(v))

	return nil
}

func readArrayBySliceString(ctx context.Context, v []string, callback func([]byte)) error {
	for i := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			callback([]byte(v[i]))
		}
	}

	return nil
}

func readArrayBySliceInterface(ctx context.Context, marshal func(interface{}) ([]byte, error), v []interface{}, callback func([]byte)) error {
	if len(v) == 0 {
		return nil
	}

	for i := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			switch v := v[i].(type) {
			case string:
				callback([]byte(v))

			case []byte:
				callback(v)

			default:
				jBytes, err := marshal(v)
				if err != nil {
					return err
				}
				callback(jBytes)
			}
		}
	}

	return nil
}

func readArrayByMapIfaceIface(ctx context.Context, marshal func(interface{}) ([]byte, error), v map[interface{}]interface{}, callback func([]byte)) error {
	for key, val := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			bKey := []byte(fmt.Sprint(key) + ": ")
			b, err := marshal(val)
			if err != nil {
				return err
			}
			callback(append(bKey, b...))
		}
	}

	return nil
}

func readArrayByMapStrStr(ctx context.Context, v map[string]string, callback func([]byte)) error {
	for key, val := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			callback([]byte(key + ": " + val))
		}
	}

	return nil
}

func readArrayByMapStrIface(ctx context.Context, marshal func(interface{}) ([]byte, error), v map[string]interface{}, callback func([]byte)) error {
	for key, val := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			bKey := []byte(key + ": ")
			b, err := marshal(val)
			if err != nil {
				return err
			}
			callback(append(bKey, b...))
		}
	}

	return nil
}

func readArrayByMapIfaceStr(ctx context.Context, v map[interface{}]string, callback func([]byte)) error {
	for key, val := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			callback([]byte(fmt.Sprint(key) + ": " + val))
		}
	}

	return nil
}
