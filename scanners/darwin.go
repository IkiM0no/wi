package scanners

import (
	"regexp"
	//"os/exec"
	"strings"
	"wi/regexputils"
)

var reLines = regexp.MustCompile(`\s*(?P<ssidName>.+?)\s*(?P<ssidMac>([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2}))\s*(?P<ssidRSSI>[\-\d]+)\s*(?P<ssidChannel>[\d,\+\-]{1,})\s*(?P<ssidHT>[Y|N]{1})\s*(?P<ssidCC>[A-Z\-]{2})\s*(?P<ssidSecurity>.*)`)

func DarwinScan(app string, args []string) (WifiNeighbors, error) {
	var sm WifiNeighbors
	sm.ClientType = "Darwin"

	// TESTS
	//str := readFile("~/go/src/wi/darwin_multi_bssid")

	// PROD
	c := Cmd{App: app, Arg: args}
	str := c.Exec()

	lines := strings.Split(string(str), "\n")
	lines = utils.KeepSliceReMatches(lines, reLines)

	var ssidNameList []string
	for _, s := range lines {
		// SSIDs
		var ssidIdx SSID
		p := &ssidIdx
		ssidName := utils.ReSubMatchMap(reLines, s)["ssidName"]
		if ! utils.StringInSlice(ssidName, ssidNameList) {
			ssidNameList = append(ssidNameList, ssidName)
			p.SSIDName   = utils.ReSubMatchMap(reLines, s)["ssidName"]
			p.Encryption = utils.ReSubMatchMap(reLines, s)["ssidSecurity"]
			sm.SSIDs     = append(sm.SSIDs, ssidIdx)
		}
		// BSSIDs
		var bSsidIdx BSSID
		pB := &bSsidIdx
		pB.BssidMac        = utils.ReSubMatchMap(reLines, s)["ssidMac"]
		pB.BssidSignal     = utils.ReSubMatchMap(reLines, s)["ssidRSSI"]
		pB.BssidChannel    = utils.ReSubMatchMap(reLines, s)["ssidChannel"]
		pB.BssidEncryption = utils.ReSubMatchMap(reLines, s)["ssidSecurity"]
		pB.BssidHT         = utils.ReSubMatchMap(reLines, s)["ssidHT"]
		pB.BssidCC         = utils.ReSubMatchMap(reLines, s)["ssidCC"]

		// Append BSSIDs to corresponding SSIDs
		lenSmSSIDs := len(sm.SSIDs)
		for i := 0; i < lenSmSSIDs; i++ {
			if sm.SSIDs[i].SSIDName == ssidName {
				sm.SSIDs[i].BSSIDS = append(sm.SSIDs[i].BSSIDS, bSsidIdx)
			}
		}
	}
	return sm, nil
}
