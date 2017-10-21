package webu

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/////////////////////////////////////////
// IO Helpers
////////////////

//WriteJSON Writes json to writer
//
//Example:
//  webu.WriteJSON(w, data)
//
func WriteJSON(w http.ResponseWriter, obj interface{}) error {
	return json.NewEncoder(w).Encode(obj)
}

//ReadJSON reads json into obj
//
//Example:
//  webu.ReadJSON(r, obj)
func ReadJSON(r *http.Request, obj interface{}) error {

	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(obj)
}

//WriteStatus Writes the httpStatus with StatusText as body
func WriteStatus(w http.ResponseWriter, code int, extras ...interface{}) {
	w.WriteHeader(code)
	extra := fmt.Sprint(extras...)
	WriteJSON(w, http.StatusText(code)+"\r\n"+extra)
}
