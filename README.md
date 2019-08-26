# sawtooth-onlinestore
Offer five functions for managing the goods of online store

- Buy: Get new amount of goods in an address
- Sell: Remove amount of goods in an address
- Tranport: Transport goods to somewhere
- Empty:Clear the warehouse in address
- Show: Show the amount of goods in address

### Install and Use

```
// * means relevant url and workspace
$ git clone * 
$ cd *
```

#### 1. Build in Go

```
$ cd Transaction-Processor
$ go build src/sawtooth_onlinestore/onlinestore.go
```

This commend generate an onlinestore executiontable file, which will be used in latter docker file

#### 2. Use sawtooth with docker

Docker enviroment is required, however you doesn't need a sawtooth environment on your host, all the service with be run on docker.

```
$ docker-compose -f onlinestore-build.yaml -d
$ docker exec -it onlinestore-client bash
```

#### 3. Examples

```
// form
$ onlinestore action value adress
// use example
$ sawtooth keygen
$ python3 setup.py install
$ sawtooth keygen beijing
$ sawtooth keygen hangzhou 
// 1. buy
$ onlinestore buy 100 beijing
// 2. sell
$ onlinestore sell 10 beijing
// 3. show
$ onlinestore show beijing
// 4. empty
$ onlinestore empty bejing
// 5. transport 
$ onlinestore transport 20 beijing hangzhou
```

