// Package nestStructure calls rest API Nest, puts it in the structure and gives somes functions
package nestStructure

//go:generate echo Go Generate!
//go:generate ./command/bindata.sh

import (
	"encoding/json"
	"math"
	"os"
	"time"

	"github.com/patrickalin/nest-api-go/assembly"
	http "github.com/patrickalin/http-go"
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
	url               string
	token             string
	NestStructure NestStructure
	mock              bool
}

// generate by http://mervine.net/json2struct
// you must replace your ThermostatID and you structure ID
type nestStructure struct {
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

type nestStructureShort struct {
	Devices struct {
		Thermostats interface{} `json:"thermostats"`
	} `json:"devices"`
	Metadata   Metadata    `json:"metadata"`
	Structures interface{} `json:"structures"`
}

type StructureID struct {
	Away        string   `json:"away"`
	CountryCode string   `json:"country_code"`
	Name        string   `json:"name"`
	StructureID string   `json:"structure_id"`
	Thermostats []string `json:"thermostats"`
}

type Metadata struct {
	AccessToken   string  `json:"access_token"`
	ClientVersion float64 `json:"client_version"`
}

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

// NestStructure is the Interface NestStructure
type NestStructure interface {
	GetDeviceID() string
	GetSoftwareVersion() string
	GetAmbientTemperatureC() float64
	GetTargetTemperatureC() float64
	GetAmbientTemperatureF() float64
	GetTargetTemperatureF() float64
	GetHumidity() float64
	GetAway() string
	ShowPrettyAll() int
}

func (nestInfo nestStructure) ShowPrettyAll() int {
	out, err := json.Marshal(nestInfo)
	if err != nil {
		fmt.Println("Error with parsing Json")
		mylog.Error.Fatal(err)
	}
	mylog.Trace.Printf("Decode:> \n %s \n\n", out)
	return 0
}

func (nestInfo nestStructureShort) ShowPrettyAll() int {
	out, err := json.Marshal(nestInfo)
	if err != nil {
		fmt.Println("Error with parsing Json")
		mylog.Error.Fatal(err)
	}
	mylog.Trace.Printf("Decode:> \n %s \n\n", out)
	return 0
}

func (nestInfo nestStructure) GetDeviceID() string {
	return nestInfo.Devices.Thermostats.ThermostatID.DeviceID
}

func (nestInfo nestStructure) GetSoftwareVersion() string {
	return nestInfo.Devices.Thermostats.ThermostatID.SoftwareVersion
}

func (nestInfo nestStructure) GetAmbientTemperatureC() float64 {
	return nestInfo.Devices.Thermostats.ThermostatID.AmbientTemperatureC
}

func (nestInfo nestStructure) GetTargetTemperatureF() float64 {
	return nestInfo.Devices.Thermostats.ThermostatID.TargetTemperatureF
}

func (nestInfo nestStructure) GetAmbientTemperatureF() float64 {
	return nestInfo.Devices.Thermostats.ThermostatID.AmbientTemperatureF
}

func (nestInfo nestStructure) GetTargetTemperatureC() float64 {
	return nestInfo.Devices.Thermostats.ThermostatID.TargetTemperatureC
}

func (nestInfo nestStructure) GetHumidity() float64 {
	return nestInfo.Devices.Thermostats.ThermostatID.Humidity
}

func (nestInfo nestStructure) GetAway() string {
	return nestInfo.Structures.StructureID.Away
}

// MakeNew calls Nest and get structureNest
func MakeNew(oneConfig config.ConfigStructure) NestStructure {

	var retry = 0
	var err error
	var duration = time.Minute * 5

	// get body from Rest API
        mylog.Trace.Printf("Get from Rest Nest API")
	myRest := rest.MakeNew()
	for retry < 5 {
		err = myRest.Get(oneConfig.NestURL)
		if err != nil {
			mylog.Error.Println(&nestError{err, "Problem with call rest, check the URL and the secret ID in the config file"})
			retry++
			time.Sleep(duration)
		}else  {
			retry=5
		}
	}

	if err != nil {
		mylog.Error.Fatal(&nestError{err, "Problem with call rest, check the URL and the secret ID in the config file"})
	}

	var nestInfo nestStructure
	var nestInfoShort nestStructureShort

	body := myRest.GetBody()

        mylog.Trace.Printf("Unmarshal the responce")
	//err = json.Unmarshal(body, &nestInfo)
	err = json.Unmarshal(body, &nestInfoShort)

	if err != nil {
		mylog.Error.Fatal(&nestError{err, "Problem with json to struct, problem in the struct ?"})
	}

	// not prety but works, one uid is use in the structure nest but I don't like that
	// so I found a work around
	listThermostatsInterface := nestInfoShort.Devices.Thermostats
	listThermostatsMaps := listThermostatsInterface.(map[string]interface{})

	var oneThermostat ThermostatID

	for _, value := range listThermostatsMaps {
		jsonString, _ := json.Marshal(value)
		json.Unmarshal(jsonString, &oneThermostat)
		nestInfo.Devices.Thermostats.ThermostatID = oneThermostat
	}

	listStructuresInterface := nestInfoShort.Structures
	listStructuresMaps := listStructuresInterface.(map[string]interface{})

	var oneStructure StructureID

	for _, value := range listStructuresMaps {
		jsonString, _ := json.Marshal(value)
		json.Unmarshal(jsonString, &oneStructure)
		nestInfo.Structures.StructureID = oneStructure
	}

	nestInfo.Metadata = nestInfoShort.Metadata
	// not prety but works end

	nestInfo.ShowPrettyAll()

	return nestInfo
}












// New calls nest and get structurenest
func New(nestURL, nestToken string, mock bool, l *logrus.Logger) Nest {
	initLog(l)

	logDebug(funcName(), "New nest structure", nestURL)

	// Read mock file
	if mock {
		logWarn(funcName(), "Mock activated !!!", "")
		mockFileByte = readFile(mockFile)
	}

	rest = http.New(log)

	return &nest{url: nestURL, token: nestToken, mock: mock}
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

//GetTimeStamp returns the timestamp give by Nest
func (nest *nest) GetTimeStamp() time.Time {
	return time.Unix(int64(nest.NestStructure.Data.TS), 0)
}

//GetCity returns the city name
func (nest *nest) GetCity() string {
	return nest.NestStructure.CityName
}

//GetDeviceID returns the Device Id
func (nest *nest) GetDeviceID() string {
	return nest.NestStructure.DeviceID
}

//GetNumOfFollowers returns the number of followers
func (nest *nest) GetNumOfFollowers() int {
	return int(nest.NestStructure.NumOfFollowers)
}

//GetIndexUV returns the UV index from 1 to 11
func (nest *nest) GetIndexUV() string {
	return nest.NestStructure.Storm.UVIndex
}

//IsNight returns true if it's the night
func (nest *nest) IsNight() bool {
	return nest.NestStructure.Data.Night
}

//GetTemperatureFahrenheit returns temperature in Fahrenheit
func (nest *nest) GetTemperatureFahrenheit() float64 {
	return nest.NestStructure.Data.TemperatureF
}

//GetTemperatureCelsius returns temperature in Celsius
func (nest *nest) GetTemperatureCelsius() float64 {
	return nest.NestStructure.Data.TemperatureC
}

//GetHumidity returns humidity %
func (nest *nest) GetHumidity() float64 {
	return nest.NestStructure.Data.Humidity
}

//GetPressureHPa returns pressure in HPa
func (nest *nest) GetPressureHPa() float64 {
	return nest.NestStructure.Data.Pressurehpa
}

//GetPressureInHg returns pressure in InHg
func (nest *nest) GetPressureInHg() float64 {
	return nest.NestStructure.Data.Pressure
}

//GetWindDirection returns wind direction (N,S,W,E, ...)
func (nest *nest) GetWindDirection() string {
	return nest.NestStructure.Storm.WindDirection
}

//GetWindGustMph returns Wind in Mph
func (nest *nest) GetWindGustMph() float64 {
	return nest.NestStructure.Storm.WindGust
}

//GetWindGustMs returns Wind in Ms
func (nest *nest) GetWindGustMs() float64 {
	return (nest.NestStructure.Storm.WindGust * 1.61)
}

//GetSustainedWindSpeedMph returns Sustained Wind Speed in Mph
func (nest *nest) GetSustainedWindSpeedMph() float64 {
	return nest.NestStructure.Storm.SustainedWindSpeed
}

//GetSustainedWindSpeedMs returns Sustained Wind Speed in Ms
func (nest *nest) GetSustainedWindSpeedMs() float64 {
	return (nest.NestStructure.Storm.SustainedWindSpeed * 1.61)
}

//IsRain returns true if it's rain
func (nest *nest) IsRain() bool {
	return nest.NestStructure.Data.Rain
}

//GetRainDailyIn returns rain daily in In
func (nest *nest) GetRainDailyIn() float64 {
	return nest.NestStructure.Storm.RainDaily
}

//GetRainIn returns total rain in In
func (nest *nest) GetRainIn() float64 {
	return nest.NestStructure.Storm.Rainin
}

//GetRainRateIn returns rain in In
func (nest *nest) GetRainRateIn() float64 {
	return nest.NestStructure.Storm.RainRate
}

//GetRainDailyMm returns rain daily in mm
func (nest *nest) GetRainDailyMm() float64 {
	return nest.NestStructure.Storm.RainDaily
}

//GetRainMm returns total rain in mm
func (nest *nest) GetRainMm() float64 {
	return nest.NestStructure.Storm.Rainmm
}

//GetRainRateMm returns rain in mm
func (nest *nest) GetRainRateMm() float64 {
	return nest.NestStructure.Storm.RainRate
}

//GetSustainedWindSpeedkmh returns Sustained Wind in Km/h
func (nest *nest) GetSustainedWindSpeedkmh() float64 {
	return nest.NestStructure.Storm.SustainedWindSpeedkmh
}

//GetWindGustkmh returns Wind in Km/h
func (nest *nest) GetWindGustkmh() float64 {
	return nest.NestStructure.Storm.WindGustkmh
}

func (nest *nest) GetLastCall() string {
	return nest.NestStructure.LastCall
}

func (nest *nest) GetTS() float64 {
	return nest.NestStructure.Data.TS
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

// ShowPrettyAll prints to the console the JSON
func (nest *nest) showPrettyAll() {
	out, err := json.Marshal(nest)
	checkErr(err, funcName(), "Error with parsing Json", string(out))
}

//Read file and return []byte
func readFile(fileName string) []byte {
	fileByte, err := assembly.Asset(fileName)
	checkErr(err, funcName(), "Error reading the file", fileName)
	return fileByte
}

//Call rest and refresh the structure
func (nest *nest) refreshFromRest() {
	tock := []string{nest.token}

	var headers map[string][]string
	headers = make(map[string][]string)
	headers["Authorization"] = tock

	var retry = 0
	for retry < 5 {
		if err := rest.GetWithHeaders(nest.url, headers); err != nil {
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
	var nestArray []NestStructure
	if err := json.Unmarshal(body, &nestArray); err != nil {
		logFatal(err, funcName(), "Problem with json to struct", string(body))
	}
	nest.NestStructure = nestArray[0]
	nest.NestStructure.Data.TemperatureC = toFixed(((nest.NestStructure.Data.TemperatureF - 32.00) * 5.00 / 9.00), 2)
	nest.NestStructure.Data.Pressurehpa = toFixed((nest.NestStructure.Data.Pressure * 33.8638815), 2)

	nest.NestStructure.Storm.WindGustms = toFixed(nest.NestStructure.Storm.WindGust*0.44704, 2)
	nest.NestStructure.Storm.WindGustkmh = toFixed(nest.NestStructure.Storm.WindGust*1.60934, 2)
	nest.NestStructure.Storm.SustainedWindSpeedms = toFixed(nest.NestStructure.Storm.SustainedWindSpeed*0.44704, 2)
	nest.NestStructure.Storm.SustainedWindSpeedkmh = toFixed(nest.NestStructure.Storm.SustainedWindSpeed*1.60934, 2)

	nest.NestStructure.Storm.RainDailymm = toFixed(nest.NestStructure.Storm.RainDaily*25.4, 2)
	nest.NestStructure.Storm.RainRatemm = toFixed(nest.NestStructure.Storm.RainRate*25.4, 2)
	nest.NestStructure.Storm.Rainmm = toFixed(nest.NestStructure.Storm.Rainin*25.4, 2)
	nest.NestStructure.LastCall = time.Now().Format("2006-01-02 15:04:05")

	logDebug(funcName(), "Refresh From Body", nest.NestStructure.LastCall)
}
