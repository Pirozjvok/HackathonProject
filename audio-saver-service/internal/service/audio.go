package service

import (
	"audio-saver-service/internal/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type audioService struct {
	url string
}

func (a *audioService) GetAudio(id string) (*os.File, error) {
	uri := a.url + id
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	bell := &model.BellInfo{}

	err = json.Unmarshal(body, bell)
	if err != nil {
		return nil, err
	}
	fmt.Println(bell)

	req, err = http.NewRequest(http.MethodGet, bell.ContactAudio, nil)
	if err != nil {
		return nil, err
	}

	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	idi := strconv.Itoa(bell.ID)

	file, err := os.Create("assets/" + idi + ".mp3")

	r := bytes.NewReader(body)

	_, err = io.Copy(file, r)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func newAudioService(url string) *audioService {
	return &audioService{url: url}
}

var _ Audio = (*audioService)(nil)
