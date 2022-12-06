CREATE TABLE community
(
    name TEXT DEFAULT  '',
    creator TEXT DEFAULT '',
    description TEXT DEFAULT  '',
    preview_url TEXT DEFAULT '',
    id TEXT NOT NULL PRIMARY KEY UNIQUE
);
CREATE  INDEX  community_creator_index ON community(creator);