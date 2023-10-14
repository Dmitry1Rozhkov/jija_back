package domain

type (
	Atm struct {
		Address   string     `json:"address"`
		Latitude  float64    `json:"latitude"`
		Longitude float64    `json:"longitude"`
		AllDay    bool       `json:"allDay"`
		Services  Services   `json:"services"`
		Load      []LoadItem `json:"load"`
	}

	Services struct {
		Wheelchair        ServiceInfo `json:"wheelchair"`
		Blind             ServiceInfo `json:"blind"`
		NfcForBankCards   ServiceInfo `json:"nfcForBankCards"`
		QrRead            ServiceInfo `json:"qrRead"`
		SupportsUsd       ServiceInfo `json:"supportsUsd"`
		SupportsChargeRub ServiceInfo `json:"supportsChargeRub"`
		SupportsEur       ServiceInfo `json:"supportsEur"`
		SupportsRub       ServiceInfo `json:"supportsRub"`
	}

	ServiceInfo struct {
		ServiceCapability string `json:"serviceCapability"`
		ServiceActivity   string `json:"serviceActivity"`
	}

	LoadItem struct {
		Day     int      `json:"day"`
		Loads   [][2]int `json:"loads"`
		WorkHrs []int    `json:"workHrs"`
	}
)

func (a *Atm) GetX() float64 {
	return a.Latitude
}

func (a *Atm) GetY() float64 {
	return a.Longitude
}
