package configuration

func (t *stringParamBuilder) DefaultTo(defaultValue string) *stringParamBuilder {
	current := t.f

	wrapper := func() (string, error) {
		v, err := current()

		if err != nil {
			return defaultValue, nil
		}

		return v, nil
	}

	t.f = wrapper

	return t
}

func (t *intParamBuilder) DefaultTo(defaultValue int) *intParamBuilder {
	current := t.f

	wrapper := func() (int, error) {
		v, err := current()
		if err != nil {
			return defaultValue, nil
		}
		return v, nil
	}

	t.f = wrapper

	return t
}

func (t *boolParamBuilder) DefaultTo(defaultValue bool) *boolParamBuilder {
	current := t.f

	wrapper := func() (bool, error) {
		v, err := current()
		if err != nil {
			return defaultValue, nil
		}
		return v, nil
	}

	t.f = wrapper

	return t
}

func (t *mapParamBuilder) DefaultTo(defaultValue map[string]string) *mapParamBuilder {
	current := t.f

	wrapper := func() (map[string]string, error) {
		v, err := current()
		if err != nil {
			return defaultValue, nil
		}
		return v, nil
	}

	t.f = wrapper

	return t
}
