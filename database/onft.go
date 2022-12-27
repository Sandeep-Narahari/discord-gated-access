package database

import (
	"fmt"
	
	"github.com/AutonomyNetwork/iam/types"
)

func (db *Db) SaveDenomData(denom types.Denom) error {
	return db.SaveDenomsData([]types.Denom{denom})
}

func (db Db) SaveDenomsData(denoms []types.Denom) error {
	if len(denoms) == 0 {
		return nil
	}
	
	denomQuery := `
INSERT INTO denom (id, symbol, name, creator, description, preview_uri,height,created_at, source,source_url, featured) VALUES
`
	var denomParams []interface{}
	
	for i, denom := range denoms {
		denomQuery += fmt.Sprintf("($%d),($%d),($%d),($%d),($%d),($%d),($%d),($%d),($%d),($%d),($%d),", i+1, i+2, i+3, i+4, i+5, i+6, i+7, i+8, i+9, i+10, i+11)
		denomParams = append(denomParams, denom.Id, denom.Symbol, denom.Name, denom.Creator, denom.Description, denom.Preview_URI, denom.Height, denom.CreatedAt, denom.Source, denom.Soure_URL, denom.Featured)
	}
	
	denomQuery = denomQuery[:len(denomQuery)-1]
	denomQuery += `
ON CONFLICT (id) DO UPDATE
SET id = excluded.id,
	symbol = excluded.symbol,
	name = excluded.name,
	creator = excluded.creator,
	description = excluded.description,
	preview_uri = excluded.preview_uri,
	height = excluded.height,
    created_at = excluded.height,
    source = excluded.source,
    source_url = excluded.source_url,
    featured = excluded.featured
WHERE denom.height <= excluded.height`
	
	_, err := db.Sqlx.Exec(denomQuery, denomParams...)
	if err != nil {
		return fmt.Errorf("error while storing denoms infos: %s", err)
	}
	return err
}

func (db *Db) GetDenoms() ([]types.Denom, error) {
	sqlStmt := `
SELECT DISTINCT ON (denom.id)
	denom.id,
	denom.symbol,
	denom.name,
	denom.creator,
	denom.description,
	denom.preview_uri,
	denom.height,
	denom.created_at,
	denom.source,
	denom.source_url,
	denom.featured
FROM denom
ORDER BY denom.created_at`
	
	var rows []types.Denom
	err := db.Sqlx.Select(&rows, sqlStmt) // TODO : check the types marshal and unmarshal
	if err != nil {
		return nil, err
	}
	
	return rows, nil
}
