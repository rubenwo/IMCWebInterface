package main

//Config is the upper struct containing all available settings for the IMC
type Config struct {
	ID               string           `json:"id"`
	DateTimeSettings DateTimeSettings `json:"datetime_settings"`
	NetworkSettings  NetworkSettings  `json:"network_settings"`
	AudioSettings    AudioSettings    `json:"audio_settings"`
	RadioSettings    RadioSettings    `json:"radio_settings"`
	AlarmSettings    AlarmSettings    `json:"alarm_settings"`
	GeneralSettings  GeneralSettings  `json:"general_settings`
}

//DateTimeSettings is a struct containing all settings for date & time
type DateTimeSettings struct {
	TimeHH   int `json:"time_hour"`
	TimeMM   int `json:"time_minute"`
	DateYY   int `json:"date_year"`
	DateMO   int `json:"date_month"`
	DateDD   int `json:"date_day"`
	TimeZone int `json:"timezone"`
	TimeSync int `json:"timesync"`
}

//NetworkSettings contains an integer array containing the mac address for the IMC
type NetworkSettings struct {
	MacAddress []int `json:"mac_address"`
}

//AudioSettings contains all settings contraining audio (Volume, Treble & Bass)
type AudioSettings struct {
	Volume int `json:"volume"`
	Treble int `json:"treble"`
	Bass   int `json:"bass"`
}

//RadioSettings contains the settings for a new radio_channel [WARNING] Might be removed
type RadioSettings struct {
	IpAddress string `json:"ip_address"`
	Port      int    `json:"port"`
	Route     string `json:"route"`
}

//AlarmSettings contains all settings for the alarms [WARNING] Might change
type AlarmSettings struct {
	TimeHH       int    `json:"time_hour"`
	TimeMM       int    `json:"time_minute"`
	ReArm        int    `json:"re_arm"`
	AudioType    int    `json:"audio_type"`
	RadioChannel string `json:"radio_channel"`
}

//GeneralSettings contains a flag whether the device should return to factory settings
type GeneralSettings struct {
	FactorySettings int `json:"factory_settings"`
}
