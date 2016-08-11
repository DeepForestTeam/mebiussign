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

###API

####API Endpoint
All API URLs listed are relative to ```http://mobiussign.com/```. For example, the ```/api/sign/``` call is reachable at ```http://mobiussign.com/api/sign/```.

####Response format
Responses to all requests are always in [JSON](https://en.wikipedia.org/wiki/JSON) format.

####API Requests
Getting current **MobiusTime**:
```
/api/time
```
Can be requested using `GET` method. Returns the details of current MobiusTime mark.

Example request:
```
http://mobiussign.com/api/time
```

Example Response:

```json
{
  "result": "OK",
  "time_zone": "UTC",
  "time": "2016-08-10 13:17:41",
  "unix_time": 1470835061,
  "salt_hash": "C9415DDCF03E52598D8FFBFFFE90844E1CD3ED33CE600B9E2E8BE3632FE8CFD4",
  "pepper_hash": "FB4179D2EE768092",
  "mobius_time": "C808B0E2D0693C7ED68F8E65E470562EED1F581B49A3A69FA5E9D84B57B6DB93",
  "rsa_time": "n/a"
}
```

#####Check Mobius Time value
```
/api/time/{mobius_time}
```
Can be requested using GET method. The {mobius_time} is a MobiusTime mark hash.
Validates and returns the details of provided MobiusTime mark details.
The fields returned are the same as in the ‘/api/time’ request.
Example request:

```
http://mobiussign.com/api/time/51BB0B1727BC49FD60459518CCAF01EF00A149CE7F5F368CFB778F27EB80A136
```

Example Response:

```json
{
  "result": "OK",
  "time_zone": "UTC",
  "time": "2016-08-10 13:17:41",
  "unix_time": 1470835061,
  "salt_hash": "C9415DDCF03E52598D8FFBFFFE90844E1CD3ED33CE600B9E2E8BE3632FE8CFD4",
  "pepper_hash": "FB4179D2EE768092",
  "mobius_time": "C808B0E2D0693C7ED68F8E65E470562EED1F581B49A3A69FA5E9D84B57B6DB93",
  "rsa_time": "n/a"
}
```
In case of error:

```json
{
  "result": "Hash not found",
  "note": "",
  "result_code": 404
}
```

####Sign data

```
/api/sign
```

Can be requested using POST method, data encoded in JSON format. Signs user provided data.
Accepts the following params:

| Param |  JSON param | type | Format/Values|   |
|---|---|---|---|---|
| ServiceID |`service_id` | string | A-Za-z0-9 | *[optional]* should be provided only by registered services |
| ObjectID | `object_id` | string | A-Za-z0-9 | *[optional]* contains ID of customer’s associated object |
| ConsumerID | `consumer_id` | string | A-Za-z0-9 | *[optional]* contains ID of customer’s associated consumer |
| **Data Section** |
| DataUrl | `data_url` | string | | *[optional]* URL where the signed information can be found |
| DataNote | `data_note` | string |  | *[optional]* Any additional notes to the data |
| DataHash | `data_hash` | string |  | prepared data hash |
| DataHash | `data_block` | bin or string |  | data to be signed |
| DataBlockFormat | `data_format` | string | base64/string/hex | indicates what type of data was passed in the DataBlock field. Default value is ‘string’ |
| Security |
| ServiceSign | `service_sign` | string |  |  *[optional]*  should be provided only by registered services |

If `DataUrl` provided - it will be validated against URL standards and service will return error in case if validation fails.
One of the DataHash or DataBlock fields is always required. You can send only one of them or both, but at least one of them is required.

If only `DataHash` provided - system will just sign it.

If only `DataBlock` provided - system will create the hash itself and store both DataBlock and DataHash.

If both `DataHash` and `DataBlock` provided - system will create hash from DataBlock (using *SHA2* *512* or *512-256*), will check it against provided `DataHash` and return error `error in hash block` if they differ.

If data you provide in `DataBlock` is binary - you should encode it into `base64` or `hex` first, also you should provide the appropriate format description in `DataBlockFormat` field.

Fields ServiceID and ServiceSign should be provided only by registered customers. For regular users this fields will be ignored. If ServiceID and ServiceSign are used for registered customers and ServiceSign will be incorrect for the provided ServiceID and data - `signature error` will be returned. 

If some data required for signing the data is missing - the appropriate error will be returned.

Returns the following fields: 

| JSON Field | Description |
|---|---|
| `result` | result code of the request execution. Should have `OK` value for all successful requests. |
| `sign_id` | short signature ID (can be used to check the signature as shortcat instead of full `mobius_sign`) |
| `row_id` | database record id |
| `block_id`	 | *[reserved]* |
| `service_id` | unique ID of the MobiusSign™ registred customer |
| `object_id` | field provided by MobiusSign™ customer with ID of customer’s associated object |
| `consumer_id` | field provided by MobiusSign™ customer with ID of customer’s associated consumer |
| `data_url` | URL where the signed information can be found |
| `data_note` | any additional notes to the signed data |
| `data_hash` | hash of the signed data |
| `data_block` | signed data |
| `data_format` |  signed data format (string, base64 or hex) |
| `time` | time when the data was signed |
| `unix_time` | time when the data was signed as UnixTimeStamp |
| `time_hash` | MobiusTime hash for the moment when data was signed |
| `salt_id` | `sign_id` of previous MobiusSign record, which was used as a salt |
| `salt_hash` | previous MobiusSign record, which was used as a salt |
| `pepper_hash` | random 8 bytes in hex string representation |
| `mobius_sign` | MobiusSign of signed data |
| `rsa_sign ` | *[reserved* |

Example Response:

```json
{
  "result": "OK",
  "sign_id": "51A502C5CEB5457C",
  "row_id": 52,
  "block_id": "0000000",
  "service_id": "",
  "object_id": "",
  "consumer_id": "",
  "data_url": "",
  "data_note": "",
  "data_hash": "E157DF98228347C240C860715AF46DCC6D67F37F67855AA7A5888488262E1B916829872ADE39DB2C619A9A1E02A64870BB21568108E0225239E37AA8351F2388",
  "data_block": "The Ultimate Question of Life, the Universe, and Everything",
  "data_format": "string",
  "time": "2016-08-10 13:58:49",
  "unix_time": 1470837529,
  "time_hash": "7B09AC5D090A4A9FCA9AB8A3493D2B4D0B0BE0200AF78A1F3ED7FF373F65B891",
  "salt_id": "438956E0A64A6788",
  "salt_hash": "A2CA1F51187E334C07475CF4CF7171C27B2312E0B7F07F8F84F0EA061E1E18FB751C70BB742378D110412435B1BD94EEADE3FBE30023A82349AEF776B25052C8",
  "pepper_hash": "2BEF6F728301D40A",
  "mobius_sign": "4EF5E0F1E903F1BBD437E9CA8FCF50F36FFA817E4D57BA555B93A42A536E53E5B9120E7DE9BD5CF820DB18BE7CB2525055F9F80B693671F7CA021D4DD6D597F4",
  "rsa_sign": "n/a"
}
```

#####Validate MobiusSign Data Signature

```
/api/sign/{mobius_sign}
```

Can be requested using `GET` method. The `{mobius_sign}` is a MobiusSign (`mobius_sign`) hash or MobiusSign ID (`sign_id`).

Validates and returns the details of provided MobiusSign/MobiusSignID details.
The fields returned are the same as in the ‘/api/sign’ request.
Example request:

```
http://mobiussign.com/api/sign/9C5B52B8914B30AB0F4459F62B34D1FEFD5B2957E0E73B874D6AD9E8C3303208523F261F8BD4233A81A51142EC776DB9211875B6B89BFB29FDC256DE8625824C
```

Example Response:

```json
{
  "result": "OK",
  "sign_id": "51A502C5CEB5457C",
  "row_id": 52,
  "block_id": "0000000",
  "service_id": "",
  "object_id": "",
  "consumer_id": "",
  "data_url": "",
  "data_note": "",
  "data_hash": "E157DF98228347C240C860715AF46DCC6D67F37F67855AA7A5888488262E1B916829872ADE39DB2C619A9A1E02A64870BB21568108E0225239E37AA8351F2388",
  "data_block": "The Ultimate Question of Life, the Universe, and Everything",
  "data_format": "string",
  "time": "2016-08-10 13:58:49",
  "unix_time": 1470837529,
  "time_hash": "7B09AC5D090A4A9FCA9AB8A3493D2B4D0B0BE0200AF78A1F3ED7FF373F65B891",
  "salt_id": "438956E0A64A6788",
  "salt_hash": "A2CA1F51187E334C07475CF4CF7171C27B2312E0B7F07F8F84F0EA061E1E18FB751C70BB742378D110412435B1BD94EEADE3FBE30023A82349AEF776B25052C8",
  "pepper_hash": "2BEF6F728301D40A",
  "mobius_sign": "4EF5E0F1E903F1BBD437E9CA8FCF50F36FFA817E4D57BA555B93A42A536E53E5B9120E7DE9BD5CF820DB18BE7CB2525055F9F80B693671F7CA021D4DD6D597F4",
  "rsa_sign": "n/a"
}
```

In case of error:

```json
{
  "result": "Signature not found",
  "note": "",
  "result_code": 404
}
```

##PHP Usage example

```php
<?php
/**
 * MobiusSign PHP Example
 *
 */

// Compose you sign request:

$SignRequest = array(
    'object_id' => "0",
    'consumer_id' => "1",
    'data_url' => 'http://example.site/document.html',
    'data_note' => 'Example of MobiusSign usage',
    'data_block' => 'The Ultimate Question of Life, the Universe, and Everything:42',
    'data_hash' => '',
);
//Now you can calculate data hash  SHA512
$SignRequest['data_hash'] = hash('sha512', $SignRequest['data_block']);
echo "DH:" . $SignRequest['data_hash'] . PHP_EOL;
//Now you can generate JSON
$JSonRequest = json_encode($SignRequest);

// ...and make API Request

$Options = array(
    'http' => array(
        'method' => "POST",
        'header' => "Content-Type: application/json; charset=utf-8\r\n",
        'content' => $JSonRequest
    )
);
$r = stream_context_create($Options);
$Result = file_get_contents("http://mobiussign.com/api/sign", null, $r);

// decoding result
$ResultArray = json_decode($Result, true);

// check for valid response from service
if (!isset($ResultArray['result']) || !is_array($ResultArray))
{
	echo 'Something went absolutely wrong - incorrect answer from server'.PHP_EOL;
	die();
}

// handling successful response
if ($ResultArray['result'] == 'OK')
{
	echo 'Your data was signed successfully'.PHP_EOL;
	echo 'MobiusSign: '.$ResultArray['mobius_sign'].PHP_EOL;
	echo 'MobiusSign ID: '.$ResultArray['sign_id'].PHP_EOL;
	die();
}

// output the error
echo 'Error occurred: '.$ResultArray['result'].PHP_EOL;
```

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

```
DataBlock:=([SaltHash][Pepper][ServiceID][ObjectID][ConsumerID][DataHash][MobiusTime])
```
Where: 

- ```SaltHash``` - previous element in sequence (previous MobiusSign generated by the system)
- ```Pepper``` - random 8 bytes in hex string representation
- ```ServiceID``` [optional] - unique ID of the MobiusSign™ customer
- ```ObjectID``` [optional] - field provided by MobiusSign™ customer with ID of customer’s object
- ```ConsumerID``` [optional] - field provided by MobiusSign™ customer with ID of customer’s consumer
- ```DataHash``` - hash of user-provided data
- ```MobiusTime``` - current value of MobiusTime

```DataHash``` can be provided directly as a hash or by data itself (in this case MobiusSign™ will create hash with SHA2 512 algorithm itself).



#### Used packets:
* https://github.com/op/go-logging
* https://github.com/boltdb/bolt
* https://github.com/gorilla/mux