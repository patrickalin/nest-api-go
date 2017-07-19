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

func TestNestStructure_GetTargetTemperatureC(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   float64
	}{
		{"Test1", mynestTest1, 22},
		{"Test2", mynestTest2, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetTargetTemperatureC(); got != tt.want {
				t.Errorf("NestStructure.GetTargetTemperatureC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestStructure_GetAway(t *testing.T) {
	tests := []struct {
		name   string
		fields Nest
		want   string
	}{
		{"Test1", mynestTest1, "home"},
		{"Test2", mynestTest2, "away"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetAway(); got != tt.want {
				t.Errorf("NestStructure.GetAway() = %v, want %v", got, tt.want)
			}
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
		fields Nest
	}{
		{"Test1", mynestTest1},
		{"Test2", mynestTest2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Refresh()
		})
	}
}

func Benchmark_GetHumidity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mybloomsky := New("", "", true, nil)
		mybloomsky.RefreshFromBody(readFile(testFile1))
		if got := mybloomsky.GetHumidity(); got != 50 {
			b.Errorf("BloomskyStructure.IsNight() = %v, want %v", got, false)
		}
	}
}
