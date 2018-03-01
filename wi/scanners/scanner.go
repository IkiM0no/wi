package scanners

import (
	"fmt"
	"bufio"
	"os/exec"
	"io/ioutil"
)

type WifiNeighbors struct {
	ClientType    string `json:"client_type"`
	CurrentUTC    string `json:"current_utc"`
	SSIDs         []SSID `json:"ssids"`
}

type SSID struct {
	SSIDName        string `json:"ssid_name"`
	NetworkType     string `json:"network_type,omitempty"`
	Authentication  string `json:"auth,omitempty"`
	Encryption      string `json:"encryption,omitempty"`
	BSSIDS          []BSSID `json:"bssids,omitempty"`
}

type BSSID struct {
	BssidMac        string `json:"mac,omitempty"`
	BssidSignal     string `json:"signal,omitempty"`
	BssidRadioType  string `json:"radio_type,omitempty"`
	BssidChannel    string `json:"channel,omitempty"`
	BssidEncryption string `json:"bssid_encryption,omitempty"`
	BssidBasicRates string `json:"basic_rates,omitempty"`
	BssidOtherRates string `json:"other_rates,omitempty"`
	BssidHT         string `json:"ht,omitempty"`
	BssidCC         string `json:"cc,omitempty"`
}

type Cmd struct {
	App string
	Arg []string
}

// Return cmd output as string
func (c Cmd) Exec() (string) {
	cmd := exec.Command(c.App, c.Arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}
	in := bufio.NewScanner(stdout)
	var cmdOutput string

	for in.Scan() {
		str := in.Text()
		cmdOutput += (str + "\n")
	}
	return cmdOutput
}

// TMP for Darwin - read from static file
func readFile(path string) (string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)
	return str
}
