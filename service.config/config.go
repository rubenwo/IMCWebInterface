package main

//Config is a struct used for decoding the configs.json file
type Config struct {
	ID              string          `json:"id"`
	RadioSettings   RadioSettings   `json:"radio_settings"`
	AlarmSettings   AlarmSettings   `json:"alarm_settings"`
	GeneralSettings GeneralSettings `json:"general_settings"`
}

//RadioSettings is a struct used for decoding the configs.json file
type RadioSettings struct {
	Volume int    `json:"volume"`
	Treble int    `json:"treble"`
	Bass   int    `json:"bass"`
	Sound  string `json:"sound"`
}

//AlarmSettings is a struct used for decoding the configs.json file
type AlarmSettings struct {
	AlarmID string `json:"alarm_id"`
	Time    string `json:"time"`
}

//GeneralSettings is a struct used for decoding the configs.json file
type GeneralSettings struct {
	Language string `json:"language"`
	Time     string `json:"time"`
}
