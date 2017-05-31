package webu

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//HttpJson answer json
func WriteJSON(w http.ResponseWriter, obj interface{}) error {
	return json.NewEncoder(w).Encode(obj)
}

func ReadJSON(r *http.Request, obj interface{}) error {

	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(obj)
}

func WriteStatus(w http.ResponseWriter, code int, extras ...interface{}) {
	w.WriteHeader(code)
	extra := fmt.Sprint(extras)
	WriteJSON(w, http.StatusText(code)+"\r\n"+extra)
}
