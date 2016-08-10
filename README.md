# Möbius Sign™
 *Möbius Sign™* is a up-to-date technology of cryptographic security from information changing which corresponding to BlockChain, allows any wishful to ensure in the integrality of not only single document, fact and docchain, but also of the whole system information in general

## MobiusSign API
The service is generally a storage of consistent hashes linked to time marks which ensures integrity and continuity of the signatures chain. The storage is organized in the way which avoids any ability of modification or removal of the chain links without affecting all next links.

Service consists of two main parts:

- **MobiusTime** - mechanism of creating unique time marks
- **MobiusSign** - mechanism of integral continuous signs chain

**Attention:**
*You should always remember that all information gathered by **MobiusSign™** is considered as public data.*

If you intend to sign any confidential data you should use one of the following methods:

- You can skip the DataBlock and pass only DataHash. This way you can confirm the data authenticity without making it public. The data itself should be stored on the side of the service which requested the signature. This way you can also avoid the data size limit. If you create DataHash by yourself we recommend appending MobiusTime to your data to ensure the signing time on your side.
- You can send your pre-encrypted information to DataBlock. The encryption should be made with any resistant algorithm without passing keys to our service. This data volume is limited to 32Kb.

###Algorithms

####MobiusTime™
The time mark (time hash) is an sha512.Sum512_256 encoded **DataBlock**.
**DataBlock** consists of merged fields:```SaltHash``` , ```SignPepper```, ```UnixTimeStampBytes``` and ```BaseSalt```.

```
DataBlock:=([SaltHash][SignPepper][UnixTimeStampBytes][BaseSalt])
```
Where:

- ```SignPepper``` - random 8 bytes in hex string representation
- ```BaseSalt``` - basic salt string for the whole project
- ```SaltHash``` - previous time mark hash (for the first mark it has a BaseSalt value)
- ```UnixTimeStampBytes``` - int64 value of the current UnixTimeStamp

**MobiusTime** value example: ```A7532AEFD875EB4A159A1C02DFF3DE59DF6CE72F592B192543C01F0493F4AEB2```

####MobiusSign™
The ***MobiusSign*** is a simple SHA2 512 value of given DataBlock which consists of several merged values.

#### Used packets:
* https://github.com/op/go-logging
* https://github.com/boltdb/bolt
