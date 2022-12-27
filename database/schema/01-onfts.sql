CREATE TABLE denom
(
    id TEXT NOT NULL PRIMARY KEY UNIQUE,
    symbol TEXT DEFAULT '',
    name TEXT DEFAULT '',
    creator TEXT DEFAULT '',
    description TEXT DEFAULT '',
    preview_uri TEXT DEFAULT '',
    height BIGINT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    source TEXT DEFAULT '',
    source_url TEXT DEFAULT  '',
    featured BOOLEAN NOT NULL DEFAULT FALSE,
    UNIQUE (creator, created_at)
);
CREATE INDEX denom_creator_index ON denom(creator);

CREATE TABLE nfts(
    id TEXT NOT NULL PRIMARY KEY UNIQUE,
    name TEXT NOT NULL  DEFAULT '',
    description TEXT NOT NULL  DEFAULT '',
    preview_uri TEXT NOT NULL DEFAULT '',
    owner TEXT NOT NULL  DEFAULT '',
    denom_id TEXT NOT NULL REFERENCES denom(id),
    price TEXT NOT NULL  DEFAULT '',
    source_url TEXT DEFAULT '',
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    height BIGINT NOT NULL DEFAULT 0
)
