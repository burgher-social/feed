package PSQL

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbinstance *gorm.DB

func Connect() *gorm.DB {
	if dbinstance == nil {
		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		port := os.Getenv("DB_PORT")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", host, user, password, dbname, port)
		// fmt.Println(dsn)
		// dsn := "host=localhost user=user password=password dbname=burgher port=5432 sslmode=disable TimeZone=Asia/Kolkata"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		dbinstance = db
	}
	return dbinstance
}

// var connectionString string = "127.0.0.1"
// var bucketName string = "burgher"
// var username string = "Administrator"
// var password string = "password"

// var once sync.Once
// var cluster *gocb.Scope

// func Cluster() *gocb.Scope {
// 	once.Do(func() {
// 		cb, clusterError := gocb.Connect(connectionString, gocb.ClusterOptions{
// 			Username: username,
// 			Password: password,
// 		})

// 		if clusterError != nil {
// 			panic(clusterError)
// 		}

// 		bucket := cb.Bucket(bucketName)
// 		print(bucket)

// 		cluster = bucket.Scope("burgher")
// 	})
// 	return cluster
// }

// // File: mongoDBConnection.go
// // Open new connection
// func setupCouch() {
// 	cluster, err := gocb.Connect("couchbase://"+connectionString, gocb.ClusterOptions{
// 		Authenticator: gocb.PasswordAuthenticator{
// 			Username: username,
// 			Password: password,
// 		},
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	bucket := cluster.Bucket(bucketName)

// 	err = bucket.WaitUntilReady(5*time.Second, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	col := bucket.Scope("tenant_agent_00").Collection("users")
// 	_, err = col.Upsert("u:jade",
// 		User{
// 			Name:      "Jade",
// 			Email:     "jade@test-email.com",
// 			Interests: []string{"Swimming", "Rowing"},
// 		}, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
// func SetupMongoDB() (*mongo.Collection, *mongo.Client, context.Context, context.CancelFunc) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@mongo:27017/"))
// 	if err != nil {
// 		panic(fmt.Sprintf("Mongo DB Connect issue %s", err))
// 	}
// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		panic(fmt.Sprintf("Mongo DB ping issue %s", err))
// 	}
// 	collection := client.Database("mongo-golang-test").Collection("Users")
// 	return collection, client, ctx, cancel
// }

// // Close the connection
// func CloseConnection(client *mongo.Client, context context.Context, cancel context.CancelFunc) {
// 	defer func() {
// 		cancel()
// 		if err := client.Disconnect(context); err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("Close connection is called")
// 	}()
// }
