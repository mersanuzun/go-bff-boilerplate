package providers

type Provider interface {
	GetString(key string) (string, error)
	GetInt(key string) (int, error)
	GetBool(key string) (bool, error)
	GetMap(key string) (map[string]string, error)
}
