// Package nestStructure calls rest API Nest, puts it in the structure and gives somes functions
package nestStructure

//go:generate echo Go Generate!
//go:generate ./command/bindata.sh

import (
	"encoding/json"
	"math"
	"os"
	"time"

	http "github.com/patrickalin/http-go"
	"github.com/patrickalin/nest-api-go/assembly"
	"github.com/sirupsen/logrus"
)

const (
	logFile  = "nestapi.log"
	mockFile = "mock/mock.json"
)

var (
	log          *logrus.Logger
	rest         http.HTTP
	mockFileByte []byte
)

type nest struct {
	url           string
	token         string
	NestStructure NestStructure
	mock          bool
	LastCall      string
}

// NestStructure generate by http://mervine.net/json2struct
// you must replace your ThermostatID and you structure ID
type NestStructure struct {
	Devices struct {
		Thermostats struct {
			ThermostatID ThermostatID `json:"noNeeded"`
		} `json:"thermostats"`
	} `json:"devices"`
	Metadata   Metadata `json:"metadata"`
	Structures struct {
		StructureID StructureID `json:"noNeeded"`
	} `json:"structures"`
}

//nestStructureShort is the structure from the API Nest
type nestStructureShort struct {
	Devices struct {
		Thermostats interface{} `json:"thermostats"`
	} `json:"devices"`
	Metadata   Metadata    `json:"metadata"`
	Structures interface{} `json:"structures"`
}

//StructureID sub structure from NestStructure
type StructureID struct {
	Away        string   `json:"away"`
	CountryCode string   `json:"country_code"`
	Name        string   `json:"name"`
	StructureID string   `json:"structure_id"`
	Thermostats []string `json:"thermostats"`
}

//Metadata sub structure from NestStructure
type Metadata struct {
	AccessToken   string  `json:"access_token"`
	ClientVersion float64 `json:"client_version"`
}

//ThermostatID sub structure from NestStructure
type ThermostatID struct {
	AmbientTemperatureC    float64 `json:"ambient_temperature_c"`
	AmbientTemperatureF    float64 `json:"ambient_temperature_f"`
	AwayTemperatureHighC   float64 `json:"away_temperature_high_c"`
	AwayTemperatureHighF   float64 `json:"away_temperature_high_f"`
	AwayTemperatureLowC    float64 `json:"away_temperature_low_c"`
	AwayTemperatureLowF    float64 `json:"away_temperature_low_f"`
	CanCool                bool    `json:"can_cool"`
	CanHeat                bool    `json:"can_heat"`
	DeviceID               string  `json:"device_id"`
	FanTimerActive         bool    `json:"fan_timer_active"`
	HasFan                 bool    `json:"has_fan"`
	HasLeaf                bool    `json:"has_leaf"`
	Humidity               float64 `json:"humidity"`
	HvacMode               string  `json:"hvac_mode"`
	IsOnline               bool    `json:"is_online"`
	IsUsingEmergencyHeat   bool    `json:"is_using_emergency_heat"`
	LastConnection         string  `json:"last_connection"`
	Locale                 string  `json:"locale"`
	Name                   string  `json:"name"`
	NameLong               string  `json:"name_long"`
	SoftwareVersion        string  `json:"software_version"`
	StructureID            string  `json:"structure_id"`
	TargetTemperatureC     float64 `json:"target_temperature_c"`
	TargetTemperatureF     float64 `json:"target_temperature_f"`
	TargetTemperatureHighC float64 `json:"target_temperature_high_c"`
	TargetTemperatureHighF float64 `json:"target_temperature_high_f"`
	TargetTemperatureLowC  float64 `json:"target_temperature_low_c"`
	TargetTemperatureLowF  float64 `json:"target_temperature_low_f"`
	TemperatureScale       string  `json:"temperature_scale"`
}

// Nest is the Interface NestStructure
type Nest interface {
	GetDeviceID() string
	GetSoftwareVersion() string
	GetAmbientTemperatureC() float64
	GetTargetTemperatureC() float64
	GetAmbientTemperatureF() float64
	GetTargetTemperatureF() float64
	GetHumidity() float64
	GetAway() string
	GetNestStruct() NestStructure
	Refresh()
	RefreshFromBody(body []byte)
	GetLastCall() string
}

func (nest *nest) GetDeviceID() string {
	return nest.NestStructure.Devices.Thermostats.ThermostatID.DeviceID
}

func (nest *nest) GetSoftwareVersion() string {
	return nest.NestStructure.Devices.Thermostats.ThermostatID.SoftwareVersion
}

func (nest *nest) GetAmbientTemperatureC() float64 {
	return nest.NestStructure.Devices.Thermostats.ThermostatID.AmbientTemperatureC
}

func (nest *nest) GetAmbientTemperatureF() float64 {
	return nest.NestStructure.Devices.Thermostats.ThermostatID.AmbientTemperatureF
}

func (nest *nest) GetTargetTemperatureF() float64 {
	return nest.NestStructure.Devices.Thermostats.ThermostatID.TargetTemperatureF
}

func (nest *nest) GetTargetTemperatureC() float64 {
	return nest.NestStructure.Devices.Thermostats.ThermostatID.TargetTemperatureC
}

func (nest *nest) GetHumidity() float64 {
	return nest.NestStructure.Devices.Thermostats.ThermostatID.Humidity
}

func (nest *nest) GetAway() string {
	return nest.NestStructure.Structures.StructureID.Away
}

// New calls Nest and get structureNest
func New(url, token string, mock bool, l *logrus.Logger) Nest {

	initLog(l)

	logDebug(funcName(), "New nest structure", url)

	// Read mock file
	if mock {
		logWarn(funcName(), "Mock activated !!!", "")
		mockFileByte = readFile(mockFile)
	}

	rest = http.New(log)

	return &nest{url: url, token: token, mock: mock}

}

func (nest *nest) Refresh() {
	if nest.mock {
		nest.RefreshFromBody(mockFileByte)
		return
	}
	nest.refreshFromRest()
}

func (nest *nest) GetNestStruct() NestStructure {
	return nest.NestStructure
}

func (nest *nest) GetLastCall() string {
	return bloomsky.BloomskyStructure.LastCall
}

/* Func private ------------------------------------ */

//Init the logger
func initLog(l *logrus.Logger) {
	if l != nil {
		log = l
		logDebug(funcName(), "Use the logger pass in New", "")
		return
	}

	log = logrus.New()

	logDebug(funcName(), "Create new logger", "")

	log.Formatter = new(logrus.TextFormatter)

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY, 0666)
	checkErr(err, funcName(), "Failed to log to file, using default stderr", "")

	log.Out = file
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

//Read file and return []byte
func readFile(fileName string) []byte {
	fileByte, err := assembly.Asset(fileName)
	checkErr(err, funcName(), "Error reading the file", fileName)
	return fileByte
}

//Call rest and refresh the structure
func (nest *nest) refreshFromRest() {
	var retry = 0
	for retry < 5 {
		if err := rest.Get(nest.url + nest.token); err != nil {
			logFatal(err, funcName(), "Problem with call rest, check the URL and the secret ID in the config file", nest.url)
			retry++
			time.Sleep(time.Minute * 5)
		} else {
			retry = 5
		}
	}

	nest.RefreshFromBody(rest.GetBody())
}

//Refresh from body without call rest
func (nest *nest) RefreshFromBody(body []byte) {

	var nestInfoShort nestStructureShort

	err := json.Unmarshal(body, &nestInfoShort)
	checkErr(err, funcName(), "Problem with json to struct, problem in the struct ?", "")

	// not prety but works, one uid is use in the structure nest but I don't like that
	// so I found a work around
	listThermostatsInterface := nestInfoShort.Devices.Thermostats
	listThermostatsMaps := listThermostatsInterface.(map[string]interface{})

	var oneThermostat ThermostatID

	for _, value := range listThermostatsMaps {
		jsonString, _ := json.Marshal(value)
		err = json.Unmarshal(jsonString, &oneThermostat)
		checkErr(err, funcName(), "Problem Unmarshal jsonString", "")
		nest.NestStructure.Devices.Thermostats.ThermostatID = oneThermostat
	}

	listStructuresInterface := nestInfoShort.Structures
	listStructuresMaps := listStructuresInterface.(map[string]interface{})

	var oneStructure StructureID

	for _, value := range listStructuresMaps {
		jsonString, _ := json.Marshal(value)
		err = json.Unmarshal(jsonString, &oneStructure)
		checkErr(err, funcName(), "Problem Unmarshal jsonString2", "")
		nest.NestStructure.Structures.StructureID = oneStructure
	}

	nest.NestStructure.Metadata = nestInfoShort.Metadata

	nest.LastCall = time.Now().Format("2006-01-02 15:04:05")

	logDebug(funcName(), "Refresh From Body", nest.LastCall)
}
