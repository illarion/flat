package flat

import "testing"

func TestFlatten(t *testing.T) {
	cases := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "simple",
			input: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"t": 123,
					},
				},
			},
			expected: map[string]interface{}{
				"foo.bar.t": 123,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := Flatten(c.input, nil)
			if !mapsEqual(output, c.expected) {
				t.Errorf("expected %v, got %v", c.expected, output)
			}
		})
	}
}

func TestUnflatten(t *testing.T) {
	cases := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "simple",
			input: map[string]interface{}{
				"foo.bar.t": 123,
			},
			expected: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"t": 123,
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := Unflatten(c.input, nil)
			if !mapsEqual(output, c.expected) {
				t.Errorf("expected %v, got %v", c.expected, output)
			}
		})
	}
}

func TestFlattenThenUnflatten(t *testing.T) {
	cases := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "simple",
			input: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"t": 123,
					},
				},
			},
			expected: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": map[string]interface{}{
						"t": 123,
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			flat := Flatten(c.input, nil)
			output := Unflatten(flat, nil)
			if !mapsEqual(output, c.expected) {
				t.Errorf("expected %v, got %v", c.expected, output)
			}
		})
	}

}

func mapsEqual(actual map[string]interface{}, expected map[string]interface{}) bool {
	if len(actual) != len(expected) {
		return false
	}

	for k, v := range actual {
		am, aok := v.(map[string]interface{})
		em, eok := expected[k].(map[string]interface{})
		if aok && eok {
			return mapsEqual(am, em)
		}
		if aok != eok {
			return false
		}
		if v != expected[k] {
			return false
		}
	}

	return true
}
