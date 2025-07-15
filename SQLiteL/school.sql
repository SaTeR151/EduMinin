create table subject(
    id integer primary key autoincrement,
    name varchar(256) not null default ""
);
create index subject_name on subject (name);

create table teachers(
    id integer primary key autoincrement,
    fio varchar(256) not null default "",
    phone varchar(256) not null default ""
);
create index teachers_fio on teachers (fio);

create table teacherssubject(
    teacher_id integer not null default 0,
    subject_id integer not null default 0
);
create index teacherssubject_teacher on teacherssubject (teacher_id);
create index teacherssubject_subject on teacherssubject (subject_id);

create table students(
    id integer primary key autoincrement,
    fio varchar(256) not null default "",
    parent_phone varchar(256) not null default "",
    birthday char(8) not null default "",
    adress varchar(256) not null default "",
    foreign key (classId) references classes (id)
);
create index students_classfio on students (class_id, fio);


create table classes(
    id integer primary key autoincrement,
    years integer not null default 0,
    letter char(1) not null default ""
);

create table teachersclasses(
    teacher_id integer not null default 0,
    class_id integer not null default 0
);
create index teachersclasses_teacher on teachersclasses (teacher_id);
create index teachersclasses_class on teachersclasses (class_id);
