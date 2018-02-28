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

func runner() {
	var runTime = runtime.GOOS
	var m scanners.WifiNeighbors
	switch runTime {
	case "linux":
		fmt.Println({"os": "Unix/Linux"})
		m, _ = scanners.WinScan(WinCmd, WinArg)
		//m, _ = scanners.DarwinScan(DarwinCmd, DarwinArg)
		//fmt.Println("Coming Soon")
		//os.Exit(1)
	case "darwin":
		fmt.Println({"os": "Darwin"})
		m, _ = scanners.DarwinScan(DarwinCmd, DarwinArg)
	case "windows":
		fmt.Println({"os": "Windows"})
		m, _ = scanners.WinScan(WinCmd, WinArg)
	default:
		fmt.Println({"os": "Unknown/Unsupported runtime detected"})
		os.Exit(0)
	}
	loop(m)
}

func loop(m scanners.WifiNeighbors) {
	for {
		jsonify(m)
		sleep()
	}
}

func jsonify(w scanners.WifiNeighbors) {
	j, err := json.MarshalIndent(w, "", "    ")
	if err != nil {
	        fmt.Println("Json err:", err)
	}
	sj := string(j)
	fmt.Println(sj)
}

func sleep() {
	switch {
	case interval != 0:
		time.Sleep(time.Duration(interval) * time.Second)
	default:
		time.Sleep(time.Duration(10) * time.Second)
	}
}
