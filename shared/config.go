package shared

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2"

	pbList "github.com/wizelineacademy/GoWorkshop/proto/list"
	pbNotifier "github.com/wizelineacademy/GoWorkshop/proto/notifier"
	"google.golang.org/grpc"
)

// Init bootstrap entry point
func Init() {
	initConfig()
	initServices()
}

type (
	configuration struct {
		Server, ListService, UsersService, NotifierService, MongoDBHost, DBUser, DBPwd, Database, SMTPHost, SMTPPort, SMTPUser, SMTPPass string
	}
)

var (
	// AppConfig var
	AppConfig      configuration
	mongoSession   *mgo.Session
	listClient     pbList.ListClient
	notifierClient pbNotifier.NotifierClient
)

// Initialize AppConfig
func initConfig() {
	file, err := os.Open("../shared/config.json")
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
	connList, _ := grpc.Dial(AppConfig.ListService, grpc.WithInsecure())
	listClient = pbList.NewListClient(connList)

	connNotifier, _ := grpc.Dial(AppConfig.NotifierService, grpc.WithInsecure())
	notifierClient = pbNotifier.NewNotifierClient(connNotifier)
}

func GetListClient() pbList.ListClient {
	return listClient
}

// GetNotifierClient func
func GetNotifierClient() pbNotifier.NotifierClient {
	return notifierClient
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
