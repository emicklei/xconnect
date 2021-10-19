package xconnect

import "log"

func findInMap(path []string, tree map[string]interface{}) (interface{}, bool) {
	if len(tree) == 0 {
		log.Println("warn: xconnect, empty extra fields")
		return nil, false
	}
	if len(path) == 0 {
		log.Println("warn: xconnect, empty key")
		return nil, false
	}
	if len(path) == 1 {
		f, ok := tree[path[0]]
		if !ok {
			log.Println("warn: xconnect, no such key", path[0])
			return nil, false
		}
		return f, true
	}
	// > 1
	f, ok := tree[path[0]]
	if !ok {
		log.Println("warn: xconnect, no such key", path[0])
		return nil, false
	}
	mi, ok := f.(map[interface{}]interface{})
	if !ok {
		log.Printf("warn: xconnect, value is not a map, but a %T for key %s\n", f, path[0])
		return nil, false
	}
	// do the copy
	m := map[string]interface{}{}
	for k, v := range mi {
		sk, ok := k.(string)
		if !ok {
			log.Printf("warn: xconnect, key %s is not a string but %T\n", k, k)
		} else {
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
