package api

type GetSetString struct {
	Value string
}

type GetSetFloat struct {
	Value float64
}

type GetSetBool struct {
	Value bool
}

type GetSetInt struct {
	Value int
}

type GetStatus struct {
	Mode               string
	CurrentSetpoint    float64
	BrewSetpoint       float64
	SteamSetpoint      float64
	CurrentPressure    float64
	CurrentTemperature float64
}
