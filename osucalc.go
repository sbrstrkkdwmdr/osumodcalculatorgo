package osumodcalculatorgo

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
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
	Range100 := 139 - (od * 8) + 0.5
	Range50 := 199 - (od * 8) + 0.5

	rangeobj.Range300 = range300
	rangeobj.Range100 = Range100
	rangeobj.Range50 = Range50
	rangeobj.OD = od

	return rangeobj

	//javascript version
	/*
			function ODtoms(od) {
		    let range300 = 79 - (od * 6) + 0.5
		    let Range100 = 139 - (od * 8) + 0.5
		    let Range50 = 199 - (od * 10) + 0.5

		    let rangeobj = {
		        range300: range300,
		        Range100: Range100,
		        Range50: Range50,
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
func MsToOD(hitwin300 float32, hitwin100 float32, hitwin50 float32) *hitwindowobj {
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

	odobj.Range300 = hitwin300
	//odobj.Range100 = hitwin100
	//odobj.Range50 = hitwin50
	odobj.OD = od

	return odobj
}

//convert ms to approach rate
func MsToAR(ms int) int {
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
func ODtoDT(od float32) *hitwindowobj {

	hitwins := new(hitwindowobj)

	hitwins.Range300 = (79 - (od * 6) + 0.5) * 2 / 3
	hitwins.Range100 = (139 - (od * 8) + 0.5) * 2 / 3
	hitwins.Range50 = (199 - (od * 8) + 0.5) * 2 / 3
	hitwins.OD = (79.5 - (od * 4 / 3)) / 6

	return hitwins
}

//convert overall difficulty to half time
func ODtoHT(od float32) *hitwindowobj {
	hitwins := new(hitwindowobj)

	hitwins.Range300 = (79 - (od * 6) + 0.5) * 4 / 3
	hitwins.Range100 = (139 - (od * 8) + 0.5) * 4 / 3
	hitwins.Range50 = (199 - (od * 8) + 0.5) * 4 / 3
	hitwins.OD = (79.5 - (od*2/3)/6)
	return hitwins
}

//calculate accuracy and grade/rank in osu! standard. accuracy is in decimal
func CalcGradeSTD(hit300 int, hit100 int, hit50 int, miss int) *accuracygrade {

	grades := new(accuracygrade)
	tophalf := float32((hit300 * 300) + (hit100 * 100) + (hit50 * 50))
	bottomhalf := float32(300 * (hit300 + hit100 + hit50 + miss))

	equationfull := float64(tophalf / bottomhalf)
	equationshort := float64(tophalf / bottomhalf)

	totalhits := float32(hit300 + hit100 + hit50 + miss)
	hit300fl := float32(hit300)
	hit50fl := float32(hit50)

	grades.Grade = "D"
	if float32(hit300fl/totalhits) > 0.6 && miss == 0 || float32(hit300fl/totalhits) > 0.7 {
		grades.Grade = "C"
	}
	if float32(hit300fl/totalhits) > 0.7 && miss == 0 || float32(hit300fl/totalhits) > 0.8 {
		grades.Grade = "B"
	}
	if float32(hit300fl/totalhits) > 0.8 && miss == 0 || float32(hit300fl/totalhits) > 0.9 {
		grades.Grade = "A"
	}
	if float32(hit300fl/totalhits) > 0.9 && miss == 0 && float32(hit50fl/totalhits) < 0.01 {
		grades.Grade = "S"
	}
	if float32(hit300fl/totalhits) == 1 {
		grades.Grade = "SS"
	}

	formattedacc, err := strconv.ParseFloat(strconv.FormatFloat(equationshort, 'f', -1, 32), 32)
	grades.Accuracy = float32(formattedacc)
	grades.Fullacc = float64(equationfull)

	if err != nil {
		fmt.Println(err.Error())
	}

	return grades
}

//calculates accuracy and grade for osu! taiko. hits: great(100%), good(50%), miss(0%)
func CalcGradeTaiko(hit300 int, hit100 int, miss int) *accuracygrade {
	grades := new(accuracygrade)

	tophalf := float32(float32(hit300) + (float32(hit100) / 2))
	bottomhalf := float32(hit300 + hit100 + miss)
	grades.Accuracy = float32(math.Round(float64(tophalf/bottomhalf)*100) / 100)
	grades.Fullacc = float64(tophalf / bottomhalf)

	grades.Grade = "D"

	if float32(tophalf/bottomhalf) > 0.8 {
		grades.Grade = "B"
	}
	if float32(tophalf/bottomhalf) > 0.9 {
		grades.Grade = "A"
	}
	if float32(tophalf/bottomhalf) > 0.95 {
		grades.Grade = "S"
	}
	if float32(tophalf/bottomhalf) == 1 {
		grades.Grade = "SS"
	}

	return grades
}

//calculates accuracy and grade for osu! catch the beat / fruits. hits: fruits, drops, droplets, miss
func CalcGradeCatch(hit300 int, hit100 int, hit50 int, hitkatu int, miss int) *accuracygrade {
	totalhits := hit300 + hit100 + hit50 + hitkatu

	tophalf := hit300 + hit100 + hit50
	bottomhalf := totalhits

	grades := new(accuracygrade)

	grades.Accuracy = float32(math.Round(float64(tophalf/bottomhalf)*100) / 100)
	grades.Fullacc = float64(tophalf / bottomhalf)

	grades.Grade = "D"
	if float32(tophalf/bottomhalf) > 0.85 {
		grades.Grade = "C"
	}
	if float32(tophalf/bottomhalf) > 0.9 {
		grades.Grade = "B"
	}
	if float32(tophalf/bottomhalf) > 0.94 {
		grades.Grade = "A"
	}
	if float32(tophalf/bottomhalf) > 0.98 {
		grades.Grade = "S"
	}
	if float32(tophalf/bottomhalf) == 1 {
		grades.Grade = "SS"
	}

	return grades
}

//calculates accuracy for osu! mania. hits: 300+/max, 300, 200, 100, 50, miss
func CalcGradeMania(hitgeki int, hit300 int, hitkatu int, hit100 int, hit50 int, miss int) *accuracygrade {
	grades := new(accuracygrade)

	totalhits := hitgeki + hit300 + hitkatu + hit100 + hit50 + miss

	tophalf := float32((300 * (hitgeki + hit300)) + (200 * hitkatu) + (100 * hit100) + (50 * hit50))
	bottomhalf := float32(300 * totalhits)

	grades.Accuracy = float32(math.Round(float64(tophalf/bottomhalf)*100) / 100)
	grades.Fullacc = float64(tophalf / bottomhalf)

	grades.Grade = "D"
	if float32(tophalf/bottomhalf) > 0.7 {
		grades.Grade = "C"
	}
	if float32(tophalf/bottomhalf) > 0.8 {
		grades.Grade = "B"
	}
	if float32(tophalf/bottomhalf) > 0.9 {
		grades.Grade = "A"
	}
	if float32(tophalf/bottomhalf) > 0.95 {
		grades.Grade = "S"
	}
	if float32(tophalf/bottomhalf) == 1 {
		grades.Grade = "SS"
	}

	return grades
}
func ToHR(cs float32, ar float32, od float32, hp float32) *basicmapval {
	values := new(basicmapval)

	csn := cs * 1.3
	arn := ar * 1.4
	hpn := hp * 1.4
	odn := od * 1.4

	if csn > 10 {
		values.CS = 10
	} else {
		values.CS = csn
	}
	if arn > 10 {
		values.AR = 10
	} else {
		values.AR = arn
	}
	if hpn > 10 {
		values.HP = 10
	} else {
		values.HP = hpn
	}
	if odn > 10 {
		values.OD = 10
	} else {
		values.OD = odn
	}
	return values
}
func ToEZ(cs float32, ar float32, hp float32, od float32) *basicmapval {
	values := new(basicmapval)

	csn := cs / 2
	arn := ar / 2
	hpn := hp / 2
	odn := od / 2

	if csn > 10 {
		values.CS = 10
	} else {
		values.CS = csn
	}
	if arn > 10 {
		values.AR = 10
	} else {
		values.AR = arn
	}
	if hpn > 10 {
		values.HP = 10
	} else {
		values.HP = hpn
	}
	if odn > 10 {
		values.OD = 10
	} else {
		values.OD = odn
	}
	return values
}

//takes the mods as a string and returns the total int value. useful for api v1.0
func ModStringToInt(mods string) int {
	var modint int
	if strings.Contains(mods, "NF") {
		modint += 1
	}
	if strings.Contains(mods, "EZ") {
		modint += 2
	}
	if strings.Contains(mods, "TD") {
		modint += 4
	}
	if strings.Contains(mods, "HD") {
		modint += 8
	}
	if strings.Contains(mods, "HR") {
		modint += 16
	}
	if strings.Contains(mods, "SD") {
		modint += 32
	}
	if strings.Contains(mods, "DT") {
		modint += 64
	}
	if strings.Contains(mods, "RX") || strings.Contains(mods, "RL") {
		modint += 128
	}
	if strings.Contains(mods, "HT") {
		modint += 256
	}
	if strings.Contains(mods, "NC") {
		modint += 512
	}
	if strings.Contains(mods, "FL") {
		modint += 1024
	}
	if strings.Contains(mods, "AT") {
		modint += 2048
	}
	if strings.Contains(mods, "SO") {
		modint += 4096
	}
	if strings.Contains(mods, "AP") {
		modint += 8192
	}
	if strings.Contains(mods, "PF") {
		modint += 16384
	}
	if strings.Contains(mods, "1K") {
		modint += 67108864
	}
	if strings.Contains(mods, "2K") {
		modint += 268435456
	}
	if strings.Contains(mods, "3K") {
		modint += 134217728
	}
	if strings.Contains(mods, "4K") {
		modint += 32768
	}	
	if strings.Contains(mods, "5K") {
		modint += 65536
	}
	if strings.Contains(mods, "6K") {
		modint += 131072
	}
	if strings.Contains(mods, "7K") {
		modint += 262144
	}
	if strings.Contains(mods, "8K") {
		modint += 524288
	}
	if strings.Contains(mods, "9K") {
		modint += 16777216
	}	
	if strings.Contains(mods, "FI") {
		modint += 1048576
	}
	if strings.Contains(mods, "RDM") {
		modint += 2097152
	}
	if strings.Contains(mods, "CN") {
		modint += 4194304
	}
	if strings.Contains(mods, "TP") {
		modint += 8388608
	}
	if strings.Contains(mods, "KC") {
		modint += 33554432
	}
	if strings.Contains(mods, "SV2") || strings.Contains(mods, "S2") {
		modint += 536870912
	}
	if strings.Contains(mods, "MR") {
		modint += 1073741824
	}

    if strings.Contains(mods, "NC") && !strings.Contains(mods, "DT") {
        modint += 64
    }
	return modint
}
//make a func that takes the mod int and returns the string\
func ModIntToString(modint int) string {
	var mods string
	if modint&1 == 1 {
		mods += "NF"
	}
	if modint&2 == 2 {
		mods += "EZ"
	}
	if modint&4 == 4 {
		mods += "TD"
	}
	if modint&8 == 8 {
		mods += "HD"
	}
	if modint&16 == 16 {
		mods += "HR"
	}
	if modint&32 == 32 {
		mods += "SD"
	}
	if modint&64 == 64 {
		mods += "DT"
	}
	if modint&128 == 128 {
		mods += "RX"
	}
	if modint&256 == 256 {
		mods += "HT"
	}
	if modint&512 == 512 {
		mods += "NC"
	}
	if modint&1024 == 1024 {
		mods += "FL"
	}
	if modint&2048 == 2048 {

		mods += "AT"
	}
	if modint&4096 == 4096 {
		mods += "SO"
	}
	if modint&8192 == 8192 {
		mods += "AP"
	}
	if modint&16384 == 16384 {
		mods += "PF"
	}
	if modint&67108864 == 67108864 {
		mods += "1K"
	}
	if modint&268435456 == 268435456 {
		mods += "2K"
	}
	if modint&134217728 == 134217728 {
		mods += "3K"
	}
	if modint&32768 == 32768 {
		mods += "4K"
	}
	if modint&65536 == 65536 {
		mods += "5K"
	}
	if modint&131072 == 131072 {
		mods += "6K"
	}
	if modint&262144 == 262144 {
		mods += "7K"
	}
	if modint&524288 == 524288 {
		mods += "8K"
	}
	if modint&16777216 == 16777216 {
		mods += "9K"
	}
	if modint&1048576 == 1048576 {
		mods += "FI"
	}
	if modint&2097152 == 2097152 {
		mods += "RDM"
	}
	if modint&4194304 == 4194304 {
		mods += "CN"
	}
	if modint&8388608 == 8388608 {
		mods += "TP"
	}
	if modint&33554432 == 33554432 {
		mods += "KC"
	}
	if modint&536870912 == 536870912 {
		mods += "SV2"
	}
	if modint&1073741824 == 1073741824 {
		mods += "MR"
	}
	return mods
}

func OrderMods(mods string) string {
	var ModsOrder []string = []string{"AT", "RX", "AP", "TP", "SO", "EZ", "HD", "HT", "DT", "NC", "HR", "SD", "PF", "FL", "NF"}
	var orderedMods string
	for _, mod := range ModsOrder {
		if strings.Contains(mods, mod) {
			orderedMods += mod
		}
	}
	return orderedMods
}