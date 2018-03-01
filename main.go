package main

import (
	"os"
	"fmt"
	"time"
	"runtime"
	"encoding/json"

	"wi/scanners"
)

const WinCmd = "/mnt/c/Windows/System32/netsh.exe"
var WinArg   = []string{"wlan", "show", "networks", "mode=bssid"}

const DarwinCmd = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"
var DarwinArg   = []string{"-s"}

var runTime = runtime.GOOS


func runner() {
	for {
		s := fetchScan()
		jsonify(s, pp)
		sleep()
	}
}

func fetchScan() scanners.WifiNeighbors {
	var m scanners.WifiNeighbors
	switch runTime {
	case "linux":
		m, _ = scanners.WinScan(WinCmd, WinArg)
		fmt.Println("Coming Soon")
		os.Exit(1)
		return m
	case "darwin":
		m, _ = scanners.DarwinScan(DarwinCmd, DarwinArg)
		return m
	case "windows":
		m, _ = scanners.WinScan(WinCmd, WinArg)
		return m
	default:
		fmt.Println(`{"os": "Unknown/Unsupported runtime detected"}`)
		os.Exit(0)
		return m
	}
}

func jsonify(w scanners.WifiNeighbors, prettyPrint bool) {
	if prettyPrint {
		j, err := json.MarshalIndent(w, "", "    ")
		if err != nil {
		        fmt.Println("Json err:", err)
			return
		}
		sj := string(j)
		fmt.Println(sj)
	} else {
		s, err := json.Marshal(w)
		if err != nil {
		        fmt.Println("Json err:", err)
			return
		}
		fmt.Println(string(s))
	}
}

func sleep() {
	switch {
	case interval != 0:
		time.Sleep(time.Duration(interval) * time.Second)
	default:
		time.Sleep(time.Duration(10) * time.Second)
	}
}

