package statistics

import (
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
	"time"
)

func AddCustomRules () {
	govalidator.AddCustomRule("sex_validator", func(field string, rule string, message string, value interface{}) error {
		val := value.(string)
		if val != "M" && val != "W" {
			return fmt.Errorf("The %s field must be M or W", field)
		}
		return nil
	})

	govalidator.AddCustomRule("action_validator", func(field string, rule string, message string, value interface{}) error {
		val := value.(string)
		if val != "login" && val != "like" && val != "comments" && val != "logout" {
			return fmt.Errorf("The %s field must be login or like or comment or logout", field)
		}
		return nil
	})

	govalidator.AddCustomRule("short_rffc339", func(field string, rule string, message string, value interface{}) error {
		val := value.(string)
		_, err := time.Parse("2006-01-02T15:04:05", val)
		if err != nil {
			return fmt.Errorf("The %s field should be date in format RFC339 without tz, error: %s", field, err.Error())
		}
		return nil
	})

	govalidator.AddCustomRule("id_validator", func(field string, rule string, message string, value interface{}) error {
		val := value.(int)
		if val <= 0 {
			return fmt.Errorf("The %s field should greater than 1", field)
		}
		return nil
	})
}

func validatorErrorHandler(e url.Values, w http.ResponseWriter) bool {
	if len(e) != 0 {
		err := map[string]interface{}{"error": e}
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return true
	}
	return false
}

func validateGetStats(r *http.Request) url.Values {
	rules := govalidator.MapData{
		"date1":  []string{"date:yyyy-dd-mm"},
		"date2":  []string{"date:yyyy-dd-mm"},
		"action": []string{"action_validator"},
		"limit":  []string{"numeric_between:1,"},
	}
	opts := govalidator.Options{
		Request:         r,
		Rules:           rules,
		RequiredDefault: true,
	}
	v := govalidator.New(opts)
	e := v.Validate()
	return e
}

func validateCreateStat(r *http.Request) (*Stat, url.Values) {
	var stats Stat
	rules := govalidator.MapData{
		"user":   []string{"id_validator"},
		"ts":     []string{"short_rffc339"},
		"action": []string{"action_validator"},
	}
	opts := govalidator.Options{
		Request:         r,
		Data:            &stats,
		Rules:           rules,
		RequiredDefault: true,
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	return &stats, e
}

func validateCreateUser(r *http.Request) (*User, url.Values) {
	var user User
	rules := govalidator.MapData{
		"id":  []string{"id_validator"},
		"age": []string{"between:1,100"},
		"sex": []string{"sex_validator"},
	}
	opts := govalidator.Options{
		Request:         r,
		Data:            &user,
		Rules:           rules,
		RequiredDefault: true,
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	return &user, e
}
