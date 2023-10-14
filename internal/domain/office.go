package domain

type (
	Office struct {
		SalePointName       string      `json:"salePointName"`
		Address             string      `json:"address"`
		Status              string      `json:"status"`
		OpenHours           []OpenHours `json:"openHours"`
		RKO                 string      `json:"rko"`
		OpenHoursIndividual []OpenHours `json:"openHoursIndividual"`
		OfficeType          string      `json:"officeType"`
		SalePointFormat     string      `json:"salePointFormat"`
		SUOAvailability     string      `json:"suoAvailability"`
		HasRamp             string      `json:"hasRamp"`
		Latitude            float64     `json:"latitude"`
		Longitude           float64     `json:"longitude"`
		MetroStation        interface{} `json:"metroStation"`
		Distance            int         `json:"distance"`
		Kep                 bool        `json:"kep"`
		MyBranch            bool        `json:"myBranch"`
		Load                []LoadItem  `json:"load"`
	}

	OpenHours struct {
		Days  string `json:"days"`
		Hours string `json:"hours"`
	}
)

func (o *Office) GetX() float64 {
	return o.Latitude
}

func (o *Office) GetY() float64 {
	return o.Longitude
}
