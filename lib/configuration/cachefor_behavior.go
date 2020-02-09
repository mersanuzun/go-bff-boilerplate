package configuration

import "time"

func (t *stringParamBuilder) CacheFor(duration time.Duration) *stringParamBuilder {
	var (
		current     = t.f
		cachedValue = ""
		cachedErr   error
	)

	wrapper := func() (string, error) {
		if isCacheExpired(t.fetchedTime, duration) {
			v, err := current()
			cachedValue, cachedErr = v, err

			return v, err
		}

		return cachedValue, cachedErr
	}

	t.f = wrapper

	return t
}

func (t *intParamBuilder) CacheFor(duration time.Duration) *intParamBuilder {
	var (
		current     = t.f
		cachedValue = 0
		cachedErr   error
	)

	wrapper := func() (int, error) {
		if isCacheExpired(t.fetchedTime, duration) {
			v, err := current()
			cachedValue, cachedErr = v, err

			return v, err
		}

		return cachedValue, cachedErr
	}

	t.f = wrapper

	return t
}

func (t *boolParamBuilder) CacheFor(duration time.Duration) *boolParamBuilder {
	var (
		current     = t.f
		cachedValue = false
		cachedErr   error
	)

	wrapper := func() (bool, error) {
		if isCacheExpired(t.fetchedTime, duration) {
			v, err := current()
			cachedValue, cachedErr = v, err

			return v, err
		}

		return cachedValue, cachedErr
	}

	t.f = wrapper

	return t
}

func (t *mapParamBuilder) CacheFor(duration time.Duration) *mapParamBuilder {
	var (
		current     = t.f
		cachedValue map[string]string
		cachedErr   error
	)

	wrapper := func() (map[string]string, error) {
		if isCacheExpired(t.fetchedTime, duration) {
			v, err := current()
			cachedValue, cachedErr = v, err

			return v, err
		}

		return cachedValue, cachedErr
	}

	t.f = wrapper

	return t
}

func isCacheExpired(fetchedTime time.Time, expirationDuration time.Duration) bool {
	return time.Now().Sub(fetchedTime) > expirationDuration
}
