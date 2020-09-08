package neivyfunctions

import (
	"net/http"
	"fmt"
)

func GetImages(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "test")
}