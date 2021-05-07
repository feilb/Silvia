package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/feilb/Silvia/api"
	"github.com/feilb/Silvia/chips/ads1115"
	"github.com/feilb/Silvia/chips/mcp9600"
	"github.com/feilb/Silvia/machine/manager"
	"github.com/feilb/Silvia/machine/pid"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

func main() {
	// **************************************************
	// periph.io Init
	// **************************************************
	host.Init()

	// I2C
	i2cBus, err := i2creg.Open("")

	if err != nil {
		fmt.Println(err)
		return
	}

	//GPIO
	boilerPin := gpioreg.ByName("GPIO19")
	boilerPin.Out(gpio.Low)

	valvePin := gpioreg.ByName("GPIO13")
	valvePin.Out(gpio.Low)

	pumpPin := gpioreg.ByName("GPIO26")
	pumpPin.Out(gpio.Low)

	zeroCrossingPin := gpioreg.ByName("GPIO24")
	zeroCrossingPin.In(gpio.PullUp, gpio.RisingEdge)

	// **************************************************
	// MCP9600 Thermocouple Sensor Init
	// **************************************************
	mcp := mcp9600.Dev{Device: &i2c.Dev{Addr: 0x60, Bus: i2cBus}}

	mcp.SetThermocoupleType(mcp9600.TypeK)
	mcp.SetFilterCoefficient(mcp9600.FilterCoefficient4)
	mcp.SetADCResolution(mcp9600.ADCResolution16Bit)
	mcp.SetColdJunctionResolution(mcp9600.ColdJuncitonResolution0p25)

	var CurTemp float64 = 0

	// **************************************************
	// ADS1115 ADC Init
	// **************************************************
	adc := ads1115.Dev{Device: &i2c.Dev{Addr: 0x48, Bus: i2cBus}}
	adcRange := ads1115.Range4p096

	adc.SetRange(adcRange)
	adc.SetMux(ads1115.Mux0v3)
	adc.SetMode(ads1115.ModeSingle)

	//var CurPress float64 = 0

	// **************************************************
	// PID Init
	// **************************************************
	pid := pid.PID{
		Kp: 0.035,
		//Ki: 0.0002727,
		Ki: 0.00015,
		Kd: 1.3 / 5,
		//Kd:     0,
		LastInput: make([]float64, 5),
		OutMax:    1,
		OutMin:    0,
	}

	// **************************************************
	// Manager Init
	// **************************************************
	mgr := manager.Manager{
		BrewSetpoint:  100,
		SteamSetpoint: 140,
		Mode:          manager.ModeOff,
		BoilerPin:     boilerPin,
		PumpPin:       pumpPin,
		ValvePin:      valvePin,
		Controller:    &pid,
		TempUpdate:    make(chan float64),
		Tick:          make(chan time.Time),
	}

	// **************************************************
	// REST Init
	// **************************************************
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, world!\n")
	})

	router.HandleFunc("/status", func(rw http.ResponseWriter, r *http.Request) {
		resp := api.GetStatus{
			Mode:               mgr.Mode.String(),
			BrewTime:           mgr.BrewTime().Seconds(),
			CurrentSetpoint:    mgr.CurrentSetpoint(),
			BrewSetpoint:       mgr.BrewSetpoint,
			SteamSetpoint:      mgr.SteamSetpoint,
			CurrentPressure:    0,
			CurrentTemperature: CurTemp,
		}

		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(resp)
	}).Methods("GET")

	router.HandleFunc("/autooff", func(rw http.ResponseWriter, r *http.Request) {
		resp := api.GetSetBool{
			Value: mgr.AutoOffEnabled(),
		}

		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(resp)
	}).Methods("GET")

	router.HandleFunc("/autooff/enable", func(rw http.ResponseWriter, r *http.Request) {
		mgr.EnableAutoOff()
	}).Methods("POST")

	router.HandleFunc("/autooff/disable", func(rw http.ResponseWriter, r *http.Request) {
		mgr.DisableAutoOff()
	}).Methods("POST")

	router.HandleFunc("/autooff/duration", func(rw http.ResponseWriter, r *http.Request) {
		resp := api.GetSetFloat{
			Value: mgr.AutoOffDuration().Minutes(),
		}

		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(resp)
	}).Methods("GET")

	router.HandleFunc("/autooff/duration", func(rw http.ResponseWriter, r *http.Request) {
		jsonVal := api.GetSetFloat{}
		err := json.NewDecoder(r.Body).Decode(&jsonVal)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		mgr.SetAutoOffDuration(time.Minute * time.Duration(jsonVal.Value))

	}).Methods("POST")

	router.HandleFunc("/mode", func(rw http.ResponseWriter, r *http.Request) {
		resp := api.GetSetString{
			Value: mgr.Mode.String(),
		}

		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(resp)
	}).Methods("GET")

	router.HandleFunc("/mode/{mode}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		newMode := manager.ModeFromString(vars["mode"])

		if newMode == manager.ModeInvalid {
			http.Error(rw, "Invalid Mode", http.StatusBadRequest)
		} else {
			mgr.Mode = newMode
		}
	}).Methods("POST")

	router.HandleFunc("/setpoint/{type}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		resp := api.GetSetFloat{}
		badReq := false

		switch strings.ToUpper(vars["type"]) {
		case "CURRENT":
			resp.Value = mgr.CurrentSetpoint()
		case "BREW":
			resp.Value = mgr.BrewSetpoint
		case "STEAM":
			resp.Value = mgr.SteamSetpoint
		default:
			badReq = true
		}

		if badReq {
			http.Error(rw, "", http.StatusNotFound)
		} else {
			rw.Header().Set("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(resp)
		}
	}).Methods("GET")

	router.HandleFunc("/setpoint/{type}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		jsonVal := api.GetSetFloat{}
		err := json.NewDecoder(r.Body).Decode(&jsonVal)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		switch strings.ToUpper(vars["type"]) {
		case "CURRENT":
			http.Error(rw, "", http.StatusMethodNotAllowed)
		case "BREW":
			mgr.BrewSetpoint = jsonVal.Value
		case "STEAM":
			mgr.SteamSetpoint = jsonVal.Value
		default:
			http.Error(rw, "", http.StatusNotFound)
		}
	}).Methods("POST")

	// **************************************************
	// START
	// **************************************************

	go func() {
		for {
			zeroCrossingPin.WaitForEdge(time.Second)

			mgr.Tick <- time.Now()
		}
	}()

	defer func() {
		boilerPin.Out(gpio.Low)
	}()

	go func() {
		for {
			CurTemp, _ = mcp.GetTemp()

			mgr.TempUpdate <- CurTemp
			time.Sleep(time.Second * 1)
		}
	}()

	defer mgr.Close()
	go mgr.Run()
	time.Sleep(time.Second)
	mgr.SetAutoOffDuration(time.Hour)
	mgr.EnableAutoOff()
	/*
		fmt.Println("Waiting 10s...")
		time.Sleep(time.Second * 10)
		mod.Setpoint = 0.05
		time.Sleep(time.Hour * 6)
	*/

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for range c {
			mgr.Close()
			os.Exit(1)
		}
	}()

	handler := cors.Default().Handler(router)
	fmt.Println("Listening on 8081...")
	http.ListenAndServe(":8081", handler)

}
