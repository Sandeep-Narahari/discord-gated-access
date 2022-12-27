package types

import "time"

type Denom struct {
	Id          string
	Symbol      string
	Name        string
	Creator     string
	Description string
	Preview_URI string
	Height      int64
	CreatedAt   time.Time
	Source      string
	Soure_URL   string
	Featured    bool
}

func NewDenom(id, symbol, name, creator, desc, prev string, height int64, created_at time.Time, source, source_url string, featured bool) Denom {
	return Denom{
		Id:          id,
		Symbol:      symbol,
		Name:        name,
		Creator:     creator,
		Description: desc,
		Preview_URI: prev,
		Height:      height,
		CreatedAt:   created_at,
		Source:      source,
		Soure_URL:   source_url,
		Featured:    featured,
	}
}
