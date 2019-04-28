package config

//MongoPort port on which mongo is running
var MongoPort = "27017"

//MongoIP server ip fror mongo server
var MongoIP = "192.168.0.108"

//MongoProto protocol to use to connect to mongo
var MongoProto = "mongodb"

type key string

const (
	ProtoKey = key("protoKey")
	IPKey    = key("ipKey")
	PortKey  = key("portKey")
)
