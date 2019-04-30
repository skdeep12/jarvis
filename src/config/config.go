package config

//MongoPort port on which mongo is running
var MongoPort = "27017"

//MongoIP server ip fror mongo server
var MongoIP = "127.0.0.1"

//MongoProto protocol to use to connect to mongo
var MongoProto = "mongodb"

type key string

const (
	ProtoKey = key("protoKey")
	IPKey    = key("ipKey")
	PortKey  = key("portKey")
)
