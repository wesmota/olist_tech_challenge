CREATE TABLE author (
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL UNIQUE
);

CREATE TABLE book (
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
    edition TEXT NULL,
    publication_year INTEGER(4)
);

CREATE TABLE book_author (
	id INTEGER PRIMARY KEY,
	id_author INTEGER NOT NULL,
    id_book INTEGER NOT NULL,
    FOREIGN KEY (id_author) 
      REFERENCES author (id) 
         ON DELETE CASCADE 
         ON UPDATE NO ACTION,
    FOREIGN KEY (id_book) 
      REFERENCES book (id) 
         ON DELETE CASCADE 
         ON UPDATE NO ACTION
);



