CREATE KEYSPACE IF NOT EXISTS scylla_talk_ks WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

CREATE TABLE IF NOT EXISTS scylla_talk_ks.user (
    id uuid,
    name text,
    email text,
    about text,
    properties map<text, text>,
    PRIMARY KEY (id)
);
