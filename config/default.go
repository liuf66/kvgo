package config

var DefaultServerConfig *ServerConfig

func init() {
	DefaultServerConfig = &ServerConfig{
		Host: "127.0.0.1",
		Port: 7547,
	}
}
