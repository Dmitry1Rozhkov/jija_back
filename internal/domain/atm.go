package domain

type (
	Atm struct {
		Address   string   `json:"address"`
		Latitude  float64  `json:"latitude"`
		Longitude float64  `json:"longitude"`
		AllDay    bool     `json:"allDay"`
		Services  Services `json:"services"`
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
)

/*
{
            "address": "ул. Богородский Вал, д. 6, корп. 1",
            "latitude": 55.802432,
            "longitude": 37.704547,
            "allDay": false,
            "services": {
                "wheelchair": {
                    "serviceCapability": "UNKNOWN",
                    "serviceActivity": "UNKNOWN"
                },
                "blind": {

                    "serviceCapability": "UNKNOWN",
                    "serviceActivity": "UNKNOWN"
                },
                "nfcForBankCards": {
                    "serviceCapability": "UNKNOWN",
                    "serviceActivity": "UNAVAILABLE"
                },
                "qrRead": {
                    "serviceCapability": "UNSUPPORTED",
                    "serviceActivity": "UNAVAILABLE"
                },
                "supportsUsd": {
                    "serviceCapability": "UNSUPPORTED",
                    "serviceActivity": "UNAVAILABLE"
                },
                "supportsChargeRub": {
                    "serviceCapability": "SUPPORTED",
                    "serviceActivity": "AVAILABLE"
                },
                "supportsEur": {
                    "serviceCapability": "UNSUPPORTED",
                    "serviceActivity": "UNAVAILABLE"
                },
                "supportsRub": {
                    "serviceCapability": "SUPPORTED",
                    "serviceActivity": "AVAILABLE"
                }
            }
        },


*/
