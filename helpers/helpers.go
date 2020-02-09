package helpers

import (
	"encoding/json"
	"github.com/mersanuzun/go-bff-boilerplate/models/shared"
	"github.com/parnurzeal/gorequest"
)

func HandleClientResponse(sa *gorequest.SuperAgent, errorStatusStart int) (gorequest.Response, string) {
	resp, body, errs := sa.End()

	if len(errs) > 0 {
		panic(errs[0])
	}

	if resp.StatusCode > errorStatusStart {
		bytes, _ := json.Marshal(&shared.ErrorResponse{Status: resp.StatusCode, Body: body})

		panic(string(bytes))
	}

	return resp, body
}
