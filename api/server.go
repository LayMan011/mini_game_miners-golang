package api

import (
	"errors"
	"fmt"
	"myproj/data/info"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	handlers *HTTPHandlers
}

func NewHTTPServer(handlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		handlers: handlers,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/miners/info").Methods("GET").HandlerFunc(s.handlers.HandlerMinersInfo);
	router.Path("/miners/{class}").Methods("POST").HandlerFunc(s.handlers.HandlerMinerAdd);
	router.Path("/miners").Methods("GET").Queries("now", "true").HandlerFunc(s.handlers.HandlerMinersNow);
	router.Path("/miners/{class}").Methods("GET").HandlerFunc(s.handlers.HandlerMinersNowClass);
	router.Path("/equipments/info").Methods("GET").HandlerFunc(s.handlers.HandlerPriceEquipmetns);
	router.Path("/equipments/{equipment}").Methods("PATCH").HandlerFunc(s.handlers.HandlerBuyEquipments);
	router.Path("/equipments").Methods("GET").HandlerFunc(s.handlers.HandlerCompleteEquipments);
	router.Path("/info").Methods("GET").HandlerFunc(s.handlers.HandlerInfoCompany);
	router.Path("/end").Methods("GET").HandlerFunc(s.handlers.HandlerEndGame);

    go func() {
        ch := info.StartPassiveIncome()
        info.ReadFromCh(ch)
    }()
    
    fmt.Println("Server starting on :9091")

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}