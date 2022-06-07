package osumodcalculatorgo

import (
	"math"
	"reflect"
)

/*
convert approach rate to milliseconds
*/
func ApproachMS(inputar float32) float32 {
	var ogms float32
	if inputar > 5 {
		ogms = 1200 - (((inputar - 5) * 10) * 15)
	} else {
		ogms = 1800 - (((inputar) * 10) * 12)
	}
	finalnumber := ogms
	return finalnumber
}

//approach rate if double time
func ApproachDT(inputar float32) float32 {
	millisex := ApproachMS(inputar) * 2 / 3
	var outputar float32
	if millisex < 300 {
		outputar = 11
	} else if millisex < 1200 {
		outputar = (((11 - (millisex-300)/150) * 100) / 100)
	} else {
		outputar = (((5 - (millisex-1200)/120) * 100) / 100)
	}
	return outputar
}

//convert AR to half time
func ApproachHT(inputar float32) float32 {
	millisex := ApproachMS(inputar) * 4 / 3
	var outputar float32
	if millisex < 300 {
		outputar = 11
	} else if millisex < 1200 {

		outputar = (((11 - (millisex-300)/150) * 100) / 100)
	} else {
		outputar = (((5 - (millisex-1200)/120) * 100) / 100)
	}
	return outputar
}

//convert OD to hit windows in ms
func ODtoms(od float32) *hitwindowobj {

	rangeobj := new(hitwindowobj)

	range300 := 79 - (od * 6) + 0.5
	range100 := 139 - (od * 8) + 0.5
	range50 := 199 - (od * 8) + 0.5

	rangeobj.range300 = range300
	rangeobj.range100 = range100
	rangeobj.range50 = range50
	rangeobj.od = od

	return rangeobj

	//javascript version
	/*
			function ODtoms(od) {
		    let range300 = 79 - (od * 6) + 0.5
		    let range100 = 139 - (od * 8) + 0.5
		    let range50 = 199 - (od * 10) + 0.5

		    let rangeobj = {
		        range300: range300,
		        range100: range100,
		        range50: range50,
		    }
		    return rangeobj;
		}
	*/
}

//convert approach rate to ms
func ARtoms(ar float32) float32 {
	var ms float32
	if ar > 5 {
		ms = 1200 - (((ar - 5) * 10) * 15)
	} else {
		ms = 1800 - ((ar * 10) * 12)
	}
	return ms
}

//convert hit windows to od
func msToOD(hitwin300 float32, hitwin100 float32, hitwin50 float32) *hitwindowobj {
	odobj := new(hitwindowobj)
	var od float32
	od = 0
	if reflect.TypeOf(hitwin300).Kind() == reflect.Float32 {
		od = ((79.5 - hitwin300) / 6)
	} /*else if reflect.TypeOf(hitwin100).Kind() == reflect.Float32 {
		od = ((139.5 - hitwin100) / 8)
	} else if reflect.TypeOf(hitwin50).Kind() == reflect.Float32 {
		od = ((199.5 - hitwin50) / 10)
	}*/

	odobj.range300 = hitwin300
	//odobj.range100 = hitwin100
	//odobj.range50 = hitwin50
	odobj.od = od

	return odobj
}

//convert ms to approach rate
func msToAR(ms int) int {
	ar := 0
	if ms < 300 {
		ar = 11
	} else if ms < 1200 {
		ar = ((11 - (ms-300)/150) * 100) / 100
	} else {
		ar = ((5 - (ms-1200)/120) * 100) / 100
	}
	return ar
}

//convert overall difficulty to double time
func odDT(od float32) *hitwindowobj {

	hitwins := new(hitwindowobj)

	hitwins.range300 = (79 - (od * 6) + 0.5) * 2 / 3
	hitwins.range100 = (139 - (od * 8) + 0.5) * 2 / 3
	hitwins.range50 = (199 - (od * 8) + 0.5) * 2 / 3
	hitwins.od = (79.5 - (od * 4 / 3)) / 6

	return hitwins
}

//convert overall difficulty to half time
func odHT(od float32) *hitwindowobj {
	hitwins := new(hitwindowobj)

	hitwins.range300 = (79 - (od * 6) + 0.5) * 4 / 3
	hitwins.range100 = (139 - (od * 8) + 0.5) * 4 / 3
	hitwins.range50 = (199 - (od * 8) + 0.5) * 4 / 3
	hitwins.od = (79.5 - (od*2/3)/6)
	return hitwins
}

//calculate accuracy and grade/rank in osu! standard. accuracy is in decimal
func calcGradeSTD(hit300 int, hit100 int, hit50 int, miss int) *accuracygrade {

	grades := new(accuracygrade)
	tophalf := ((hit300 * 300) + (hit100 * 100) + (hit50 * 50))
	bottomhalf := (300 * (hit300 + hit100 + hit50 + miss))

	equationfull := float32(tophalf / bottomhalf)
	equationshort := float64(tophalf / bottomhalf)

	totalhits := hit300 + hit100 + hit50 + miss

	grades.grade = "D"
	if float32(hit300/totalhits) > 0.6 && miss == 0 || float32(hit300/totalhits) > 0.7 {
		grades.grade = "C"
	}
	if float32(hit300/totalhits) > 0.7 && miss == 0 || float32(hit300/totalhits) > 0.8 {
		grades.grade = "B"
	}
	if float32(hit300/totalhits) > 0.8 && miss == 0 || float32(hit300/totalhits) > 0.9 {
		grades.grade = "A"
	}
	if float32(hit300/totalhits) > 0.9 && miss == 0 && float32(hit50/totalhits) < 0.01 {
		grades.grade = "S"
	}
	if float32(hit300/totalhits) > 1 {
		grades.grade = "SS"
	}

	grades.accuracy = float32(math.Round(float64(equationshort)*100) / 100)
	grades.fullacc = float64(equationfull)

	return grades
}

//calculates accuracy and grade for osu! taiko. hits: great(100%), good(50%), miss(0%)
func calcGradeTaiko(hit300 int, hit100 int, miss int) *accuracygrade {
	grades := new(accuracygrade)

	tophalf := (hit300 + (hit100 / 2))
	bottomhalf := (hit300 + hit100 + miss)
	grades.accuracy = float32(math.Round(float64(tophalf/bottomhalf)*100) / 100)
	grades.fullacc = float64(tophalf / bottomhalf)

	grades.grade = "D"

	if float32(tophalf/bottomhalf) > 0.8 {
		grades.grade = "B"
	}
	if float32(tophalf/bottomhalf) > 0.9 {
		grades.grade = "A"
	}
	if float32(tophalf/bottomhalf) > 0.95 {
		grades.grade = "S"
	}
	if float32(tophalf/bottomhalf) == 1 {
		grades.grade = "SS"
	}

	return grades
}

//calculates accuracy and grade for osu! catch the beat / fruits. hits: fruits, drops, droplets, miss
func calcGradeCatch(hit300 int, hit100 int, hit50 int, hitkatu int, miss int) *accuracygrade {
	totalhits := hit300 + hit100 + hit50 + hitkatu

	tophalf := hit300 + hit100 + hit50
	bottomhalf := totalhits

	grades := new(accuracygrade)

	grades.accuracy = float32(math.Round(float64(tophalf/bottomhalf)*100) / 100)
	grades.fullacc = float64(tophalf / bottomhalf)

	grades.grade = "D"
	if float32(tophalf/bottomhalf) > 0.85 {
		grades.grade = "C"
	}
	if float32(tophalf/bottomhalf) > 0.9 {
		grades.grade = "B"
	}
	if float32(tophalf/bottomhalf) > 0.94 {
		grades.grade = "A"
	}
	if float32(tophalf/bottomhalf) > 0.98 {
		grades.grade = "S"
	}
	if float32(tophalf/bottomhalf) == 1 {
		grades.grade = "SS"
	}

	return grades
}

//calculates accuracy for osu! mania. hits: 300+/max, 300, 200, 100, 50, miss
func calcGradeMania(hitgeki int, hit300 int, hitkatu int, hit100 int, hit50 int, miss int) *accuracygrade {
	grades := new(accuracygrade)

	totalhits := hitgeki + hit300 + hitkatu + hit100 + hit50 + miss

	tophalf := (300 * (hitgeki + hit300)) + (200 * hitkatu) + (100 * hit100) + (50 * hit50)
	bottomhalf := 300 * totalhits

	grades.accuracy = float32(math.Round(float64(tophalf/bottomhalf)*100) / 100)
	grades.fullacc = float64(tophalf / bottomhalf)

	grades.grade = "D"
	if float32(tophalf/bottomhalf) > 0.7 {
		grades.grade = "C"
	}
	if float32(tophalf/bottomhalf) > 0.8 {
		grades.grade = "B"
	}
	if float32(tophalf/bottomhalf) > 0.9 {
		grades.grade = "A"
	}
	if float32(tophalf/bottomhalf) > 0.95 {
		grades.grade = "S"
	}
	if float32(tophalf/bottomhalf) == 1 {
		grades.grade = "SS"
	}

	return grades
}
func toHR(cs float32, ar float32, od float32, hp float32) *basicmapval {
	values := new(basicmapval)

	csn := cs * 1.3
	arn := ar * 1.4
	hpn := hp * 1.4
	odn := od * 1.4

	if csn > 10 {
		values.cs = 10
	} else {
		values.cs = csn
	}
	if arn > 10 {
		values.ar = 10
	} else {
		values.ar = arn
	}
	if hpn > 10 {
		values.hp = 10
	} else {
		values.hp = hpn
	}
	if odn > 10 {
		values.od = 10
	} else {
		values.od = odn
	}
	return values
}
func toEZ(cs float32, ar float32, hp float32, od float32) *basicmapval {
	values := new(basicmapval)

	csn := cs / 2
	arn := ar / 2
	hpn := hp / 2
	odn := od / 2

	if csn > 10 {
		values.cs = 10
	} else {
		values.cs = csn
	}
	if arn > 10 {
		values.ar = 10
	} else {
		values.ar = arn
	}
	if hpn > 10 {
		values.hp = 10
	} else {
		values.hp = hpn
	}
	if odn > 10 {
		values.od = 10
	} else {
		values.od = odn
	}
	return values
}
