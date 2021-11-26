package main

import (
	"database/sql"
)


type Result struct {
//	AHRIRefElectric             string              `json:"AHRIRefElectric"`
//    AHRIRefFurnace              string              `json:"AHRIRefFurnace"`
//    OutdoorUnit                 string              `json:"outdoorUnit"`
//    IndoorUnit                  string              `json:"indoorUnit"`
//    Furnace                     string              `json:"furnace"`
    CoolingCapacity             int                 `json:"coolingCapacity"`
    EER                         float64             `json:"eer"`
    SEER                        float64             `json:"seer"`
//    Phase                       int                 `json:"phase"`
//    HCAP47                      int                 `json:"hcap47"`
//    HCAP17                      int                 `json:"hcap17"`
    HSPF                        *float64             `json:"hspf"`
//    FurnaceConfigurations       *helpers.RawJSON    `json:"furnaceConfigurations"`
//    FurnaceInputBTUH            int                 `json:"furnaceInputBTUH"`
    AFUE                        float64             `json:"afue"`
    OutdoorUnitSKU              string              `json:"outdoorUnitSKU"`
    IndoorUnitSKU               string              `json:"indoorUnitSKU"`
    FurnaceSKU                  string              `json:"furnaceSKU"`

//    ProductLines                *helpers.RawJSON    `json:"productLines"`
//    FuelTypes                   *helpers.RawJSON    `json:"fuelTypes"`
//    HeatedCooled                *helpers.RawJSON    `json:"heatedCooled"`
//    EnergyDistributionMethod    string              `json:"energyDistributionMethod"`
//    MaxNrDuctlessZones          int                 `json:"maxNrDuctlessZones"`
//    CompressorStages            string              `json:"compressorStages"`
    NominalCoolingTons          *float64             `json:"nominalCoolingTons"`
//    NominalCoolingTonsMatch     bool                `json:"nominalCoolingTonsMatch"`
//    BrandedTiersMatch           bool                `json:"brandedTiersMatch"`
//    StackHeight                 float64             `json:"stackHeight"`         
   SensibleCapacity            *int                 `json:"sensibleCapacity"`
   LatentCapacity              *int                 `json:"latentCapacity"`

    // heatingCapacity
}

// New returns a new *Result variable.
func New() *Result {
	r := new(Result)
	return r
}

// Load with manual J
func LoadWithManualJ(my_heated_cooled *HeatedCooled, my_fuel_source string, my_energy_distribution_method string, my_evaporator_ewb *int,
    my_indoor_db *int, my_outdoor_db *int, my_mj8_sensible_cooling_btuh float64, my_mj8_latent_cooling_btuh float64, my_mj8_heating_btuh float64, myFilters map[string][]string, db *sql.DB) ([]Result, error) {
    myData := make([]Result, 0)


	rows, err := db.Query("SELECT * FROM oem_carrier.equipment_search_raw_with_manual_j($1,$2, $3, $4, $5, $6, $7, $8, $9)", my_heated_cooled, my_fuel_source, my_energy_distribution_method, my_evaporator_ewb, my_indoor_db, my_outdoor_db, int(my_mj8_sensible_cooling_btuh), int(my_mj8_latent_cooling_btuh), int(my_mj8_heating_btuh))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		r := New()

        err = rows.Scan(&r.CoolingCapacity, &r.EER, &r.SEER,  &r.HSPF, &r.AFUE, &r.OutdoorUnitSKU, &r.IndoorUnitSKU, &r.FurnaceSKU, &r.NominalCoolingTons, &r.SensibleCapacity, &r.LatentCapacity)
		if err != nil {
			return nil, err
		}
		myData = append(myData, *r)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	rows.Close()

	return myData, nil

}



// Load without manual J
func LoadWithoutManualJ(my_heated_cooled *HeatedCooled, my_fuel_source string, my_energy_distribution_method string, my_nominal_tons float64, my_heating_btuh float64, my_evaporator_ewb *int,
    my_indoor_db *int, my_outdoor_db *int, myFilters map[string][]string, db *sql.DB) ([]Result, error) {
    myData := make([]Result, 0)


	rows, err := db.Query("SELECT * FROM oem_carrier.equipment_search_raw_without_manual_j($1,$2, $3, $4, $5, $6, $7, $8)", my_heated_cooled, my_fuel_source, my_energy_distribution_method, my_nominal_tons, int(my_heating_btuh), my_evaporator_ewb, my_indoor_db, my_outdoor_db)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		r := New()
        err = rows.Scan(&r.CoolingCapacity, &r.EER, &r.SEER,  &r.HSPF, &r.AFUE, &r.OutdoorUnitSKU, &r.IndoorUnitSKU, &r.FurnaceSKU, &r.NominalCoolingTons, &r.SensibleCapacity, &r.LatentCapacity)
		if err != nil {
			return nil, err
		}
		myData = append(myData, *r)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	rows.Close()

	return myData, nil

}

