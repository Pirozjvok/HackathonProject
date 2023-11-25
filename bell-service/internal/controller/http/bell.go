package http

import (
	"bell-service/internal/model"
	"bell-service/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type bellController struct {
	apiUC service.APIRequester
	rudUC service.RUD
}

func newBellController(apiUC service.APIRequester) *bellController {
	return &bellController{apiUC: apiUC}
}

func newBellRoutes(uc *service.Services, r *mux.Router) {
	bellCtrl := newBellController(uc.APIRequester)

	r.HandleFunc("/dima", bellCtrl.requester).Methods(http.MethodGet)
	r.HandleFunc("/get/{id:[0-9]+}", bellCtrl.getByID).Methods(http.MethodGet)
	r.HandleFunc("/get/all", bellCtrl.getAll).Methods(http.MethodGet)
	r.HandleFunc("/update/{id:[0-9]+}", bellCtrl.update).Methods(http.MethodPost)
	r.HandleFunc("/remove/{id:[0-9]+}", bellCtrl.remove).Methods(http.MethodGet)
}

func (b *bellController) requester(w http.ResponseWriter, r *http.Request) {

	err := b.apiUC.Request(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody := map[string]interface{}{
		"ok": "ok",
	}

	response, err := json.Marshal(responseBody)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(response)
	if err != nil {
		log.Println(err)
		return
	}
}
func (b *bellController) getByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bell, err := b.rudUC.GetByID(r.Context(), vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ans, err := json.Marshal(bell)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(ans)
	if err != nil {
		log.Println(err)
		return
	}
}
func (b *bellController) getAll(w http.ResponseWriter, r *http.Request) {
	all, err := b.rudUC.GetAll(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ans, err := json.Marshal(all)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(ans)
	if err != nil {
		log.Println(err)
		return
	}
}
func (b *bellController) update(w http.ResponseWriter, r *http.Request) {
	var body model.BellInfo

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = b.rudUC.Update(r.Context(), &body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("ok"))
	if err != nil {
		log.Println(err)
		return
	}
}
func (b *bellController) remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	err := b.rudUC.Delete(r.Context(), vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("ok"))
	if err != nil {
		log.Println(err)
		return
	}
}
