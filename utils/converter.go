package utils

// Here you can give your routings keys instead women
const parsingKey = "women"

// ConvertYamlToJSON will return a json interface
func ConvertYamlToJSON(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			if k == parsingKey {
				m2[k.(string)] = ConvertYamlToJSON(v)
			}
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = ConvertYamlToJSON(v)
		}
	}
	return i
}
