package config

type Config struct {
	LinPhone LinPhone `json:"linphone"`
}

type LinPhone struct {
	Username string `json:"username"`
	Domain   string `json:"domain"`
	Address  string `json:"address"`
	ApiKey   string `json:"apikey"`
	Proxy    string `json:"proxy"`
}
