CREATE DATABASE IF NOT EXISTS testdb;
CREATE USER test IDENTIFIED WITH plaintext_password BY 'test';
GRANT ALL ON testdb.* TO test;