package models

import (
	"fmt"
	"time"
)

type SearchPoint struct {
	Lat    string
	Long   string
	Radius string
}

func (s SearchPoint) ToString() string {
	return fmt.Sprintf("%s,%s,%s", s.Lat, s.Long, s.Radius)
}

type VelibStation struct {
	Name            string `json:"station_name"`
	Code            string `json:"station_code"`
	State           string `json:"station_state"`
	Dist            string `json:"dist"`
	Maxbikeoverflow int    `json:"maxbikeoverflow"`
	Densitylevel    string `json:"densitylevel"`
	Nbfreedock      int    `json:"nbfreedock"`
}

type VelibStatus struct {
	Nhits      int `json:"nhits"`
	Parameters struct {
		Dataset           string   `json:"dataset"`
		Timezone          string   `json:"timezone"`
		Rows              int      `json:"rows"`
		Format            string   `json:"format"`
		GeofilterDistance []string `json:"geofilter.distance"`
		Facet             []string `json:"facet"`
	} `json:"parameters"`
	Records []struct {
		Datasetid string `json:"datasetid"`
		Recordid  string `json:"recordid"`
		Fields    struct {
			StationState       string    `json:"station_state"`
			Maxbikeoverflow    int       `json:"maxbikeoverflow"`
			Densitylevel       string    `json:"densitylevel"`
			Nbbikeoverflow     int       `json:"nbbikeoverflow"`
			Dist               string    `json:"dist"`
			Nbedock            int       `json:"nbedock"`
			StationName        string    `json:"station_name"`
			Kioskstate         string    `json:"kioskstate"`
			Nbfreeedock        int       `json:"nbfreeedock"`
			StationType        string    `json:"station_type"`
			StationCode        string    `json:"station_code"`
			Creditcard         string    `json:"creditcard"`
			Nbfreedock         int       `json:"nbfreedock"`
			Duedate            string    `json:"duedate"`
			Nbebikeoverflow    int       `json:"nbebikeoverflow"`
			Nbebike            int       `json:"nbebike"`
			Overflow           string    `json:"overflow"`
			Nbdock             int       `json:"nbdock"`
			Geo                []float64 `json:"geo"`
			Overflowactivation string    `json:"overflowactivation"`
			Nbbike             int       `json:"nbbike"`
		} `json:"fields"`
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		RecordTimestamp time.Time `json:"record_timestamp"`
	} `json:"records"`
	FacetGroups []struct {
		Facets []struct {
			Count int    `json:"count"`
			Path  string `json:"path"`
			State string `json:"state"`
			Name  string `json:"name"`
		} `json:"facets"`
		Name string `json:"name"`
	} `json:"facet_groups"`
}

// ExtractStations extract the station information from the velib status
func (status VelibStatus) ExtractStations() []VelibStation {
	var stations []VelibStation
	for _, s := range status.Records {
		newStation := VelibStation{
			Name:            s.Fields.StationName,
			Code:            s.Fields.StationCode,
			State:           s.Fields.StationState,
			Nbfreedock:      s.Fields.Nbfreeedock,
			Maxbikeoverflow: s.Fields.Maxbikeoverflow,
			Densitylevel:    s.Fields.Densitylevel,
			Dist:            s.Fields.Dist}
		stations = append(stations, newStation)
	}
	return stations
}
