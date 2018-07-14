# pault.ag/go/ais

**This is also absolutely not ready to be used for navigational purposes.**
Please do not use this library in any situation that may cause harm to life
or property.

`go-ais` is a library to handle reading and parsing NMEA AIS messages into
an idiomatic and workable set of Go types.

This library is in a very raw and prototypical format, it can not parse all
the AIS message types yet, and it may have errors that make it unsuitable for
production.

## How to catch AIS messages with rtl_ais

```
./rtl_ais -l 161.966M -r 162.015M -g 48 -n
```
