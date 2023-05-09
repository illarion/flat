package flat

import (
	"strconv"
	"strings"
)

type Options struct {
	Delimiter string
}

// Flatten flattens a nested map[string]interface{} into a one level map[string]interface{}.
// The keys of the resulting map will be the path to the values in the nested map.
func Flatten(input map[string]interface{}, opts *Options) map[string]interface{} {

	if opts == nil {
		opts = &Options{
			Delimiter: ".",
		}
	}

	output := make(map[string]interface{})
	flattenHelper(output, input, "", *opts)
	return output
}

func flattenHelper(output map[string]interface{}, input interface{}, prefix string, opts Options) {
	switch val := input.(type) {
	case map[string]interface{}:
		for k, v := range val {
			newKey := k
			if prefix != "" {
				newKey = prefix + opts.Delimiter + k
			}
			flattenHelper(output, v, newKey, opts)
		}
	case []interface{}:
		for i, v := range val {
			newKey := strconv.Itoa(i)
			if prefix != "" {
				newKey = prefix + opts.Delimiter + newKey
			}
			flattenHelper(output, v, newKey, opts)
		}
	default:
		output[prefix] = input
	}

}

func Unflatten(input map[string]interface{}, opts *Options) map[string]interface{} {

	if opts == nil {
		opts = &Options{
			Delimiter: ".",
		}
	}

	output := make(map[string]interface{})
	for key, value := range input {
		parts := strings.Split(key, opts.Delimiter)
		currMap := output
		for i := 0; i < len(parts)-1; i++ {
			part := parts[i]
			nextPart := parts[i+1]
			if _, ok := currMap[part]; !ok {
				currMap[part] = make(map[string]interface{})
			}
			currMap = currMap[part].(map[string]interface{})
			if _, ok := currMap[nextPart]; !ok {
				currMap[nextPart] = make(map[string]interface{})
			}
		}
		currMap[parts[len(parts)-1]] = value
	}
	return output
}
