package webu

import (
	"encoding/json"
	"net/http"
)

//HttpJson answer json
func HttpJSON(w http.ResponseWriter, obj interface{}) error {
	return json.NewEncoder(w).Encode(obj)
}

func ReadJSON(r *http.Request, obj interface{}) error {

	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(obj)
}
