package server

import "testServerStats/pkg/statistics"

func InitValidators() {
	statistics.AddCustomRules()
}
