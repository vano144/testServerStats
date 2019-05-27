package statistics

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func errorWriter(e error, w http.ResponseWriter, status int) {
	err := map[string]interface{}{"error": e.Error()}
	w.Header().Set("Content-type", "applciation/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(err)
}

func CreateStat(w http.ResponseWriter, r *http.Request) {
	stat, verr := validateCreateStat(r)
	if validatorErrorHandler(verr, w) {
		return
	}
	err := createStatController(stat)
	if err != nil {
		switch eType := err.(type) {
		case *EntityDoesNotAlreadyExists:
			errorWriter(eType, w, http.StatusBadRequest)
			return
		default:
			log.WithError(eType).Error("createStat problem")
			errorWriter(eType, w, http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-type", "applciation/json")
	w.WriteHeader(http.StatusCreated)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user, e := validateCreateUser(r)
	if validatorErrorHandler(e, w) {
		return
	}
	err := createUserController(user)
	if err != nil {
		switch eType := err.(type) {
		case *EntityAlreadyExists:
			errorWriter(err, w, http.StatusBadRequest)
			return
		default:
			log.WithError(eType).Error("createUser problem")
			errorWriter(eType, w, http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-type", "applciation/json")
	w.WriteHeader(http.StatusCreated)

}

func GetAccumulateStats(w http.ResponseWriter, r *http.Request) {
	e := validateGetStats(r)
	if validatorErrorHandler(e, w) {
		return
	}
	queryValues := r.URL.Query()
	date1 := queryValues.Get("date1")
	date2 := queryValues.Get("date2")
	action := queryValues.Get("action")
	limit, _ := strconv.Atoi(queryValues.Get("limit"))
	result, err := getAccumulateStatController(date1, date2, action, limit)
	z, err := json.Marshal(&result)
	if err != nil {
		log.WithError(err).Error("GetAccumulateStats problem")
		errorWriter(err, w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "applciation/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(z)
	if err != nil {
		log.WithError(err).Error("Failed to write to output")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetStatsPerDay(w http.ResponseWriter, r *http.Request) {
	e := validateGetStats(r)
	if validatorErrorHandler(e, w) {
		return
	}
	queryValues := r.URL.Query()
	date1 := queryValues.Get("date1")
	date2 := queryValues.Get("date2")
	action := queryValues.Get("action")
	limit, _ := strconv.Atoi(queryValues.Get("limit"))
	result, err := getStatPerDayController(date1, date2, action, limit)
	z, err := json.Marshal(&result)
	if err != nil {
		log.WithError(err).Error("GetStatsPerDay problem")
		errorWriter(err, w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "applciation/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(z)
	if err != nil {
		log.WithError(err).Error("Failed to write to output")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
