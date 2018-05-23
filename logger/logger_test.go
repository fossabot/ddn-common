package logger

import (
	"testing"
)

func TestShouldLog(t *testing.T) {
	tests := []struct {
		Level     LogLevel
		CheckName string
		Check     map[LogLevel]bool
	}{
		{CheckName: "FATAL", Level: FATAL, Check: map[LogLevel]bool{FATAL: true, ERROR: false, WARN: false, INFO: false, DEBUG: false}},
		{CheckName: "ERROR", Level: ERROR, Check: map[LogLevel]bool{FATAL: true, ERROR: true, WARN: false, INFO: false, DEBUG: false}},
		{CheckName: "WARN", Level: WARN, Check: map[LogLevel]bool{FATAL: true, ERROR: true, WARN: true, INFO: false, DEBUG: false}},
		{CheckName: "INFO", Level: INFO, Check: map[LogLevel]bool{FATAL: true, ERROR: true, WARN: true, INFO: true, DEBUG: false}},
		{CheckName: "DEBUG", Level: DEBUG, Check: map[LogLevel]bool{FATAL: true, ERROR: true, WARN: true, INFO: true, DEBUG: true}},
	}

	for _, test := range tests {
		Level = test.Level
		for lvl, expected := range test.Check {
			if shouldLog(lvl) != expected {
				t.Errorf("shouldLog(%s) <-> %s is expected to be %t", lvl, test.CheckName, expected)
			}
		}
	}
}

func TestParse(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name    string
		args    args
		want    LogLevel
		wantErr bool
	}{
		{"valid debug", args{"debug"}, DEBUG, false},
		{"valid DEBUG", args{"DEBUG"}, DEBUG, false},
		{"valid info", args{"info"}, INFO, false},
		{"valid INFO", args{"INFO"}, INFO, false},
		{"valid warn", args{"warn"}, WARN, false},
		{"valid WARN", args{"WARN"}, WARN, false},
		{"valid error", args{"error"}, ERROR, false},
		{"valid ERROR", args{"ERROR"}, ERROR, false},
		{"valid fatal", args{"fatal"}, FATAL, false},
		{"valid FATAL", args{"FATAL"}, FATAL, false},

		{"invalid", args{"asdf"}, 0, true},
		{"INVALID", args{"NOTHING"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
