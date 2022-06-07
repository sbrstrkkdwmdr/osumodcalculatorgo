package osumodcalculatorgo

//object for hit300,100,50 windows
type hitwindowobj struct {
	range300 float32
	range100 float32
	range50  float32
	od       float32
}

type accuracygrade struct {
	grade    string
	accuracy float32
	fullacc  float64
}

// circle size, approach rate, overall difficulty, health drain
type basicmapval struct {
	cs float32
	ar float32
	od float32
	hp float32
}
