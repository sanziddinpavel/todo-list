package todo

import (
	"Todo-list/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {

	reqQuery := r.URL.Query()

	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	todoList, err := h.svc.List(page, limit)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	cnt, err := h.svc.Count()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	paginatedData := util.PaginatedData{
		Data: todoList,
		Pagination: util.Pagination{
			Page:       page,
			Limit:      limit,
			TotalItems: cnt,
			TotalPages: cnt / limit,
		},
	}

	util.SendPage(w, paginatedData, page, limit, cnt)

	// encoder := json.NewEncoder(w)
	// encoder.Encode(Todolist)

}
