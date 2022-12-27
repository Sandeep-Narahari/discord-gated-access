package database

import (
	"fmt"
	"log"
)

type nft struct {
	address       string
	nftId         string
	creator       string
	collection_id string
}
type gatedList struct {
	collection_id string
	role_name     string
	roleid        string
}

func (db *Db) CheckDiscordId(discordId string) string {
	q, err := db.Database.Sql.Query(`SELECT address from backend where discordid=$1`, discordId)
	if err != nil {
		panic(err.Error())
	}

	defer q.Close()
	var address string

	if q.Next() {

		if err := q.Scan(&address); err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Println("Db address", address)
	}
	fmt.Println(address)
	return address

}
func (db *Db) GetAllNftFromAddress(address string) []string {
	q, err := db.Database.Sql.Query(`SELECT nftid,owner_address FROM nfts WHERE owner_address=$1`, address)
	if err != nil {
		log.Fatalf("%s", err.Error())
		panic(err.Error())
	}
	defer q.Close()
	var nfts []string
	for q.Next() {
		var id string

		if err := q.Scan(&id); err != nil {
			log.Fatalf("%s", err.Error())
			panic(err.Error())
		}
		nfts = append(nfts, id)

		// nftArray = append(nftArray, id)
		// fmt.Println(nftArray)
	}
	return nfts
}
func (db *Db) GetGatedList() []string {
	q, err := db.Database.Sql.Query(`SELECT collection_id,roleid,role_name FROM gatedList`)
	if err != nil {
		log.Fatalf("%s", err.Error())
		panic(err.Error())
	}
	defer q.Close()
	var communityArray []string
	for q.Next() {
		var collection_id, roleid, role_name string
		if err := q.Scan(&collection_id, &roleid, &role_name); err != nil {
			log.Fatalf("%s", err.Error())
			panic(err.Error())
		}
		communityArray = append(communityArray, collection_id)

	}

	fmt.Println("Community Array", communityArray)
	return communityArray
}

// func (db *Db) GetDiscordId(address string) {
// 	q, err := db.Database.Sql.Query(`SELECT `)
// }
