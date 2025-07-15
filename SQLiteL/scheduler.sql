CREATE TABLE scheduler (
    "id" INTEGER NOT NULL PRIMARY KEY,
    "date" integer NOT NULL,
    "title" VARCHAR(100) NOT NULL DEFAULT "",
    "comment" VARCHAR(250),
    "repeat" VARCHAR(120) NOT NULL
);
create index scheduler_data on scheduler (data);
