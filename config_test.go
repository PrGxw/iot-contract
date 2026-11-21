package contract

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestDisplayColors_OrderAndValid(t *testing.T) {
	want := []DisplayColor{"red", "green", "yellow", "purple", "blue", "light_yellow", "white"}
	if got := DisplayColors(); !reflect.DeepEqual(got, want) {
		t.Fatalf("colors = %v, want %v", got, want)
	}
	for _, c := range want {
		if !c.Valid() {
			t.Errorf("%q should be valid", c)
		}
	}
	if DisplayColor("orange").Valid() {
		t.Error("orange must not be valid")
	}
}

func TestArrowDirection_Valid(t *testing.T) {
	for _, d := range []ArrowDirection{DirectionUp, DirectionDown, DirectionLeft, DirectionRight} {
		if !d.Valid() {
			t.Errorf("%q should be valid", d)
		}
	}
	if ArrowDirection("diagonal").Valid() {
		t.Error("diagonal must not be valid")
	}
}

func TestConfigKinds_StableOrder(t *testing.T) {
	want := []ConfigKind{ConfigDisplayConfig, ConfigDisplayHolds}
	if got := ConfigKinds(); !reflect.DeepEqual(got, want) {
		t.Fatalf("kinds = %v, want %v", got, want)
	}
}

func TestConfigDocument_EmptyItemsMarshalAsArray(t *testing.T) {
	b, err := json.Marshal(ConfigDocument{Kind: ConfigDisplayHolds})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), `"items":[]`) {
		t.Fatalf("want items:[] in %s", b)
	}
}

func TestConfigReport_EmptyCollectionsMarshalAsArrays(t *testing.T) {
	b, err := json.Marshal(ConfigReport{Kinds: []KindReport{{Kind: ConfigDisplayConfig}}})
	if err != nil {
		t.Fatal(err)
	}
	if got := string(b); got != `{"kinds":[{"kind":"display_config","items":[]}]}` {
		t.Fatal(got)
	}
	b, err = json.Marshal(ConfigReport{})
	if err != nil {
		t.Fatal(err)
	}
	if got := string(b); got != `{"kinds":[]}` {
		t.Fatal(got)
	}
}

func TestConfigReport_RoundTrip(t *testing.T) {
	in := ConfigReport{Kinds: []KindReport{
		{Kind: ConfigDisplayConfig, Items: []HeldItem{
			{Key: "AA:BB:CC:DD:EE:01", Version: "v1"},
			{Key: "AA:BB:CC:DD:EE:02", Version: "v1", RefusedVersion: "v2", Reason: RefusalInvalidConfig},
			{Key: "AA:BB:CC:DD:EE:03", RefusedVersion: "v9", Reason: RefusalStoreFailed},
		}},
		{Kind: ConfigDisplayHolds, Items: []HeldItem{}},
	}}
	b, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	var out ConfigReport
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(in, out) {
		t.Fatalf("round trip: %+v != %+v", out, in)
	}
}

func TestConfigDocument_RoundTripWithContent(t *testing.T) {
	content, _ := json.Marshal(DisplayConfigContent{Screens: []DisplayScreen{
		{Kind: "arrow", Direction: DirectionLeft, Color: ColorGreen},
		{Kind: "text", Digits: 3, Color: ColorWhite, Bays: []CameraBay{{CameraMAC: "AA:BB:CC:DD:EE:FF", Bay: 2}}},
	}})
	in := ConfigDocument{Kind: ConfigDisplayConfig, Items: []ConfigItem{{Key: "K", Version: "v", Content: content, Retry: true}}}
	b, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	var out ConfigDocument
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	var got DisplayConfigContent
	if err := json.Unmarshal(out.Items[0].Content, &got); err != nil {
		t.Fatal(err)
	}
	if !out.Items[0].Retry || got.Screens[1].Bays[0].Bay != 2 || got.Screens[0].Color != ColorGreen {
		t.Fatalf("lost data: %+v %+v", out, got)
	}
}

func TestDisplayHoldsContent_EmptyHeldIsArray(t *testing.T) {
	b, _ := json.Marshal(DisplayHoldsContent{})
	if string(b) != `{"held":[]}` {
		t.Fatal(string(b))
	}
}
