CREATE TABLE books (
   ID UUID PRIMARY KEY UNIQUE NOT NULL DEFAULT gen_random_uuid (),
   AUTHOR varchar,
   TITLE varchar,
   PAGEAMOUNT INTEGER
);

INSERT INTO books(author, title, pageamount)
VALUES ('Dostoevsky', 'Besy', 300);

INSERT INTO books(author, title, pageamount)
VALUES ('Strugatskiy', 'Niichavo', 500);

INSERT INTO books(author, title, pageamount)
VALUES ('Joyce', 'Ulysses', 2000);