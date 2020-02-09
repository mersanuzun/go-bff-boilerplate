package providers

import (
	"errors"
)

type mapConfiguration struct {
	configMap map[string]interface{}
}

func Map(configMap map[string]interface{}) Provider {
	return &mapConfiguration{configMap}
}

func (c *mapConfiguration) GetString(key string) (string, error) {
	rawValue, ok := c.configMap[key]

	if !ok {
		return "", errors.New("Key not found in the Map. Key: " + key)
	}

	value, ok := rawValue.(string)

	if !ok {
		return "", errors.New("Value in Map cannot be converted to a string. Key: " + key)
	}

	return value, nil
}

func (c *mapConfiguration) GetInt(key string) (int, error) {
	rawValue, ok := c.configMap[key]

	if !ok {
		return 0, errors.New("Key not found in the Map. Key: " + key)
	}

	value, ok := rawValue.(int)

	if !ok {
		return 0, errors.New("Value in Map cannot be converted to an int. Key: " + key)
	}

	return value, nil
}

func (c *mapConfiguration) GetBool(key string) (bool, error) {
	rawValue, ok := c.configMap[key]

	if !ok {
		return false, errors.New("Key not found in the Map. Key: " + key)
	}

	value, ok := rawValue.(bool)

	if !ok {
		return false, errors.New("Value in Map cannot be converted to a bool. Key: " + key)
	}

	return value, nil
}

func (c *mapConfiguration) GetMap(key string) (map[string]string, error) {
	rawValue, ok := c.configMap[key]

	if !ok {
		return nil, errors.New("Key not found in the Map. Key: " + key)
	}

	value, ok := rawValue.(map[string]string)

	if !ok {
		return nil, errors.New("Value in Map cannot be converted to an int. Key: " + key)
	}

	return value, nil
}
