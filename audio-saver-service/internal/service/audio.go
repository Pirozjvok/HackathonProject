package service

import (
	"audio-saver-service/internal/model"
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"strconv"
)

type audioService struct {
	url string
}

func (a *audioService) GetAudio(id string) ([]byte, error) {
	uri := a.url + id
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, "GetAudio: fail to create in constructor request1")
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "GetAudio: fail to create request1")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "GetAudio: fail to read response body1")
	}
	bell := &model.BellInfo{}

	err = json.Unmarshal(body, bell)
	if err != nil {
		return nil, errors.Wrap(err, "GetAudio: fail to unmarshalling JSON to struct")
	}

	req, err = http.NewRequest(http.MethodGet, bell.ContactAudio, nil)
	if err != nil {
		return nil, errors.Wrap(err, "GetAudio: fail to create in constructor request2")
	}

	res, err = client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "GetAudio: fail to create request2")
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "GetAudio: fail to read response body2")
	}

	idi := strconv.Itoa(bell.ID)

	file, err := os.Create("assets/" + idi + ".mp3")
	if err != nil {
		return nil, errors.Wrap(err, "GetAudio: fail to create new file in assets")
	}
	defer file.Close()

	r := bytes.NewReader(body)

	_, err = io.Copy(file, r)
	if err != nil {
		return nil, errors.Wrap(err, "GetAudio: fail to copy body in file")
	}

	data := make([]byte, 64)

	_, err = file.Read(data)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return data, nil
		}
		return nil, errors.Wrap(err, "GetAudio: fail to read file")
	}

	return nil, nil
}

func newAudioService(url string) *audioService {
	return &audioService{url: url}
}

var _ Audio = (*audioService)(nil)
