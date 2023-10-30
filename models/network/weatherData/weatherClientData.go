package weatherData

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TimezoneID     string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type Day struct {
	MaxTempC      float64   `json:"maxtemp_c"`
	MaxTempF      float64   `json:"maxtemp_f"`
	MinTempC      float64   `json:"mintemp_c"`
	MinTempF      float64   `json:"mintemp_f"`
	AvgTempC      float64   `json:"avgtemp_c"`
	AvgTempF      float64   `json:"avgtemp_f"`
	MaxWindMPH    float64   `json:"maxwind_mph"`
	MaxWindKPH    float64   `json:"maxwind_kph"`
	TotalPrecipMM float64   `json:"totalprecip_mm"`
	TotalPrecipIN float64   `json:"totalprecip_in"`
	AvgVisKM      float64   `json:"avgvis_km"`
	AvgVisMiles   float64   `json:"avgvis_miles"`
	AvgHumidity   float64   `json:"avghumidity"`
	Condition     Condition `json:"condition"`
	UV            float64   `json:"uv"`
}

type Astro struct {
	Sunrise          string `json:"sunrise"`
	Sunset           string `json:"sunset"`
	Moonrise         string `json:"moonrise"`
	Moonset          string `json:"moonset"`
	MoonPhase        string `json:"moon_phase"`
	MoonIllumination int64  `json:"moon_illumination"`
}

type Hour struct {
	TimeEpoch    int64     `json:"time_epoch"`
	Time         string    `json:"time"`
	TempC        float64   `json:"temp_c"`
	TempF        float64   `json:"temp_f"`
	IsDay        int       `json:"is_day"`
	Condition    Condition `json:"condition"`
	WindMPH      float64   `json:"wind_mph"`
	WindKPH      float64   `json:"wind_kph"`
	WindDegree   int       `json:"wind_degree"`
	WindDir      string    `json:"wind_dir"`
	PressureMB   float64   `json:"pressure_mb"`
	PressureIN   float64   `json:"pressure_in"`
	PrecipMM     float64   `json:"precip_mm"`
	PrecipIN     float64   `json:"precip_in"`
	Humidity     int       `json:"humidity"`
	Cloud        int       `json:"cloud"`
	FeelslikeC   float64   `json:"feelslike_c"`
	FeelslikeF   float64   `json:"feelslike_f"`
	WindchillC   float64   `json:"windchill_c"`
	WindchillF   float64   `json:"windchill_f"`
	HeatindexC   float64   `json:"heatindex_c"`
	HeatindexF   float64   `json:"heatindex_f"`
	DewpointC    float64   `json:"dewpoint_c"`
	DewpointF    float64   `json:"dewpoint_f"`
	WillItRain   int       `json:"will_it_rain"`
	ChanceOfRain int       `json:"chance_of_rain"`
	WillItSnow   int       `json:"will_it_snow"`
	ChanceOfSnow int       `json:"chance_of_snow"`
	VisKM        float64   `json:"vis_km"`
	VisMiles     float64   `json:"vis_miles"`
	GustMPH      float64   `json:"gust_mph"`
	GustKPH      float64   `json:"gust_kph"`
}

type Forecastday struct {
	Date       string `json:"date"`
	DateEpoch  int64  `json:"date_epoch"`
	Day        Day    `json:"day"`
	Astro      Astro  `json:"astro"`
	HourlyData []Hour `json:"hour"`
}

type Forecast struct {
	Forecastday []Forecastday `json:"forecastday"`
}

type ForeCastData struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

type RealTimeData struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int64   `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int64   `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		UV         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}

type TimeZoneData struct {
	Location Location `json:"location"`
}

type Current struct {
	LastUpdatedEpoch int64   `json:"last_updated_epoch"`
	LastUpdated      string  `json:"last_updated"`
	TempC            float64 `json:"temp_c"`
	TempF            float64 `json:"temp_f"`
	IsDay            int     `json:"is_day"`
	Condition        struct {
		Text string `json:"text"`
		Icon string `json:"icon"`
		Code int    `json:"code"`
	} `json:"condition"`
	WindMph    float64 `json:"wind_mph"`
	WindKph    float64 `json:"wind_kph"`
	WindDegree int     `json:"wind_degree"`
	WindDir    string  `json:"wind_dir"`
	PressureMb float64 `json:"pressure_mb"`
	PressureIn float64 `json:"pressure_in"`
	PrecipMm   float64 `json:"precip_mm"`
	PrecipIn   float64 `json:"precip_in"`
	Humidity   int     `json:"humidity"`
	Cloud      int     `json:"cloud"`
	FeelslikeC float64 `json:"feelslike_c"`
	FeelslikeF float64 `json:"feelslike_f"`
	VisKm      float64 `json:"vis_km"`
	VisMiles   float64 `json:"vis_miles"`
	UV         float64 `json:"uv"`
	GustMph    float64 `json:"gust_mph"`
	GustKph    float64 `json:"gust_kph"`
}

type FutureData struct {
	Location Location `json:"location"`
	Forecast Forecast `json:"forecast"`
}
