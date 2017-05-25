package webu

import (
	"encoding/json"
	"net/http"
)

//HttpJson answer json
func HttpJson(w http.ResponseWriter, obj interface{}) error {

	return json.NewEncoder(w).Encode(obj)
}
