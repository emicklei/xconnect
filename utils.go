package xconnect

type finder interface {
	find(keys []string) (interface{}, bool)
}

func findInMap(path []string, tree map[string]interface{}) (interface{}, bool) {
	if len(tree) == 0 {
		return nil, false
	}
	if len(path) == 0 {
		return nil, false
	}
	if len(path) == 1 {
		f, ok := tree[path[0]]
		if !ok {
			return nil, false
		}
		return f, true
	}
	// > 1
	f, ok := tree[path[0]]
	if !ok {
		return nil, false
	}
	mi, ok := f.(map[interface{}]interface{})
	if !ok {
		return nil, false
	}
	// do the copy
	m := map[string]interface{}{}
	for k, v := range mi {
		sk, ok := k.(string)
		if ok {
			m[sk] = v
		}
	}
	return findInMap(path[1:], m)
}

func copy(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for key, value := range src {
		dst[key] = value
	}
	return dst
}
