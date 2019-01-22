package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/MartinToruan/tax_calculator/logic"
	"github.com/MartinToruan/tax_calculator/persistence"
	"github.com/MartinToruan/tax_calculator/svc"
	"github.com/google/uuid"
)

var (
	persistenceClient svc.Persistence
	logicClient       svc.Logic
)

// Response
var responseCreated = map[string]string{"status": "created"}
var invalidURLError = map[string]string{"error": "invalid api url"}
var systemError = map[string]string{"error": "system error. Call IT Support."}

func init() {
	// Setup Persistence
	persistenceClient = &persistence.TaxPersistence{
		DB_NAME:       "tax",
		DB_HOST:       "tax_calculator_db",
		DB_PORT:       5432,
		DB_USER:       "postgres",
		DB_PASSWORD:   "postgres",
		MAX_DB_CLIENT: 20,
	}
	persistenceClient.Init()

	// Setup Logic
	logicClient = &logic.TaxLogic{
		MAX_CONCURRENT: 20,
	}
	logicClient.Init()
}

func deinit() {
	var wg sync.WaitGroup

	fmt.Println("Shutting Down Server...")
	wg.Add(2)
	go func() {
		defer wg.Done()
		persistenceClient.DeInit()
	}()

	go func() {
		defer wg.Done()
		logicClient.DeInit()
	}()
	wg.Wait()
}

func main() {
	// Trap Signal
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		// Run The Server
		http.HandleFunc("/", handleRequest)

		// Server will run on Localhost:8000
		if err := http.ListenAndServe(":8000", nil); err != nil {
			fmt.Println("failed start server. err: ", err)
			close(stopChan)
			return
		}
	}()

	// Wait Signal to Stop the Server
	// Signal May be `ctrl+c` || `kill -9 PID` || etc
	<-stopChan

	// When Server Stopped, Run Deinit function in Logic and Persistence, to close all openned connection
	deinit()
	fmt.Println("Server Down.")
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Construct Request ID
	reqId := uuid.New().String()
	fmt.Println(reqId)

	// Routing Request
	// The Request will be processed by appropriate Logic Handler
	switch path := r.URL.Path; path {
	case "/add/tax":
		err := logicClient.HandleAddTax(r, persistenceClient)
		if err != nil {
			fmt.Println("err: ", err)

			// System Failure
			errorResp, _ := json.Marshal(systemError)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write(errorResp)
			return
		}
		// Response
		responseCreated, _ := json.Marshal(responseCreated)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write(responseCreated)
		return

	case "/get_bill":
		var data []byte
		data, err := logicClient.HandleGetBill(persistenceClient)
		if err != nil {
			fmt.Println("err: ", err)

			// System Failure
			errorResp, _ := json.Marshal(systemError)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write(errorResp)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(data)
		return
	default:
		// Invalid URL
		errorResp, _ := json.Marshal(invalidURLError)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write(errorResp)
		return
	}
}
