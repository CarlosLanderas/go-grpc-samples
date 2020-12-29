package dbclient

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	. "go-grpc-samples/user-service"
	"log"
)

const UsersBucket = "Users"

type IBoltClient interface {
	GetUser(id int64) (User, error)
	Seed()
	OpenDb()
}

type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) GetUser(id int64) (User, error) {

	user := User{}

	err := bc.boltDB.View(func (tx *bolt.Tx) error {
		b := tx.Bucket([]byte(UsersBucket))
		userBytes := b.Get([]byte(string(id)))

		if userBytes == nil {
			return fmt.Errorf("No user found for %s", id)
		}

		json.Unmarshal(userBytes, &user)

		return nil
	})

	if err != nil {
		return user, err
	}

	return user, nil
}

func (bc *BoltClient) Seed()  {

	err := bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(UsersBucket))
		if err !=  nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})

	if err != nil { panic(err)}

	var i int64

	for i = 0; i < 100; i ++ {
		name := "User_" + string(i)
		user := User{ Id: i, Name:  name}

		jsonBytes, _ := json.Marshal(user)
		bc.boltDB.Update(func (tx *bolt.Tx) error {
			b := tx.Bucket([]byte(UsersBucket))
			err := b.Put([]byte(string(i)), jsonBytes)
			return err
		})
	}

	fmt.Println("Users seed completed...")
}

func (bc *BoltClient) OpenDb() {

	var err error
	bc.boltDB, err = bolt.Open("users.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDatabase() BoltClient {
	return BoltClient{}
}