package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gocolly/colly"
)

func main() {
	lambda.Start(Handler)
}

// Handler handle the lambda function
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod != http.MethodGet {
		return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 405}, nil
	}

	m := collect()

	response, err := json.Marshal(m)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
}

func collect() *MetData {
	c := colly.NewCollector()

	var m *MetData

	c.OnHTML("#table-meteorologico tbody tr", func(e *colly.HTMLElement) {
		t := make([]string, 0)
		e.ForEach("td", func(i int, el *colly.HTMLElement) {
			t = append(t, strings.TrimSpace(el.Text))
		})

		m = &MetData{t[0], t[1], toF(t[2]), toF(t[3]), toF(t[4]), toF(t[5]), toF(t[6]), t[7]}
	})

	c.Visit("http://alertablu.cob.sc.gov.br/d/")

	return m
}

// MetData meteorological data
type MetData struct {
	Station             string  `json:"station"`
	RawReadTime         string  `json:"rawReadTime"`
	Temperature         float64 `json:"temerature"`
	AirHumidity         float64 `json:"airHumidity"`
	ThermalSensation    float64 `json:"thermalSensation"`
	AtmosphericPressure float64 `json:"atmosphericPressure"`
	WindSpeed           float64 `json:"windSpeed"`
	WindDirection       string  `json:"windDirection"`
}

func toF(s string) float64 {
	s = strings.Replace(s, ",", ".", 1)
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
