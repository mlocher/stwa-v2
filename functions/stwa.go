package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-lambda-go/lambda"
)

type stwaData struct {
	Sites []stwaSite `json:"aaData"`
}

type stwaSite struct {
	Name  string
	State string
}

func (o *stwaSite) UnmarshalJSON(data []byte) error {
	var v [2]string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	o.Name = v[0]
	switch v[1] {
	case "#EC1C24":
		o.State = "gale_warning"
	case "#FFDD15":
		o.State = "advance_warning"
	case "#8BC53F":
		o.State = "standby"
	case "#6c757d--change-me":
		o.State = "out_of_order"
	default:
		o.State = "unknown"
	}

	return nil
}

type gjFeatureCollection struct {
	Type     string      `json:"type"`
	Features []gjFeature `json:"features"`
}

type gjFeature struct {
	Type       string       `json:"type"`
	Properties gjProperties `json:"properties"`
	Geometry   gjGeometry   `json:"geometry"`
}

type gjProperties struct {
	Name        string `json:"name"`
	State       string `json:"state"`
	Description string `json:"description,omitempty"`
	Height      int    `json:"height,omitempty"`
}

type gjGeometry struct {
	Type        string     `json:"type"`
	Coordinates [2]float32 `json:"coordinates"`
}

func newFeature(name string, state string, coordinates [2]float32, description string, height int) gjFeature {
	f := gjFeature{
		Type:       "Feature",
		Properties: gjProperties{},
		Geometry: gjGeometry{
			Type: "Point",
		},
	}
	f.Properties.Name = name
	f.Properties.State = state
	f.Properties.Description = description
	f.Properties.Height = height
	f.Geometry.Coordinates = coordinates
	return f
}

func newFeatureCollection(features []gjFeature) gjFeatureCollection {
	fc := gjFeatureCollection{
		Type: "FeatureCollection",
	}
	fc.Features = features
	return fc
}

func stwa(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	type stwaMetaData struct {
		name        string
		description string
		height      int
		coordinates [2]float32
	}
	metaData := map[string]stwaMetaData{
		"Neusiedl": stwaMetaData{
			name:        "Neusiedl",
			description: "Mast auf der Mole im Südhafen",
			height:      6,
			coordinates: [2]float32{16.83833333, 47.92666667},
		},
		"Weiden": stwaMetaData{
			name:        "Weiden",
			description: "Mast Kanaleinfahrt, von See kommend links",
			height:      0,
			coordinates: [2]float32{16.85222222, 47.91666667},
		},
		"Podersdorf": stwaMetaData{
			name:        "Podersdorf",
			description: "Leuchtturm an der Außenmole",
			height:      0,
			coordinates: [2]float32{16.82638889, 47.86027778},
		},
		"Illmitz-Hoelle": stwaMetaData{
			name:        "Illmitz-Hölle",
			description: "Mast am Haus des Wasserleitungsverbandes",
			height:      7,
			coordinates: [2]float32{16.78305556, 47.80361111},
		},
		"Illmitz": stwaMetaData{
			name:        "Illmitz",
			description: "Mast in der Badeanlage - Gemeindehafen, südlich des Yachtclub Illmitz (YCI)",
			height:      0,
			coordinates: [2]float32{16.7425, 47.75583333},
		},
		"Fertoerakos": stwaMetaData{
			name:        "Fertörakos",
			description: "Mast am Haus des Wasserverbandes",
			height:      0,
			coordinates: [2]float32{16.69333333, 47.72027778},
		},
		"Moerbisch": stwaMetaData{
			name:        "Mörbisch",
			description: "Am Mast des Polizei- und Zollwachhauses",
			height:      8,
			coordinates: [2]float32{16.69777778, 47.75138889},
		},
		"Rust": stwaMetaData{
			name:        "Rust",
			description: "Turm mit Mast am Clubhaus des Burgenländischen Yachtclubs (BYC)",
			height:      10,
			coordinates: [2]float32{16.69944444, 47.80361111},
		},
		"Oggau": stwaMetaData{
			name:        "Oggau",
			description: "Mast bei Kanaleinfahrt",
			height:      20,
			coordinates: [2]float32{16.70111111, 47.845},
		},
		"Purbach": stwaMetaData{
			name:        "Purbach",
			description: "Mast am Fußwegende, Anfang Kanaleinfahrt",
			height:      5,
			coordinates: [2]float32{16.74611111, 47.90111111},
		},
		"Breitenbrunn": stwaMetaData{
			name:        "Breitenbrunn",
			description: "Flaggenmast im Yachtclub Breitenbrunn (YCBb)",
			height:      15,
			coordinates: [2]float32{16.77, 47.91611111},
		},
	}

	url := "https://www.lsz-b.at/fileadmin/stwa/scripts/holeStwaStatus.php"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var stwa stwaData
	json.Unmarshal(jsonData, &stwa)

	var sites []gjFeature
	for _, site := range stwa.Sites {
		if site.Name == "Neusiedler See" {
			continue
		}
		sites = append(sites, newFeature(metaData[site.Name].name, site.State, metaData[site.Name].coordinates, metaData[site.Name].description, metaData[site.Name].height))
	}

	geoJSON, err := json.Marshal(newFeatureCollection(sites))
	if err != nil {
		panic(err.Error())
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Cache-Control":               "public, s-maxage=60",
		},
		Body: string(geoJSON),
	}, nil
}

func main() {
	lambda.Start(stwa)
}
