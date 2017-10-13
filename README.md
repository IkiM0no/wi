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
