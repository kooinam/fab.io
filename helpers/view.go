package helpers

// IncludeRootInJSON used prepend root in JSON
func IncludeRootInJSON(json interface{}, includeRoot bool, root string) interface{} {
	if includeRoot == false {
		return json
	}

	jsonWithRoot := make(map[string]interface{})
	jsonWithRoot[root] = json

	return jsonWithRoot
}

// BuildJSON used to build json with parameters like root key
func BuildJSON(json interface{}, includeRoot bool, root string) interface{} {
	if includeRoot == false {
		return json
	}

	jsonWithRoot := make(map[string]interface{})
	jsonWithRoot[root] = json

	return jsonWithRoot
}
