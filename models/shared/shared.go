package shared

type AppStatus struct {
	Status string `json:"status"`
}

type AppInformation struct {
	Env string `json:"env"`
	AppName string `json:"appName"`
	Version string `json:"version"`
}

type ErrorResponse struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}