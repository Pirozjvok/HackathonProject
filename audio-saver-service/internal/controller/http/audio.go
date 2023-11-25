package http

import (
	"audio-saver-service/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type audioController struct {
	uc service.Audio
}

func newAudioController(uc service.Audio) *audioController {
	return &audioController{uc: uc}
}

func newAudioRoutes(uc service.Audio, r *mux.Router) {
	audioCtrl := newAudioController(uc)

	r.HandleFunc("/audio/{id:[0-9]+}", audioCtrl.audio).Methods(http.MethodGet)
}

func (a *audioController) audio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	audio, err := a.uc.GetAudio(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(&audio)

	marshal, err := json.Marshal(audio)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(marshal)
	if err != nil {
		log.Println(err)
		return
	}
}
