package runner

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/viniciusramosdefaria/zerohashchallenge/pkg/metrics"
	"github.com/viniciusramosdefaria/zerohashchallenge/pkg/phttp"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)


type Payload struct {
	Metadata map[string]string `json:"metadata,omitempty"`
}

type handler struct {

}

func NewRunnerHandler() phttp.HttpHandler {
	return handler{}
}

func (h handler) Handler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("reading the request body failed, %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		payload := &Payload{}
		err = json.Unmarshal(body, payload)
		if err != nil {
			log.Printf("decoding the request body failed, %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(payload.Metadata) > 0 {
			if payload.Metadata["rounds"] != ""{
				roundCounter, _ := strconv.Atoi(payload.Metadata["rounds"])
				rounds := GenerateRandomGameData(CardPack, TurnOrder, roundCounter)
				winner := RunGame(rounds)
				metrics.Add("winner", winner)
			}
		}
	}
}