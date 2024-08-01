package qr

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"gonum.org/v1/gonum/mat"

	"github.com/MikelSot/interseguro-challenge-qr/model"
)

type QR struct {
}

func New() *QR {
	return &QR{}
}

func (qr QR) FactorizeQR(matrix [][]float64) (model.QR, error) {
	rows := len(matrix)
	if rows == 0 {
		log.Warn("¡Uy! la matriz está vacía")

		customErr := model.NewError()
		customErr.SetCode(model.InvalidMatrix)
		customErr.SetAPIMessage("¡Uy! la matriz está vacía")
		customErr.SetStatusHTTP(http.StatusBadRequest)

		return model.QR{}, customErr
	}

	cols := len(matrix[0])
	data := make([]float64, 0, rows*cols)

	// fill in the data for each row of the matrix
	for _, row := range matrix {
		data = append(data, row...)
	}

	math := mat.NewDense(rows, cols, data)

	var mathQR mat.QR
	mathQR.Factorize(math)

	responseRQ, err := qr.getRQ(mathQR, rows, cols)
	if err != nil {
		return model.QR{}, err
	}

	return responseRQ, nil
}

func (qr QR) getRQ(mathQR mat.QR, rows, cols int) (model.QR, error) {
	var q mat.Dense
	var r mat.Dense
	mathQR.QTo(&q)
	mathQR.RTo(&r)

	qData := make([][]float64, rows)
	rData := make([][]float64, cols)

	for i := 0; i < rows; i++ {
		qData[i] = q.RawRowView(i)
	}
	for i := 0; i < cols; i++ {
		rData[i] = r.RawRowView(i)
	}

	if len(qData) == 0 || len(rData) == 0 {
		log.Warn("qr: ¡Uy! la matriz está vacía")

		customErr := model.NewError()
		customErr.SetCode(model.InvalidMatrix)
		customErr.SetAPIMessage("¡Uy! la matriz Q o R está vacía")
		customErr.SetStatusHTTP(http.StatusBadRequest)

		return model.QR{}, customErr
	}

	responseRQ := model.QR{
		Q: qData,
		R: rData,
	}

	return responseRQ, nil
}
