package todo

import (
	"fmt"
	"net/http"
)

func (h *Handler) HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello its my first project")
}
