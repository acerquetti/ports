package database_test

import (
	"strings"
	"testing"

	"github.com/acerquetti/ports/domain"
	"github.com/acerquetti/ports/infra/database"
)

func TestCreateMemoryDB(t *testing.T) {
	expPortIDs := []string{"AEAJM", "CNBHY", "USJAX"}
	expPortNames := []string{"Ajman", "Beihai Fucheng Apt", "Jacksonville"}

	portsReader := strings.NewReader(`{
	"` + expPortIDs[0] + `": {
		"name": "` + expPortNames[0] + `",
		"city": "Ajman",
		"country": "United Arab Emirates",
		"alias": [],
		"regions": [],
		"coordinates": [
			55.5136433,
			25.4052165
		],
		"province": "Ajman",
		"timezone": "Asia/Dubai",
		"unlocs": [
			"AEAJM"
		],
		"code": "52000"
	},
	"AEFJR": {
		"name": "Al Fujayrah",
		"coordinates": [
			56.33,
			25.12
		],
		"city": "Al Fujayrah",
		"province": "Al Fujayrah",
		"country": "United Arab Emirates",
		"alias": [],
		"regions": [],
		"timezone": "Asia/Dubai",
		"unlocs": [
			"AEFJR"
		]
	},
	"` + expPortIDs[1] + `": {
		"name": "` + expPortNames[1] + `",
		"city": "Beihai Fucheng Apt",
		"province": "Guangxi",
		"country": "China",
		"alias": [],
		"regions": [],
		"coordinates": [
			108.327546,
			22.815478
		],
		"timezone": "Asia/Shanghai",
		"unlocs": [
			"CNBHY"
		],
		"code": "57076"
	},
	"BHAHD": {
		"country": "Bahrain",
		"province": "Bahrain",
		"city": "Al HIdd",
		"code": "52500",
		"name": "Al Hidd",
		"alias": [],
		"regions": [],
		"unlocs": [
			"BHAHD"
		]
	},
	"` + expPortIDs[2] + `": {
		"name": "` + expPortNames[2] + `",
		"city": "Jacksonville",
		"province": "Florida",
		"country": "United States",
		"alias": [],
		"regions": [],
		"coordinates": [
			-81.65565099999999,
			30.3321838
		],
		"timezone": "America/New_York",
		"unlocs": [
			"USJAX"
		],
		"code": "1803"
	  }
}'`)

	db, err := database.NewMemoryDB(portsReader)
	if err != nil {
		t.Fatal("unexpected error when instantiating memory db:", err)
	}

	var gotPortIDsCount int
	for expPortIndex, expPortID := range expPortIDs {
		gotPort, ok := db.Ports[domain.ID(expPortID)]
		if !ok {
			t.Fatalf("port with id %q expected in db", expPortID)
		}
		if domain.Name(expPortNames[expPortIndex]) != gotPort.Name {
			t.Fatalf("expected port with name %q, got %q", expPortNames[expPortIndex], gotPort.Name)
		}
		gotPortIDsCount++
	}

	if len(expPortIDs) != gotPortIDsCount {
		t.Fatalf("expected %v ports, got %v ports instead", len(expPortIDs), gotPortIDsCount)
	}
}
