package busserver

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"log"
	"monbus/pkg/storage"
	"net/http"
)

type BusServer struct {
	storage *storage.Storage
}

func New(s *storage.Storage) *BusServer {
	return &BusServer{
		storage: s,
	}
}

func (b *BusServer) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/bustime/{id}", b.busTimeHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func (b *BusServer) busTimeHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("start busTimeHandler with %v", r)
	if r.Method != "POST" {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// get stationId
	id := mux.Vars(r)["id"]

	// get result from db
	res, err := b.storage.GetBusTime(id)
	if err != nil {
		http.Error(w, "get bustime from db error", http.StatusInternalServerError)
		return
	}

	resp := &SkillResponse{
		Version: "2.0",
		Template: Template{
			Outputs: []SkillOutput{
				{
					SimpleText: SimpleText{
						Text: res,
					},
				},
			},
		},
	}
	glog.Infof("result is %s", resp)

	j, _ := json.Marshal(resp)
	w.Write(j)
}
