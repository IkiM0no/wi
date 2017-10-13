# wi

wi is a humble acknowledgement of a lack of a general client-based toolkit for obtaining RF Wi-Fi metrics in Go .

It aims to provide a minimal API for interacting with system-level Wi-Fi utilities on the client such as ```airport``` for Mac OS X and ```netsh``` for Windows with Linux support coming very soon.

wi includes a minimal CLI interface as a demo. It will confinually scan neighboring SSIDs and BSSIDs and return these results as JSON. Just build and run:

```
wi -i 5
```
to scan and return results every 5 seconds.

For help, invoke:
```
wi -h
```

Planned Features:
- Linux support.
- Send results over websocket to a central server to collect/log/aggregate metrics.

Sample output Windows:
```
Windows OS detected
{
    "client_type": "Windows",
    "ssids": [
        {
            "ssid_name": "wuzzle",
            "network_type": "Infrastructure",
            "auth": "WPA2-Personal",
            "encryption": "CCMP",
            "bssids": [
                {
                    "mac": "00:00:00:00:00:00:",
                    "signal": "88%",
                    "radio_type": "802.11ac",
                    "channel": "157",
                    "basic_rates": "6 12 24",
                    "other_rates": "9 18 36 48 54"
                }
            ]
        }
    ]
}
```
Sample output Mac OS:
```
Mac OS detected
{
    "client_type": "Darwin",
    "ssids": [
        {
            "ssid_name": "wuzzle",
            "encryption": "WPA2(PSK/AES/AES) ",
            "bssids": [
                {
                    "mac": "00:00:00:00:00:00",
                    "signal": "-81",
                    "channel": "136",
                    "bssid_encryption": "WPA2(PSK/AES/AES) ",
                    "ht": "Y",
                    "cc": "US"
                }
            ]
        },
        {
            "ssid_name": "woozie",
            "encryption": "WPA(PSK/AES,TKIP/TKIP) WPA2(PSK/AES,TKIP/TKIP) ",
            "bssids": [
                {
                    "mac": "00:00:00:00:00:00",
                    "signal": "-43",
                    "channel": "11",
                    "bssid_encryption": "WPA(PSK/AES,TKIP/TKIP) WPA2(PSK/AES,TKIP/TKIP) ",
                    "ht": "Y",
                    "cc": "--"
                }
            ]
        },
...
}
```
