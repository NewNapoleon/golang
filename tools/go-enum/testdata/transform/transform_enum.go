// Code generated by "go-enum -type Transform -trimprefix=Transform -transform=lower"; DO NOT EDIT.

package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TransformOne-0]
	_ = x[TransformTwo-1]
	_ = x[TransformThree-2]
}

const _Transform_name = "onetwothree"

var _Transform_index = [...]uint8{0, 3, 6, 11}

func _() {
	var _nil_Transform_value = func() (val Transform) { return }()

	// An "cannot convert Transform literal (type Transform) to type fmt.Stringer" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ fmt.Stringer = _nil_Transform_value
}

func (i Transform) String() string {
	if i < 0 || i >= Transform(len(_Transform_index)-1) {
		return "Transform(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Transform_name[_Transform_index[i]:_Transform_index[i+1]]
}

var _Transform_values = []Transform{0, 1, 2}

var _Transform_name_to_values = map[string]Transform{
	_Transform_name[0:3]:  0,
	_Transform_name[3:6]:  1,
	_Transform_name[6:11]: 2,
}

// ParseTransformString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ParseTransformString(s string) (Transform, error) {
	if val, ok := _Transform_name_to_values[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%[1]s does not belong to Transform values", s)
}

// TransformValues returns all values of the enum
func TransformValues() []Transform {
	return _Transform_values
}

// IsATransform returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Transform) Registered() bool {
	for _, v := range _Transform_values {
		if i == v {
			return true
		}
	}
	return false
}

func _() {
	var _nil_Transform_value = func() (val Transform) { return }()

	// An "cannot convert Transform literal (type Transform) to type encoding.BinaryMarshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ encoding.BinaryMarshaler = &_nil_Transform_value

	// An "cannot convert Transform literal (type Transform) to type encoding.BinaryUnmarshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ encoding.BinaryUnmarshaler = &_nil_Transform_value
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for Transform
func (i Transform) MarshalBinary() (data []byte, err error) {
	return []byte(i.String()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for Transform
func (i *Transform) UnmarshalBinary(data []byte) error {
	var err error
	*i, err = ParseTransformString(string(data))
	return err
}

func _() {
	var _nil_Transform_value = func() (val Transform) { return }()

	// An "cannot convert Transform literal (type Transform) to type json.Marshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ json.Marshaler = _nil_Transform_value

	// An "cannot convert Transform literal (type Transform) to type encoding.Unmarshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ json.Unmarshaler = &_nil_Transform_value
}

// MarshalJSON implements the json.Marshaler interface for Transform
func (i Transform) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Transform
func (i *Transform) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Transform should be a string, got %[1]s", data)
	}

	var err error
	*i, err = ParseTransformString(s)
	return err
}

func _() {
	var _nil_Transform_value = func() (val Transform) { return }()

	// An "cannot convert Transform literal (type Transform) to type encoding.TextMarshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ encoding.TextMarshaler = _nil_Transform_value

	// An "cannot convert Transform literal (type Transform) to type encoding.TextUnmarshaler" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ encoding.TextUnmarshaler = &_nil_Transform_value
}

// MarshalText implements the encoding.TextMarshaler interface for Transform
func (i Transform) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Transform
func (i *Transform) UnmarshalText(text []byte) error {
	var err error
	*i, err = ParseTransformString(string(text))
	return err
}

//func _() {
//	var _nil_Transform_value = func() (val Transform) { return }()
//
//	// An "cannot convert Transform literal (type Transform) to type yaml.Marshaler" compiler error signifies that the base type have changed.
//	// Re-run the go-enum command to generate them again.
//	var _ yaml.Marshaler = _nil_Transform_value
//
//	// An "cannot convert Transform literal (type Transform) to type yaml.Unmarshaler" compiler error signifies that the base type have changed.
//	// Re-run the go-enum command to generate them again.
//	var _ yaml.Unmarshaler = &_nil_Transform_value
//}

// MarshalYAML implements a YAML Marshaler for Transform
func (i Transform) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Transform
func (i *Transform) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = ParseTransformString(s)
	return err
}

func _() {
	var _nil_Transform_value = func() (val Transform) { return }()

	// An "cannot convert Transform literal (type Transform) to type driver.Valuer" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ driver.Valuer = _nil_Transform_value

	// An "cannot convert Transform literal (type Transform) to type sql.Scanner" compiler error signifies that the base type have changed.
	// Re-run the go-enum command to generate them again.
	var _ sql.Scanner = &_nil_Transform_value
}

func (i Transform) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Transform) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := ParseTransformString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}

// TransformSliceContains reports whether sunEnums is within enums.
func TransformSliceContains(enums []Transform, sunEnums ...Transform) bool {
	var seenEnums = map[Transform]bool{}
	for _, e := range sunEnums {
		seenEnums[e] = false
	}

	for _, v := range enums {
		if _, has := seenEnums[v]; has {
			seenEnums[v] = true
		}
	}

	for _, seen := range seenEnums {
		if !seen {
			return false
		}
	}

	return true
}

// TransformSliceContainsAny reports whether any sunEnum is within enums.
func TransformSliceContainsAny(enums []Transform, sunEnums ...Transform) bool {
	var seenEnums = map[Transform]struct{}{}
	for _, e := range sunEnums {
		seenEnums[e] = struct{}{}
	}

	for _, v := range enums {
		if _, has := seenEnums[v]; has {
			return true
		}
	}

	return false
}
