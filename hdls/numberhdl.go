package hdls

import (
	"encoding/json"
	"net/http"
	"strconv"

	"trafilea/model"
	"trafilea/srvs"

	"github.com/gorilla/mux"
)

type NumberHdl struct {
	numberSrv srvs.NumberSrv
}

func New(numberSrv srvs.NumberSrv) NumberHdl {
	return NumberHdl{numberSrv: numberSrv}
}

func (hdl *NumberHdl) PostNumber(w http.ResponseWriter, r *http.Request) {
	var num model.Number
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&num)

	hdl.numberSrv.SaveNumber(num)
	w.WriteHeader(http.StatusCreated)
}

func (hdl *NumberHdl) GetNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	num, err := hdl.numberSrv.GetNumber(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(num)
}

func (hdl *NumberHdl) GetNumbers(w http.ResponseWriter, r *http.Request) {
	nums := hdl.numberSrv.GetNumbers()
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(nums)
}
