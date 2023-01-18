package vars

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/hashicorp/go-multierror"
	"gopkg.in/yaml.v2"
)

type Template struct {
	bytes []byte
}

type EvaluateOpts struct {
	ExpectAllKeys     bool
	ExpectAllVarsUsed bool
}

func NewTemplate(bytes []byte) Template {
	return Template{bytes: bytes}
}

func (t Template) ExtraVarNames() []string {
	return interpolator{}.extractVarNames(string(t.bytes))
}

func (t Template) Evaluate(vars Variables, opts EvaluateOpts) ([]byte, error) {
	var obj interface{}

	// Note: if we do end up changing from "gopkg.in/yaml.v2" to
	// "sigs.k8s.io/yaml" here, we'll want to ensure we call
	// `json.Decoder.UseNumber()` so that we don't lose precision unmarshaling
	// numbers to float64.
	err := yaml.Unmarshal(t.bytes, &obj)
	if err != nil {
		return []byte{}, err
	}

	obj, err = t.interpolateRoot(obj, NewVarsTracker(vars, opts.ExpectAllKeys, opts.ExpectAllVarsUsed))
	if err != nil {
		return []byte{}, err
	}

	bytes, err := yaml.Marshal(obj)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}

func (t Template) interpolateRoot(obj interface{}, tracker VarsTracker) (interface{}, error) {
	var err error
	obj, err = interpolator{}.Interpolate(obj, tracker)
	if err != nil {
		return nil, err
	}

	return obj, tracker.Error()
}

type interpolator struct{}

var (
	interpolationRegex         = regexp.MustCompile(`\(\((([-/\.\w\pL]+\:)?[-/\.:@"\w\pL]+)\)\)`)
	interpolationAnchoredRegex = regexp.MustCompile("\\A" + interpolationRegex.String() + "\\z")
)

func (i interpolator) Interpolate(node interface{}, tracker VarsTracker) (interface{}, error) {
	switch typedNode := node.(type) {
	case map[interface{}]interface{}:
		for k, v := range typedNode {
			evaluatedValue, err := i.Interpolate(v, tracker)
			if err != nil {
				return nil, err
			}

			evaluatedKey, err := i.Interpolate(k, tracker)
			if err != nil {
				return nil, err
			}

			delete(typedNode, k) // delete in case key has changed
			typedNode[evaluatedKey] = evaluatedValue
		}

	case []interface{}:
		for idx, x := range typedNode {
			var err error
			typedNode[idx], err = i.Interpolate(x, tracker)
			if err != nil {
				return nil, err
			}
		}

	case string:
		for _, name := range i.extractVarNames(typedNode) {
			varRef, err := ParseReference(name)
			if err != nil {
				return nil, err
			}

			foundVal, found, err := tracker.Get(varRef)
			if err != nil {
				return nil, err
			}

			if found {
				// ensure that value type is preserved when replacing the entire field
				if interpolationAnchoredRegex.MatchString(typedNode) {
					return foundVal, nil
				}

				switch foundVal.(type) {
				case string, int, int16, int32, int64, uint, uint16, uint32, uint64, json.Number:
					foundValStr := fmt.Sprintf("%v", foundVal)
					typedNode = strings.Replace(typedNode, fmt.Sprintf("((%s))", name), foundValStr, -1)
				default:
					return nil, InvalidInterpolationError{
						Name:  name,
						Value: foundVal,
					}
				}
			}
		}

		return typedNode, nil
	}

	return node, nil
}

func (i interpolator) extractVarNames(value string) []string {
	var names []string

	for _, match := range interpolationRegex.FindAllSubmatch([]byte(value), -1) {
		names = append(names, string(match[1]))
	}

	return names
}

type VarsTracker struct {
	vars Variables

	expectAllFound bool
	expectAllUsed  bool

	missing map[string]struct{}
	visited map[string]visitedVar
}

type visitedVar struct {
	value interface{}
	found bool
	err   error
}

func NewVarsTracker(vars Variables, expectAllFound, expectAllUsed bool) VarsTracker {
	return VarsTracker{
		vars:           vars,
		expectAllFound: expectAllFound,
		expectAllUsed:  expectAllUsed,

		missing: map[string]struct{}{},
		visited: map[string]visitedVar{},
	}
}

// Gets the value of a field in a var
func (t VarsTracker) Get(varRef FieldReference) (interface{}, bool, error) {
	identifier := varRef.Reference.String()
	if _, ok := t.visited[identifier]; !ok {
		value, found, err := t.vars.Get(varRef.Reference)
		t.visited[identifier] = visitedVar{value, found, err}
	}

	visited := t.visited[identifier]
	if !visited.found || visited.err != nil {
		t.missing[varRef.String()] = struct{}{}
		return nil, visited.found, visited.err
	}

	val, err := Traverse(visited.value, varRef)
	if err != nil {
		t.missing[varRef.String()] = struct{}{}
		return nil, false, err
	}

	return val, true, err
}

func (t VarsTracker) Error() error {
	missingErr := t.MissingError()
	extraErr := t.ExtraError()

	if missingErr != nil && extraErr != nil {
		return multierror.Append(missingErr, extraErr)
	} else if missingErr != nil {
		return missingErr
	} else if extraErr != nil {
		return extraErr
	}

	return extraErr
}

func (t VarsTracker) MissingError() error {
	if !t.expectAllFound || len(t.missing) == 0 {
		return nil
	}

	return UndefinedVarsError{Vars: names(t.missing)}
}

func (t VarsTracker) ExtraError() error {
	if !t.expectAllUsed {
		return nil
	}

	allRefs, err := t.vars.List()
	if err != nil {
		return err
	}

	unusedNames := map[string]struct{}{}

	for _, ref := range allRefs {
		id := ref.String()
		if _, found := t.visited[id]; !found {
			unusedNames[id] = struct{}{}
		}
	}

	if len(unusedNames) == 0 {
		return nil
	}

	return UnusedVarsError{Vars: names(unusedNames)}
}

func names(mapWithNames map[string]struct{}) []string {
	var names []string
	for name, _ := range mapWithNames {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}
