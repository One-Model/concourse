package vars

import (
	"fmt"
	"strings"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Variables
type Variables interface {
	Get(Reference) (interface{}, bool, error)
	List() ([]Reference, error)
}

type Reference struct {
	Source string
	Path   string
}

func NewReference(source, path string) Reference {
	return Reference{Source: source, Path: path}
}

func NewReferenceWithoutSource(path string) Reference {
	return Reference{Path: path}
}

func (r Reference) WithoutSource() Reference {
	return NewReferenceWithoutSource(r.Path)
}

func (r Reference) String() string {
	var s strings.Builder
	if r.Source != "" {
		s.WriteString(r.Source + ":")
	}

	s.WriteString(refSegmentString(r.Path))
	return s.String()
}

func refSegmentString(seg string) string {
	if strings.ContainsAny(seg, ",.:/ ") {
		return fmt.Sprintf("%q", seg)
	}
	return seg
}

type FieldReference struct {
	Reference
	Fields []string
}

func NewFieldReference(source, path string, fields []string) FieldReference {
	return FieldReference{Reference: NewReference(source, path), Fields: fields}
}

func NewFieldReferenceWithoutSource(path string, fields []string) FieldReference {
	return FieldReference{Reference: NewReferenceWithoutSource(path), Fields: fields}
}

func (r FieldReference) WithoutSource() FieldReference {
	return NewFieldReferenceWithoutSource(r.Path, r.Fields)
}

func (r FieldReference) String() string {
	var s strings.Builder
	s.WriteString(r.Reference.String())

	fields := r.Fields
	for len(fields) > 0 {
		s.WriteRune('.')
		s.WriteString(refSegmentString(fields[0]))
		fields = fields[1:]
	}

	return s.String()
}

func Traverse(val interface{}, ref FieldReference) (interface{}, error) {
	for _, seg := range ref.Fields {
		switch v := val.(type) {
		case map[interface{}]interface{}:
			var found bool
			val, found = v[seg]
			if !found {
				return nil, MissingFieldError{
					Name:  ref.String(),
					Field: seg,
				}
			}

		case map[string]interface{}:
			var found bool
			val, found = v[seg]
			if !found {
				return nil, MissingFieldError{
					Name:  ref.String(),
					Field: seg,
				}
			}

		default:
			return nil, InvalidFieldError{
				Name:  ref.String(),
				Field: seg,
				Value: val,
			}
		}
	}
	return val, nil
}

func ParseReference(name string) (FieldReference, error) {
	var ref FieldReference

	input := name
	if i, ok := findUnquoted(input, ':'); ok {
		ref.Source = input[:i]
		if strings.ContainsAny(ref.Source, `"`) {
			return FieldReference{}, fmt.Errorf("invalid var '%s': source must not be quoted", name)
		}
		input = input[i+1:]
	}

	var fields []string
	hasNextSegment := true
	for hasNextSegment {
		var field string
		field, input, hasNextSegment = readPathSegment(input)
		if field == "" {
			return FieldReference{}, fmt.Errorf("invalid var '%s': empty field", name)
		}
		fields = append(fields, field)
	}

	if len(fields) == 0 {
		// Should be impossible (since we'd error that the var is empty), but better safe than sorry
		return FieldReference{}, fmt.Errorf("invalid var '%s': no fields", name)
	}

	ref.Path = fields[0]
	ref.Fields = fields[1:]

	return ref, nil
}

func findUnquoted(s string, r rune) (int, bool) {
	quoted := false
	for i, c := range s {
		switch c {
		case r:
			if !quoted {
				return i, true
			}
		case '"':
			quoted = !quoted
		}
	}
	return 0, false
}

func readPathSegment(raw string) (string, string, bool) {
	var field string
	var rest string
	i, hasNextSegment := findUnquoted(raw, '.')
	if hasNextSegment {
		field = raw[:i]
		rest = raw[i+1:]
	} else {
		field = raw
	}
	field = strings.TrimSpace(field)
	field = strings.ReplaceAll(field, `"`, "")
	return field, rest, hasNextSegment
}
