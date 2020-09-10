package neivyfunctions

import (
	"fmt"
	"net/http"
)

func GetImages(w http.ResponseWriter, r *http.Request) {
	db := newDatabase()

	if err := db.Open; err != nil {
		fmt.Errorf("error initializing app: %v", err)
	}

	fmt.Fprint(w, "test")
}
