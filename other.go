package osumodcalculatorgo_STRUCTS

//object for hit300,100,50 windows
type hitwindowobj struct {
	Range300 float32
	Range100 float32
	Range50  float32
	OD       float32
}

type accuracygrade struct {
	Grade    string
	Accuracy float32
	Fullacc  float64
}

// circle size, approach rate, overall difficulty, health drain
type basicmapval struct {
	CS float32
	AR float32
	OD float32
	HP float32
}
