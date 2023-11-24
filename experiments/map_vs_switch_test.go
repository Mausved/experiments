package experiments

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkSome(b *testing.B) {
	id := 10
	b.Run("switch", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			switchGetCountryNameByID(id)
		}
		memStat := runtime.MemStats{}
		runtime.ReadMemStats(&memStat)
		fmt.Println(memStat.Alloc)
	})
	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mapGetCountryNameByID(id)
		}
		memStat := runtime.MemStats{}
		runtime.ReadMemStats(&memStat)
		fmt.Println(memStat.Alloc)
	})
}

const (
	BrazilID            = 11
	MexicoID            = 12
	IndiaID             = 14
	IndonesiaID         = 15
	ElSalvadorID        = 21
	ColombiaID          = 22
	GuatemalaID         = 23
	PeruID              = 24
	ChileID             = 25
	MoroccoID           = 30
	EcuadorID           = 43
	DominicanRepublicID = 46
	BoliviaID           = 71
	HondurasID          = 75
	NicaraguaID         = 76
	EgyptID             = 153
	SouthAfricaID       = 10
	PanamaID            = 77
	CostaRicaID         = 72
	NigeriaID           = 40
	KenyaID             = 8
	MalaysiaID          = 85
	PakistanID          = 60
	ArgentinaID         = 13
	KazakhstanID        = 2
	JamaicaID           = 54
	ThailandID          = 19
)

func switchGetCountryNameByID(id int) (string, bool) {
	countryName := ""

	switch id {
	case BrazilID:
		countryName = "Brazil"
	case MexicoID:
		countryName = "Mexico"
	case IndiaID:
		countryName = "India"
	case IndonesiaID:
		countryName = "Indonesia"
	case ElSalvadorID:
		countryName = "ElSalvador"
	case ColombiaID:
		countryName = "Colombia"
	case GuatemalaID:
		countryName = "Guatemala"
	case PeruID:
		countryName = "Peru"
	case ChileID:
		countryName = "Chile"
	case MoroccoID:
		countryName = "Morocco"
	case EcuadorID:
		countryName = "Ecuador"
	case DominicanRepublicID:
		countryName = "Dominican Republic"
	case BoliviaID:
		countryName = "Bolivia"
	case HondurasID:
		countryName = "Honduras"
	case NicaraguaID:
		countryName = "Nicaragua"
	case EgyptID:
		countryName = "Egypt"
	case SouthAfricaID:
		countryName = "South Africa"
	case PanamaID:
		countryName = "Panama"
	case CostaRicaID:
		countryName = "Costa-Rica"
	case NigeriaID:
		countryName = "Nigeria"
	case KenyaID:
		countryName = "Kenya"
	case MalaysiaID:
		countryName = "Malaysia"
	case PakistanID:
		countryName = "Pakistan"
	case ArgentinaID:
		countryName = "Argentina"
	case KazakhstanID:
		countryName = "Kazakhstan"
	case JamaicaID:
		countryName = "Jamaica"
	case ThailandID:
		countryName = "Thailand"
	}

	return countryName, countryName != ""
}

func mapGetCountryNameByID(id int) (string, bool) {
	var countryNameByID = map[int]string{
		BrazilID:            "Brazil",
		MexicoID:            "Mexico",
		IndiaID:             "India",
		IndonesiaID:         "Indonesia",
		ElSalvadorID:        "ElSalvador",
		ColombiaID:          "Colombia",
		GuatemalaID:         "Guatemala",
		PeruID:              "Peru",
		ChileID:             "Chile",
		MoroccoID:           "Morocco",
		EcuadorID:           "Ecuador",
		DominicanRepublicID: "DominicanRepublic",
		BoliviaID:           "Bolivia",
		HondurasID:          "Honduras",
		NicaraguaID:         "Nicaragua",
		EgyptID:             "Egypt",
		SouthAfricaID:       "SouthAfrica",
		PanamaID:            "Panama",
		CostaRicaID:         "CostaRica",
		NigeriaID:           "Nigeria",
		KenyaID:             "Kenya",
		MalaysiaID:          "Malaysia",
		PakistanID:          "Pakistan",
		ArgentinaID:         "Argentina",
		KazakhstanID:        "Kazakhstan",
		JamaicaID:           "Jamaica",
		ThailandID:          "Thailand",
	}
	name, ok := countryNameByID[id]
	return name, ok
}
