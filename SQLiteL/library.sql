CREATE TABLE books (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    isbn VARCHAR(32) NOT NULL DEFAULT "",
    author_id INTEGER NOT NULL DEFAULT 0,
    title VARCHAR(256) NOT NULL DEFAULT "",
    annotation TEXT NOT NULL DEFAULT ""
);
CREATE INDEX books_authors ON books (author_id);

CREATE TABLE authors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(128) NOT NULL DEFAULT ""
);     
CREATE INDEX authors_name ON authors (name);

CREATE TABLE themes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(128) NOT NULL DEFAULT ""
);     
CREATE INDEX themes_name ON themes (name);

CREATE TABLE booktheme (
    book_id INTEGER NOT NULL DEFAULT 0,
    theme_id INTEGER NOT NULL DEFAULT 0
);     
CREATE INDEX booktheme_book ON booktheme (book_id);
CREATE INDEX booktheme_theme ON booktheme (theme_id);

CREATE TABLE readers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(128) NOT NULL DEFAULT "",
    phone VARCHAR(32) NOT NULL DEFAULT "",
    email VARCHAR(64) NOT NULL DEFAULT "",
    birthday CHAR(8) NOT NULL DEFAULT ""
);     
CREATE INDEX readers_name ON readers (name);

CREATE TABLE actions (
    book_id INTEGER NOT NULL DEFAULT 0,
    reader_id INTEGER NOT NULL DEFAULT 0,
    date CHAR(8) NOT NULL DEFAULT "",
    return CHAR(8) NOT NULL DEFAULT ""
);     
CREATE INDEX actions_bookreader ON actions (book_id, reader_id);