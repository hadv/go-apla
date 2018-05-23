package api

import (
	"net/http"

	"github.com/AplaProject/go-apla/packages/consts"

	log "github.com/sirupsen/logrus"
)

func getVersion(w http.ResponseWriter, r *http.Request, data *apiData, logger *log.Entry) (err error) {
	data.result = consts.VERSION
	return nil
}
