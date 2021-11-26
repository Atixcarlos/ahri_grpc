package main

import (
	"time"
    "database/sql/driver"
	"database/sql"
    "encoding/json"
)

// location represents the location info for a project.
type location struct {
    Latitude        *float64 `json:"latitude"`
	Longitude       *float64 `json:"longitude"`
	Elevation       *int     `json:"elevation"`
}

// Type OutdoorDC represents ACCA outdoor design conditions.
// outdoorDC represents the type of Oudoor design condition for a project.
type outdoorDC struct {
	WeatherStation string  `json:"weatherStation"`
	State          string  `json:"state"`
	Elevation      int     `json:"elevation"`   // Applies to weather station.
	Latitude       float64 `json:"latitude"`    // Applies to weather station.
	Heating99DB    *int    `json:"heating99DB"` // Outside design temp (dry bulb) for heating.
	Cooling01DB    *int    `json:"cooling01DB"` // Outside design temp (dry bulb) for cooling.
	CoincidentWB   *int    `json:"coincidentWB"`
	DG45RH         int     `json:"DG45RH"`
	DG50RH         int     `json:"DG50RH"`
	DG55RH         int     `json:"DG55RH"`
    DailyRange     string  `json:"dailyRange"`
}

// Type indoorDC represents indoor design conditions for a conditioned space.
type indoorDC struct {
    WinterIndoorF int      `json:"winterIndoorF"`
	SummerIndoorF int      `json:"summerIndoorF"`
	CoolingRH     int      `json:"coolingRH"`
}

// if MJ8 data does not exist
type nominalSize struct {
    NominalTons   float64   `json:"nominalTons"`
    HeatingBTUH   float64       `json:"heatingBTUH"`  
}

type loadCalculation struct {
    SensibleBTUH    float64     `json:"sensibleBTUH"`
    HeatingBTUH     float64     `json:"heatingBTUH"`
    LatentBTUH      float64     `json:"latentBTUH"`
}

type HeatedCooled struct {
    ProvidesCooling     bool        `json:"providesCooling"`
    ProvidesHeating     bool        `json:"providesHeating"`        
}

type systemAttributes struct {
    HeatedCooled                *HeatedCooled  `json:"heatedCooled"`
    FuelSource                  string        `json:"fuelSource"`
    EnergyDistributionMethod    string        `json:"energyDistributionMethod"`
}

type Search struct {
	DateCreated         time.Time           `json:"dateCreated"` // Read-only
	LastUpdated         time.Time           `json:"lastUpdated"` // Read-only
	ID                  int                 `json:"id"` 
	APICustomer         string              `json:"-"`
    Location            *location           `json:"location"` 
    OutdoorDC           *outdoorDC          `json:"outdoorDesignConditions"` 
    IndoorDC            *indoorDC           `json:"indoorDesignConditions"`
    NominalSize         *nominalSize        `json:"nominalSize,omitempty"`
    LoadCalculation     *loadCalculation    `json:"loadCalculation,omitempty"`
    SystemAttributes    *systemAttributes   `json:"systemAttributes"`
}


// Load
func Load(id int, myFilters map[string][]string, db *sql.DB) (*Search, error) {

	s  := new(Search)

    filterBytes, err := json.Marshal(myFilters)
	if err != nil {
		return nil, err
	}

	err = db.QueryRow("SELECT * FROM ahri_matchups.searches WHERE id = $1 and $2::json is not null", id, string(filterBytes)).
		Scan(&s.DateCreated, &s.LastUpdated, &s.ID, &s.APICustomer, &s.Location, &s.OutdoorDC, &s.IndoorDC, &s.NominalSize, &s.LoadCalculation, &s.SystemAttributes)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Save
func (s *Search) Save(db *sql.DB) error {
	err := db.QueryRow("SELECT * FROM ahri_matchups.search_save($1, $2, $3, $4, $5, $6, $7, $8)", s.ID, s.APICustomer, s.Location, s.OutdoorDC, s.IndoorDC, s.NominalSize,s.LoadCalculation, s.SystemAttributes).Scan(&s.DateCreated, &s.LastUpdated, &s.ID, &s.APICustomer, &s.Location, &s.OutdoorDC, &s.IndoorDC, &s.NominalSize, &s.LoadCalculation, &s.SystemAttributes)

	if err != nil {
		return err
	}

	return nil
}

// Download from SQL database.
// Scan implements the sql.Scanner interface.
func (l *location) Scan(value interface{}) error {
    var err error
    if value == nil {
        l = nil
    } else {
        data := value.([]byte)
        err = json.Unmarshal(data, l)
    }
    return err
}


// Insert in SQL database.
// Value implements the driver.Valuer interface.
func (l *location) Value() (driver.Value, error) {
    if l == nil {
        return nil, nil
    }
    return json.Marshal(l)
}

func (o *outdoorDC) Scan(value interface{}) error {
    var err error
    if value == nil {
        o = nil
    } else {
        data := value.([]byte)
        err = json.Unmarshal(data, o)
    }
    return err
}

func (o *outdoorDC) Value() (driver.Value, error) {
    if o == nil {
        return nil, nil
    }
    return json.Marshal(o)
}
func (i *indoorDC) Scan(value interface{}) error {
    var err error
    if value == nil {
        i = nil
    } else {
        data := value.([]byte)
        err = json.Unmarshal(data, i)
    }
    return err
}

func (i *indoorDC) Value() (driver.Value, error) {
    if i == nil {
        return nil, nil
    }
    return json.Marshal(i)
}

func (n *nominalSize) Scan(value interface{}) error {
    var err error
    if value == nil {
        n = nil
    } else {
        data := value.([]byte)
        err = json.Unmarshal(data, n)
    }
    return err
}

func (n *nominalSize) Value() (driver.Value, error) {
    if n == nil {
        return nil, nil
    }
    return json.Marshal(n)
}

func (lc *loadCalculation) Scan(value interface{}) error {
    var err error
    if value == nil {
        lc = nil
    } else {
        data := value.([]byte)
        err = json.Unmarshal(data, lc)
    }
    return err
}

func (lc *loadCalculation) Value() (driver.Value, error) {
    if lc == nil {
        return nil, nil
    }
    return json.Marshal(lc)
}

func (s *systemAttributes) Scan(value interface{}) error {
    var err error
    if value == nil {
        s = nil
    } else {
        data := value.([]byte)
        err = json.Unmarshal(data, s)
    }
    return err
}

func (s *systemAttributes) Value() (driver.Value, error) {
    if s == nil {
        return nil, nil
    }
    return json.Marshal(s)
}

func (h *HeatedCooled) Scan(value interface{}) error {
    var err error
    if value == nil {
        h = nil
    } else {
        data := value.([]byte)
        err = json.Unmarshal(data, h)
    }
    return err
}

func (h *HeatedCooled) Value() (driver.Value, error) {
    if h == nil {
        return nil, nil
    }
    return json.Marshal(h)
}

