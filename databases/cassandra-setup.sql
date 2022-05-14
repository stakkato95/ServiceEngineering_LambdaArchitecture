-- keyspace == database
CREATE KEYSPACE analytics
WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 1};

-- SELECT * FROM system_schema.keyspaces;

USE analytics;

DROP TABLE IF EXISTS user;

-- 1 cassandra supported types
-- https://cassandra.apache.org/doc/latest/cassandra/cql/types.html?msclkid=f96cf2c2d05711ec9dee2ac82ad158e0
-- 2 role of keys
-- https://www.geeksforgeeks.org/role-of-keys-in-cassandra/?msclkid=98eb2301d05711ecb7bc89f79d5a8405
-- 3 gocql example
-- https://code.tutsplus.com/tutorials/getting-started-with-cassandra-understanding-various-cql-data-types--cms-28110?msclkid=37153042d05911ec90a62813a782239a
CREATE TABLE user(
    id UUID,
    time timestamp,
    user_count int,
    PRIMARY KEY (id)
);

SELECT * FROM user;
-- TRUNCATE TABLE user;