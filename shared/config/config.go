package config

import (
	"encoding/json"
	"log"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"

	pb "github.com/wizelineacademy/GoWorkshop/proto/list"
	"google.golang.org/grpc"
)

// Init bootstrap entry point
func Init() {
	initConfig()
	initServices()
}

type (
	configuration struct {
		Server, ListService, UsersService, NotifierService, MongoDBHost, DBUser, DBPwd, Database string
	}
)

// AppConfig var
var AppConfig configuration

var mongoSession *mgo.Session

var listService pb.ListClient

// Initialize AppConfig
func initConfig() {
	file, err := os.Open("../shared/config/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[logAppConfig]: %s\n", err)
	}
}

// Initialize connection to other services
func initServices() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(AppConfig.ListService, grpc.WithInsecure())
	if err != nil {
		log.Printf("cannot connect to list service: %v", err)
	}

	listService = pb.NewListClient(conn)
}

// GetListService func
func GetListService() pb.ListClient {
	return listService
}

// GetSession returns database mongoSession
func GetSession() *mgo.Session {
	if mongoSession == nil {
		var err error
		mongoSession, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{AppConfig.MongoDBHost},
			Username: AppConfig.DBUser,
			Password: AppConfig.DBPwd,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}
	return mongoSession
}
