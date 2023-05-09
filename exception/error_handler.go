package exception

import (
	"net/http"
)

func ErrorHalder(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Write([]byte("ERROR"))
}
