package database

import "github.com/AutonomyNetwork/iam/types"

func (db Db) SaveCommunity(community types.Community) error {
	stmt := `
INSERT INTO community (name, creator, description, preview_url, id )
VALUES ($1, $2, $3, $4, $5)`
	
	_, err := db.Sqlx.Exec(stmt, community.Name, community.Creator, community.Description, community.PreviewURI, community.Id)
	return err
}
