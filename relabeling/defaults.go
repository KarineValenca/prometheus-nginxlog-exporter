package relabeling

import (
	"github.com/martin-helmich/prometheus-nginxlog-exporter/config"
)

// DefaultRelabelings are hardcoded relabeling configs that are always there
// and do not need to be explicitly configured
var DefaultRelabelings = []*Relabeling{
	{
		config.RelabelConfig{
			TargetLabel: "method",
			SourceValue: "request",
			Split:       1,

			WhitelistExists: true,
			WhitelistMap: map[string]interface{}{
				"GET":     "GET",
				"HEAD":    "HEAD",
				"POST":    "POST",
				"PUT":     "PUT",
				"DELETE":  "DELETE",
				"CONNECT": "CONNECT",
				"OPTIONS": "OPTIONS",
				"TRACE":   "TRACE",
				"PATCH":   "PATCH",
			},
		},
	},
	{
		config.RelabelConfig{
			TargetLabel: "status",
			SourceValue: "status",
		},
	},
	{
		config.RelabelConfig{
			TargetLabel: "addr",
			SourceValue: "request",
			Split:       2,
		},
	},
	{
		config.RelabelConfig{
			TargetLabel: "type",
			SourceValue: "request",
			Split:       3,
		},
	},
	{
		config.RelabelConfig{
			TargetLabel:     "isError",
			SourceValue:     "status",
			WhitelistExists: true,
			WhitelistMap: map[string]interface{}{
				"400": "true",
				"401": "true",
				"403": "true",
				"404": "true",
				"405": "true",
				"422": "true",
				"500": "true",
				"501": "true",
				"502": "true",
				"503": "true",
				"504": "true",
				"200": "false",
				"201": "false",
				"202": "false",
				"204": "false",
				"206": "false",
				"300": "false",
				"301": "false",
				"302": "false",
				"303": "false",
				"304": "false",
				"307": "false",
			},
		},
	},
}
