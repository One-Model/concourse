package vars

type StaticVariables map[string]interface{}

var _ Variables = StaticVariables{}

func (v StaticVariables) Get(ref Reference) (interface{}, bool, error) {
	if ref.Source != "" {
		return nil, false, nil
	}

	val, found := v[ref.Path]
	return val, found, nil
}

func (v StaticVariables) List() ([]Reference, error) {
	var refs []Reference

	for name, _ := range v {
		refs = append(refs, Reference{Path: name})
	}

	return refs, nil
}

func (v StaticVariables) Flatten() KVPairs {
	flat := make(KVPairs, 0, len(v))
	for k, vv := range v {
		flat = append(flat, flatten(k, nil, vv)...)
	}

	return flat
}

func flatten(path string, fields []string, value interface{}) KVPairs {
	var flat KVPairs

	switch node := value.(type) {
	case map[string]interface{}:
		for k, v := range node {
			flat = append(flat, flatten(path, append(fields, k), v)...)
		}

	case map[interface{}]interface{}:
		for k, v := range node {
			if str, ok := k.(string); ok {
				flat = append(flat, flatten(path, append(fields, str), v)...)
			}
		}

	default:
		flat = KVPairs{{
			Ref:   NewFieldReferenceWithoutSource(path, fields),
			Value: value,
		}}
	}

	return flat
}

type KVPair struct {
	Ref   FieldReference
	Value interface{}
}

type KVPairs []KVPair

func (p KVPairs) Expand() StaticVariables {
	out := make(map[string]interface{}, len(p))
	for _, pair := range p {
		upsert(out, pair.Ref.Path, pair.Ref.Fields, pair.Value)
	}

	return out
}

func upsert(out map[string]interface{}, path string, fields []string, value interface{}) {
	node, ok := out[path]
	if !ok {
		out[path] = constructValue(fields, value)
		return
	}

	nodeMap, ok := node.(map[string]interface{})
	if !ok {
		out[path] = constructValue(fields, value)
		return
	}

	if len(fields) == 0 {
		out[path] = value
		return
	}

	upsert(nodeMap, fields[0], fields[1:], value)
}

func constructValue(fields []string, value interface{}) interface{} {
	if len(fields) == 0 {
		return value
	}

	return constructValue(fields[:len(fields)-1], map[string]interface{}{fields[len(fields)-1]: value})
}
