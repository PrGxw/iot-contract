package contract

import (
	"encoding/json"
	"testing"
)

func TestResult_JSONOmitsEmpty(t *testing.T) {
	raw, err := json.Marshal(Result{Outcome: OutcomeDone})
	if err != nil || string(raw) != `{"outcome":"done"}` {
		t.Fatalf("got %s, %v", raw, err)
	}
}

func TestResult_FailedKeepsCodeAndDetail(t *testing.T) {
	raw, _ := json.Marshal(Result{Outcome: OutcomeFailed, Code: CodeDeviceTypeUnrecognised, Detail: "roles [camera]"})
	var back Result
	if err := json.Unmarshal(raw, &back); err != nil || back.Code != CodeDeviceTypeUnrecognised || back.Detail != "roles [camera]" {
		t.Fatalf("got %+v, %v", back, err)
	}
}
