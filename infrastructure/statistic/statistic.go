package statistic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/MikelSot/interseguro-challenge-qr/model"
)

const (
	_headerContentType      = "Content-Type"
	_headerContentTypeValue = "application/json"
)

type Statistic struct {
	config model.ConfigStatistic
}

func New(config model.ConfigStatistic) Statistic {
	return Statistic{
		config: config,
	}
}

func (s Statistic) GetStatistic(QR model.QR) (model.StatisticResponse, error) {
	payload, err := json.Marshal(QR)
	if err != nil {
		return model.StatisticResponse{}, fmt.Errorf("statistic.GetStatistic(), json.Marshal: %w", err)
	}

	body, err := s.doRequest(payload)
	if err != nil {
		return model.StatisticResponse{}, fmt.Errorf("qr.statistic: %w", err)
	}

	var ms model.StatisticResponse
	if err := json.Unmarshal(body, &ms); err != nil {
		return model.StatisticResponse{}, fmt.Errorf("qr.statistic.json.Unmarshal(): %w", err)
	}

	return ms, nil
}

func (s Statistic) doRequest(payload []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, s.config.UrlStatistic, bytes.NewBuffer(payload))
	if err != nil {
		log.Warn("http.NewRequest: ", err.Error())

		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Add(_headerContentType, _headerContentTypeValue)

	client := http.Client{
		Timeout: time.Second * 10,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Warn("client.Do()", err.Error())

		return nil, fmt.Errorf("client.Do: %w", err)
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Warn("io.ReadAll", err.Error())

		return nil, fmt.Errorf("reading Body: %v", err)
	}

	if res.StatusCode == http.StatusNotFound || res.StatusCode >= http.StatusInternalServerError {
		customErr := model.NewError()

		customErr.SetCode(model.Failure)
		customErr.SetStatusHTTP(http.StatusInternalServerError)
		customErr.SetAPIMessage("Error al obtener la estadística")

		return nil, customErr
	}

	if res.StatusCode != http.StatusOK {
		customErr := model.NewError()

		customErr.SetStatusHTTP(res.StatusCode)
		customErr.SetAPIMessage("Error al obtener la estadística")

		return nil, customErr
	}

	return bodyBytes, nil
}
