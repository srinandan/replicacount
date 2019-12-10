package apis

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	count "github.com/srinandan/replicacount/count"
	types "github.com/srinandan/replicacount/types"
)

//ReplicaCountHandler handles GET /tokens
func ReplicaCountHandler(w http.ResponseWriter, r *http.Request) {

	errorMessage := types.ErrorMessage{}
	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage.Message = "id is mandatory"
		errorMessage.StatusCode = http.StatusInternalServerError
		if err := json.NewEncoder(w).Encode(errorMessage); err != nil {
			types.Error.Println(err)
		}
		return
	}

	count, err := count.GetReplicaCount(id)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage.Message = err.Error()
		errorMessage.StatusCode = http.StatusInternalServerError
		if err := json.NewEncoder(w).Encode(errorMessage); err != nil {
			types.Error.Println(err)
		}
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(count); err != nil {
		types.Error.Println(err)
	}
}
