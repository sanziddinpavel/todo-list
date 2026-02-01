package user

import (
	"Todo-list/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var req LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "invalid request data")
		return

	}

	usr, err := h.userRepo.Find(req.Email, req.Password)
	if usr == nil {
		util.SendError(w, http.StatusBadRequest, "unauthorized")
		return
	}

	accessToken, err := util.CreateJwt(h.cnf.JwtSecretKey, util.Payload{
		ID:        usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
	})
	if err != nil {
		http.Error(w, "failed to create jwt", http.StatusInternalServerError)
		return
	}

	util.SendData(w, accessToken, http.StatusOK)

}
