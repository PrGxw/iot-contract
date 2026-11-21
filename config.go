package contract

import "encoding/json"

// ConfigKind names one kind of gateway configuration: a closed set. The backend
// decides what each gateway should hold per kind; the gateway reports what it holds.
type ConfigKind string

const (
	// ConfigDisplayConfig is a display's screens (what each screen shows and counts).
	ConfigDisplayConfig ConfigKind = "display_config"
	// ConfigDisplayHolds is the camera bays a display treats as held.
	ConfigDisplayHolds ConfigKind = "display_holds"
)

// ConfigKinds returns every kind in a stable order.
func ConfigKinds() []ConfigKind {
	return []ConfigKind{ConfigDisplayConfig, ConfigDisplayHolds}
}

// DisplayColor is a screen colour by name. Firmware colour codes exist only in
// the gateway's display adapter.
type DisplayColor string

const (
	ColorRed         DisplayColor = "red"
	ColorGreen       DisplayColor = "green"
	ColorYellow      DisplayColor = "yellow"
	ColorPurple      DisplayColor = "purple"
	ColorBlue        DisplayColor = "blue"
	ColorLightYellow DisplayColor = "light_yellow"
	ColorWhite       DisplayColor = "white"
)

// DisplayColors returns every colour in a stable order.
func DisplayColors() []DisplayColor {
	return []DisplayColor{ColorRed, ColorGreen, ColorYellow, ColorPurple, ColorBlue, ColorLightYellow, ColorWhite}
}

// Valid reports whether c is one of the closed set of colour names.
func (c DisplayColor) Valid() bool {
	for _, v := range DisplayColors() {
		if c == v {
			return true
		}
	}
	return false
}

// ArrowDirection is the way an arrow screen points.
type ArrowDirection string

const (
	DirectionUp    ArrowDirection = "up"
	DirectionDown  ArrowDirection = "down"
	DirectionLeft  ArrowDirection = "left"
	DirectionRight ArrowDirection = "right"
)

// Valid reports whether d is one of the four directions.
func (d ArrowDirection) Valid() bool {
	switch d {
	case DirectionUp, DirectionDown, DirectionLeft, DirectionRight:
		return true
	}
	return false
}

// CameraBay is one bay of one camera. Bay is the camera's own bay number, 1-6.
type CameraBay struct {
	CameraMAC string `json:"camera_mac"`
	Bay       int    `json:"bay"`
}

// Screen kinds of a DisplayScreen.
const (
	ScreenArrow = "arrow"
	ScreenText  = "text"
)

// DisplayScreen is one screen of a display. Kind is ScreenArrow or ScreenText;
// Direction applies to arrows, Digits and Bays to text screens (the bays it counts).
type DisplayScreen struct {
	Kind      string         `json:"kind"`
	Direction ArrowDirection `json:"direction,omitempty"`
	Digits    int            `json:"digits,omitempty"`
	Color     DisplayColor   `json:"color"`
	Bays      []CameraBay    `json:"bays,omitempty"`
}

// DisplayConfigContent is the content of a display_config item.
type DisplayConfigContent struct {
	Screens []DisplayScreen `json:"screens"`
}

// MarshalJSON writes an empty Screens as [].
func (c DisplayConfigContent) MarshalJSON() ([]byte, error) {
	type plain DisplayConfigContent
	if c.Screens == nil {
		c.Screens = []DisplayScreen{}
	}
	return json.Marshal(plain(c))
}

// DisplayHoldsContent is the content of a display_holds item.
type DisplayHoldsContent struct {
	Held []CameraBay `json:"held"`
}

// MarshalJSON writes an empty Held as [].
func (c DisplayHoldsContent) MarshalJSON() ([]byte, error) {
	type plain DisplayHoldsContent
	if c.Held == nil {
		c.Held = []CameraBay{}
	}
	return json.Marshal(plain(c))
}

// ConfigItem is one unit of a kind. Version is the backend's content version;
// the gateway stores and echoes it and never computes it. Retry asks the gateway
// to try this version again even if it refused it before.
type ConfigItem struct {
	Key     string          `json:"key"`
	Version string          `json:"version"`
	Content json.RawMessage `json:"content"`
	Retry   bool            `json:"retry,omitempty"`
}

// ConfigDocument is the whole of one kind for one gateway: every item it should
// hold. An item absent from the document means "drop it". Sent downward wrapped
// in an Envelope of KindConfig.
type ConfigDocument struct {
	Kind  ConfigKind   `json:"kind"`
	Items []ConfigItem `json:"items"`
}

// MarshalJSON writes an empty Items as [].
func (d ConfigDocument) MarshalJSON() ([]byte, error) {
	type plain ConfigDocument
	if d.Items == nil {
		d.Items = []ConfigItem{}
	}
	return json.Marshal(plain(d))
}

// RefusalReason says why a gateway refused a version.
type RefusalReason string

const (
	RefusalInvalidConfig RefusalReason = "invalid_config"
	RefusalStoreFailed   RefusalReason = "store_failed"
)

// HeldItem is what a gateway reports for one item key: the version it holds and,
// when it refused a version, that version and why. A refused new item has no Version.
type HeldItem struct {
	Key            string        `json:"key"`
	Version        string        `json:"version,omitempty"`
	RefusedVersion string        `json:"refused_version,omitempty"`
	Reason         RefusalReason `json:"reason,omitempty"`
}

// KindReport is everything a gateway holds of one kind.
type KindReport struct {
	Kind  ConfigKind `json:"kind"`
	Items []HeldItem `json:"items"`
}

// MarshalJSON writes an empty Items as [].
func (k KindReport) MarshalJSON() ([]byte, error) {
	type plain KindReport
	if k.Items == nil {
		k.Items = []HeldItem{}
	}
	return json.Marshal(plain(k))
}

// ConfigReport is a full snapshot of what a gateway holds, one entry per kind it
// supports (empty when it holds nothing of it). Sent upward wrapped in an
// Envelope of KindConfigReport.
type ConfigReport struct {
	Kinds []KindReport `json:"kinds"`
}

// MarshalJSON writes an empty Kinds as [].
func (r ConfigReport) MarshalJSON() ([]byte, error) {
	type plain ConfigReport
	if r.Kinds == nil {
		r.Kinds = []KindReport{}
	}
	return json.Marshal(plain(r))
}
