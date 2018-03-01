package scanners

import (
	"fmt"
	"time"
	"regexp"
	"strings"
	"wi/regexputils"
)

// Raw Splitter Regex
var reWinSSID = regexp.MustCompile(`SSID\s(?P<ssidNum>\d+)\s:(\s(?P<ssidName>[A-Za-z0-9-\s_>]{1,}))?.*`)
var reBSSID   = regexp.MustCompile(`(?P<bssidSplit>\s{4})?BSSID`)

// SSID Blob parse Regex
var reSSIDName       = regexp.MustCompile(`.*SSID\s{1,}\d{1,}\s{1,}:\s{1,}(?P<ssidName>[A-Za-z\s\-\d]{1,})?.*`)
var reSSIDNetType    = regexp.MustCompile(`\s{1,}Network\stype\s{2,}:\s(?P<ssidNetworkType>[A-Za-z\s\-\d]{1,})?.*`)
var reSSIDAuthType   = regexp.MustCompile(`\s{1,}Authentication\s{2,}:\s(?P<ssidAuthType>[A-Za-z\s\-\d]{1,})?.*`)
var reSSIDEncryption = regexp.MustCompile(`\s{1,}Encryption\s{2,}:\s(?P<ssidEncryption>[A-Za-z]{1,})?\s.*`)

//BSSID Blob parse Regex
var reBSSIDMac       = regexp.MustCompile(`BSSID:\s{1,}\d{1,}\s{1,}:\s{1,}(?P<bssidMac>([\dA-Fa-f]{2}[:-]){5}([\dA-Fa-f]{2}))?.*`)
var reBSSIDSignal    = regexp.MustCompile(`\s{1,}Signal\s{1,}:\s{1}(?P<signal>[\d]{1,}\%)?.*`)
var reBSSIDRadio     = regexp.MustCompile(`\s{1,}Radio\stype\s{1,}:\s{1}(?P<radioType>[\d]{1,}\.[\d]{1,}[A-Za-z]{1,})?.*`)
var reBSSIDChannel   = regexp.MustCompile(`\s{1,}Channel\s{2,}:\s(?P<channel>([\d]{1,3}){1,})?.*`)
var reBSSIDBasicRate = regexp.MustCompile(`\s{1,}Basic\srates\s\(Mbps\)\s:\s(?P<basicRates>[\d+\s\.*]+)?`)
var reBSSIDOtherRate = regexp.MustCompile(`\s{1,}Other\srates\s\(Mbps\)\s:\s(?P<otherRates>[\d+\s\.*]+)?`)

func WinScan(winCmd string, winArg []string) (WifiNeighbors, error) {
	var ssidMetrics WifiNeighbors

	t := time.Now().Format(time.RFC3339)
	ssidMetrics.CurrentUTC = t

	c := Cmd{App: winCmd, Arg: winArg}
	strParse := c.Exec()

	var recordsSplit     []string
	var headers          [][]string
	var ssidHeaderString []string
	var bssidSlices      [][]string
	var bssidSlice       []string

	recordsSplit = strings.Split(strParse, "\n\n")

	// Throw away "Interface Name: WiFi \n There are n networks available."
	recordsSplit = recordsSplit[1:]
	recordsSplit = utils.DeleteEmpty(recordsSplit)

	// Split SSIDs into blobs. Split BSSIDs into a list of string blobs.
	for i := range recordsSplit{
		// Transform BSSID to KV pair format in header
		subMatchB := utils.ReSubMatchMap(reBSSID, recordsSplit[i])
		replB := fmt.Sprintf("<BSSID_delim>BSSID:%s", subMatchB["bssidSplit"])
		recordsSplit[i] = reBSSID.ReplaceAllString(recordsSplit[i], replB)

		// Parse SSID Header to blob
		h, h2 := utils.GetHeader(recordsSplit[i], 4)
		h = utils.DeleteEmpty(h)
		headers = append(headers, h)
		ssidHeaderString = append(ssidHeaderString, h2)

		// Parse BSSID to slice of blobs
		bssidSlice = strings.Split(recordsSplit[i], "<BSSID_delim>")
		bssidSlice = utils.DeleteEmpty(bssidSlice)
		if len(bssidSlice) > 1 {
			bssidSlice = bssidSlice[1:]
		}
		bssidSlices = append(bssidSlices, bssidSlice)

		// Parse SSIDs to winSSID
		var winSSIDIdx SSID
		ssidBlobSplit := strings.Split(ssidHeaderString[i], "\n")

		// SSIDs
		for _, l := range ssidBlobSplit {
			switch {
				case (reSSIDName.MatchString(l)) :
					winSSIDIdx.SSIDName    = utils.ReSubMatchMap(reSSIDName, l)["ssidName"]
				case (reSSIDNetType.MatchString(l)) :
					winSSIDIdx.NetworkType = utils.ReSubMatchMap(reSSIDNetType, l)["ssidNetworkType"]
				case (reSSIDAuthType.MatchString(l)) :
					winSSIDIdx.Authentication = utils.ReSubMatchMap(reSSIDAuthType, l)["ssidAuthType"]
				case (reSSIDEncryption.MatchString(l)) :
					winSSIDIdx.Encryption = utils.ReSubMatchMap(reSSIDEncryption, l)["ssidEncryption"]
			}
		}
		// BSSIDs
		var bSSIDIdxList []BSSID
		for _, bBlob := range bssidSlices[i] {
			bLineSplit := strings.Split(bBlob, "\n")
			var bSSIDIdx BSSID
			for _, l := range bLineSplit {
				switch {
					case (reBSSIDMac.MatchString(l)) :
						bSSIDIdx.BssidMac = utils.ReSubMatchMap(reBSSIDMac, l)["bssidMac"]
					case (reBSSIDSignal.MatchString(l)) :
						bSSIDIdx.BssidSignal = utils.ReSubMatchMap(reBSSIDSignal, l)["signal"]
					case (reBSSIDRadio.MatchString(l)) :
						bSSIDIdx.BssidRadioType = utils.ReSubMatchMap(reBSSIDRadio, l)["radioType"]
					case (reBSSIDChannel.MatchString(l)) :
						bSSIDIdx.BssidChannel = utils.ReSubMatchMap(reBSSIDChannel, l)["channel"]
					case (reBSSIDBasicRate.MatchString(l)) :
						bSSIDIdx.BssidBasicRates = utils.ReSubMatchMap(reBSSIDBasicRate, l)["basicRates"]
					case (reBSSIDOtherRate.MatchString(l)) :
						bSSIDIdx.BssidOtherRates = utils.ReSubMatchMap(reBSSIDOtherRate, l)["otherRates"]
				}
			}
			bSSIDIdxList = append(bSSIDIdxList, bSSIDIdx)
		}
		// Append BSSID List to winSSIDIdx
		winSSIDIdx.BSSIDS = bSSIDIdxList
		// Append SSIDs to wifiNeighbors
		ssidMetrics.SSIDs = append(ssidMetrics.SSIDs, winSSIDIdx)
		ssidMetrics.ClientType = "Windows"
	}
	return ssidMetrics, nil
}
