# GO Hotel Offer Simulation


How to Run Simulation:
1. Specify RabbitMQ and MYSQL Server Configurations in below yamls (as per comments in yaml)
```
<repo>/hotel-sub/dependency/hotelOfferMgrConfig.yaml
<repo>/hotel-pub/dependency/hotelOfferPubConfig.yaml
```
2. Create Database hoteloffer
3. Go to hotel-sub folder and start binary 

```
cd <repo>/hotel-sub
./hotel-sub
```
4. In a new terminal Start Publisher Binary when Data has to be sent (one execution send message only once)
```
cd <repo>/hotel-pub
./hotel-pub
```

######NOTE
1. If Message sent has to be changed change `hotelData` variable in `<repo>/hotel-pub/def.go`
2. The above binaries are for amd64 arch. For different architecture Please build binary before run using below
```
cd <repo>/
cd hotel-sub
go build
cd ../hotel-pub
go build
```