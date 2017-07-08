package nestStructure

import (
	"os"
	"testing"
)

var mynestTest1 Nest
var mynestTest2 Nest

const testFile1 = "testcase/test1.json"
const testFile2 = "testcase/test2.json"

func TestMain(m *testing.M) {
	mynestTest1 = New("", "", true, nil)
	mynestTest1.RefreshFromBody(readFile(testFile1))
	mynestTest2 = New("", "", true, nil)
	mynestTest2.RefreshFromBody(readFile(testFile2))
	os.Exit(m.Run())
}

func TestNestStructure_GetHumidity(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 50},
		{"Test2", mynestTest2, 70},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetHumidity(); got != tt.want {
				t.Errorf("NestStructure.GetHumidity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetSoftwareVersion(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   string
	}{
		{"Test1", mynestTest1, "5.6.1-4"},
		{"Test2", mynestTest2, "5.6.1-4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetSoftwareVersion(); got != tt.want {
				t.Errorf("NestStructure.GetSoftwareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetAmbientTemperatureC(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 27.5},
		{"Test2", mynestTest2, 28.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetAmbientTemperatureC(); got != tt.want {
				t.Errorf("NestStructure.GetAmbientTemperatureC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetDeviceId(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   string
	}{
		{"Test1", mynestTest1, "oJHB1ha6NGOT9493h-fcJY--gS80WzmN"},
		{"Test2", mynestTest2, "oJHB1ha6NGOT9493h-fcJY--gS80WzmO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetDeviceID(); got != tt.want {
				t.Errorf("NestStructure.GetDeviceID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetAmbientTemperatureF(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 82},
		{"Test2", mynestTest2, 87},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetAmbientTemperatureF(); got != tt.want {
				t.Errorf("NestStructure.GetAmbientTemperatureF() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestNewNest(t *testing.T) {
	type args struct {
		nestURL   string
		nestToken string
	}
	tests := []struct {
		name string
		args args
		want NestStructure
	}{
		{"Error token", args{"https://api.nest.com/api/skydata/", ""}, mynestTest1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNest(tt.args.nestURL, tt.args.nestToken); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNest() = %v, want %v", got, tt.want)
			}
		})
	}
}*/

/*
func TestNewNestFromBody(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name string
		args args
		want Nest
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNestFromBody(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNestFromBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
/*
func TestNestStructure_GetIndexUV(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   string
	}{
		{"Test1", mynestTest1, "1"},
		{"Test2", mynestTest2, "3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetIndexUV(); got != tt.want {
				t.Errorf("NestStructure.GetIndexUV() = %v, want %v", got, tt.want)
			}
		})
	}

}
*/
func TestNestStructure_GetTargetTemperatureF(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 72},
		{"Test2", mynestTest2, 74},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetTargetTemperatureF(); got != tt.want {
				t.Errorf("NestStructure.GetTargetTemperatureF() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*

func TestNestStructure_GetTemperatureCelsius(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 21.55},
		{"Test2", mynestTest2, 18.77},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetTemperatureCelsius(); got != tt.want {
				t.Errorf("NestStructure.GetTemperatureCelsius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetHumidity(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 64},
		{"Test2", mynestTest2, 43},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetHumidity(); got != tt.want {
				t.Errorf("NestStructure.GetHumidity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetPressureInHg(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 29.41},
		{"Test2", mynestTest2, 49.41},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetPressureInHg(); got != tt.want {
				t.Errorf("NestStructure.GetPressureInHg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetPressureHPa(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 995.94},
		{"Test2", mynestTest2, 1673.21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetPressureHPa(); got != tt.want {
				t.Errorf("NestStructure.GetPressureHPa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetWindDirection(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   string
	}{
		{"Test1", mynestTest1, "E"},
		{"Test2", mynestTest2, "W"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetWindDirection(); got != tt.want {
				t.Errorf("NestStructure.GetWindDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetWindGustMph(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 0},
		{"Test2", mynestTest2, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetWindGustMph(); got != tt.want {
				t.Errorf("NestStructure.GetWindGustMph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetWindGustMs(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 0},
		{"Test2", mynestTest2, 33.81},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetWindGustMs(); got != tt.want {
				t.Errorf("NestStructure.GetWindGustMs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetSustainedWindSpeedMph(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 0},
		{"Test2", mynestTest2, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetSustainedWindSpeedMph(); got != tt.want {
				t.Errorf("NestStructure.GetSustainedWindSpeedMph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetSustainedWindSpeedMs(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 0},
		{"Test2", mynestTest2, 19.32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetSustainedWindSpeedMs(); got != tt.want {
				t.Errorf("NestStructure.GetSustainedWindSpeedMs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_IsRain(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   bool
	}{
		{"Test1", mynestTest1, true},
		{"Test2", mynestTest2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.IsRain(); got != tt.want {
				t.Errorf("NestStructure.IsRain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetRainDailyIn(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 0},
		{"Test2", mynestTest2, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetRainDailyIn(); got != tt.want {
				t.Errorf("NestStructure.GetRainDailyIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetRainIn(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 0},
		{"Test2", mynestTest2, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetRainIn(); got != tt.want {
				t.Errorf("NestStructure.GetRainIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetRainRateIn(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 0},
		{"Test2", mynestTest2, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetRainRateIn(); got != tt.want {
				t.Errorf("NestStructure.GetRainRateIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetRainDailyMm(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 0},
		{"Test2", mynestTest2, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetRainDailyMm(); got != tt.want {
				t.Errorf("NestStructure.GetRainDailyMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetRainMm(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 0},
		{"Test2", mynestTest2, 406.4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetRainMm(); got != tt.want {
				t.Errorf("NestStructure.GetRainMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetRainRateMm(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 0},
		{"Test2", mynestTest2, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetRainRateMm(); got != tt.want {
				t.Errorf("NestStructure.GetRainRateMm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetSustainedWindSpeedkmh(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 0},
		{"Test2", mynestTest2, 19.31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetSustainedWindSpeedkmh(); got != tt.want {
				t.Errorf("NestStructure.GetSustainedWindSpeedkmh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetWindGustkmh(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 0},
		{"Test2", mynestTest2, 33.8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetWindGustkmh(); got != tt.want {
				t.Errorf("NestStructure.GetWindGustkmh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nest_GetTS(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 1496365207},
		{"Test2", mynestTest2, 1496345207},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetTS(); got != tt.want {
				t.Errorf("NestStructure.GetTS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nest_Refresh(t *testing.T) {
	type fields struct {
		url           string
		token         string
		NestStructure NestStructure
		mock          bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nest := &nest{
				url:           tt.fields.url,
				token:         tt.fields.token,
				NestStructure: tt.fields.NestStructure,
				mock:          tt.fields.mock,
			}
			nest.Refresh()
		})
	}
}

func Test_nest_refreshFromRest(t *testing.T) {
	type fields struct {
		url           string
		token         string
		NestStructure NestStructure
		mock          bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nest := &nest{
				url:           tt.fields.url,
				token:         tt.fields.token,
				NestStructure: tt.fields.NestStructure,
				mock:          tt.fields.mock,
			}
			nest.refreshFromRest()
		})
	}
}

func Test_nest_RefreshFromBody(t *testing.T) {
	mynestTest1.Refresh()
	mynestTest2.Refresh()
	mynestTest1.GetNestStruct()
	mynestTest2.GetNestStruct()
	mynestTest1.GetLastCall()
	mynestTest2.GetLastCall()
}

func BenchmarkNestStructureIsNight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mynest := New("", "", true, nil)
		mynest.RefreshFromBody(readFile(testFile1))
		if got := mynest.IsNight(); got != false {
			b.Errorf("NestStructure.IsNight() = %v, want %v", got, false)
		}
	}
}
*/
