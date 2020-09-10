package neivyfunctions

import (
	"fmt"
	"net/http"
)

func GetImages(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "test")
}
