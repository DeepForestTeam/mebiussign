# Möbius Sign™
 *Möbius Sign™* is a up-to-date technology of cryptographic security from information changing which corresponding to BlockChain, allows any wishful to ensure in the integrality of not only single document, fact and docchain, but also of the whole system information in general

## MobiusTime API
MobiusTime API allow get time stamp, signed by SHA2 hash by http get request: `/api/time`
```
{
  "time_zone": "UTC",
  "time": "2016-07-26 08:04:52",
  "unix_time": 1469520292,
  "salt_hash": "BF5D2961D09BA0478AA684198C6FAB216417E53AB406EEE81E8FF35E4E95EE0A",
  "pepper_hash": "290E12CC7E8B491F",
  "mobius_time": "73578B6826998B293FCE468A8CC72EDCA66109A50B212A200902286701594170",
  "rsa_time": "n/a"
}
```
Now you can check this time-stamp: `/api/time/73578B6826998B293FCE468A8CC72EDCA66109A50B212A200902286701594170`

#### Used packets:
* https://github.com/op/go-logging
* https://github.com/boltdb/bolt
