package configuration

func (t *stringParamBuilder) CacheOnce() *stringParamBuilder {
	current := t.f
	var (
		isCached    = false
		cachedValue = ""
		cachedErr   error
	)

	wrapper := func() (string, error) {
		if !isCached {
			v, err := current()
			isCached, cachedValue, cachedErr = true, v, err

			return v, err
		}

		return cachedValue, cachedErr
	}

	t.f = wrapper

	return t
}

func (t *intParamBuilder) CacheOnce() *intParamBuilder {
	current := t.f
	var (
		isCached    = false
		cachedValue = 0
		cachedErr   error
	)

	wrapper := func() (int, error) {
		if !isCached {
			v, err := current()
			isCached, cachedValue, cachedErr = true, v, err
			return v, err
		}

		return cachedValue, cachedErr
	}

	t.f = wrapper

	return t
}

func (t *boolParamBuilder) CacheOnce() *boolParamBuilder {
	current := t.f
	var (
		isCached    = false
		cachedValue = false
		cachedErr   error
	)

	wrapper := func() (bool, error) {
		if !isCached {
			v, err := current()
			isCached, cachedValue, cachedErr = true, v, err
			return v, err
		}

		return cachedValue, cachedErr
	}

	t.f = wrapper

	return t
}

func (t *mapParamBuilder) CacheOnce() *mapParamBuilder {
	current := t.f
	var (
		isCached    = false
		cachedValue map[string]string
		cachedErr   error
	)

	wrapper := func() (map[string]string, error) {
		if !isCached {
			v, err := current()
			isCached, cachedValue, cachedErr = true, v, err
			return v, err
		}

		return cachedValue, cachedErr
	}

	t.f = wrapper

	return t
}
