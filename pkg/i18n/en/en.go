package en

var TranslateToml = map[string]string{
	"hello_world":                            "hello {{.name}}",
	"unnamed":                                "unnamed",
	"time_y":                                 "{{.v}}years",
	"time_m":                                 "{{.v}}months",
	"time_d":                                 "{{.v}}days",
	"time_h":                                 "{{.v}}hours",
	"time_mm":                                "{{.v}}minutes",
	"time_s":                                 "{{.v}}seconds",
	"zerotrust_connector_device_offline":     "offline",
	"zerotrust_connector_device_online":      "online",
	"zerotrust_deviceDiagnose_openresty":     "Checking web service components",
	"zerotrust_deviceDiagnose_openrestyPort": "Check web service port",
	"zerotrust_deviceDiagnose_webTerminal":   "Check protocol application component 1",
	"zerotrust_deviceDiagnose_webTerminalPort": "Check protocol application component 1 Port",
	"zerotrust_deviceDiagnose_guacamole":       "Check protocol application component 2",
	"zerotrust_deviceDiagnose_guacamolePort":   "Check protocol application component 2 Port",
}
