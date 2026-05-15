package services

var Services = map[string]map[string]int {
	"database": {
		"status": 0,
		"restartCount": 0,
	},
	"nginx": {
		"status": 0,
		"restartCount": 0,
	},
	"frontend": {
		"status": 0,
		"restartCount": 0,
	},
	"backend": {
		"status": 0,
		"restartCount": 0,
	},
}