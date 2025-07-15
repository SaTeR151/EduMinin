CREATE TABLE predpr(
    id_pr SMALLINT NOT NULL,
    nazvp VARCHAR(20) NOT NULL,
    city VARCHAR(15) NOT NULL,
    header VARCHAR (30),
    phone VARCHAR(12),
    CONSTRAINT pkpr PRIMARY KEY (id_pr)
);

CREATE TABLE agents(
    id_ag SMALLINT NOT NULL,
    fam VARCHAR(30) NOT NULL,
    name VARCHAR(15) NOT NULL,
    otch VARCHAR(25) NULL,
    dater DATE NOT NULL,
    Payment DECIMAL(8,2) NOT NULL,
    id_pr SMALLINT NOT NULL,
    CONSTRAINT pkagent PRIMARY KEY (id_ag),
    CONSTRAINT FKAgent  FOREIGN KEY (id_pr) REFERENCES predpr(id_pr)
);

CREATE TABLE tovary(
    id_tov SMALLINT NOT NULL,
    nametov VARCHAR(35) NOT NULL,
    cenaopt DECIMAL NOT NULL,
    cenarozn DECIMAL NOT NULL,
    CONSTRAINT pktov PRIMARY KEY (id_tov)
);

CREATE TABLE zakaz(
    zakazano SMALLINT NOT NULL,
    prodano SMALLINT NOT NULL,
    date_zak DATE NOT NULL,
    id_ag SMALLINT NOT NULL,
    id_tov SMALLINT NOT NULL,
    CONSTRAINT pkz PRIMARY KEY (id_ag, id_tov, date_zak),
    CONSTRAINT fkag FOREIGN KEY (id_ag) REFERENCES agents(id_ag),
    CONSTRAINT fktov FOREIGN KEY (id_tov) REFERENCES tovary(id_tov)
);

INSERT INTO predpr(id_pr,nazvp,city,header,phone) 
VALUES 
	(1,'АО Интерсервис', 'Новосибирск', 'Иванов ИИ', '57-34-54'),
	(2,'АО Компьютер', 'Одесса', 'Петров ПП', '38-32-24'),
	(3,'ТОО Интеркомбанк', 'Одесса', 'Сидоров СС', '38-45-31'),
	(4,'Банк Программ', 'Новгород', 'Олегов ОО', '32-45-31'),
	(5,'ИЧП Элвис', 'Новгород', 'Сергеев СС', '81-60-17'),
	(6,'Bel-IBM', 'Новгород', 'Орлов РЛ', '32-45-31'),
	(7,'СП КомпСофт', 'Новочеркасск', 'Грибов ГГ', '35-31-24'),
	(8,'Bel-IBM', 'Н.Новгород', 'Ильин ИИ', '81-66-25'),
	(9,'СП АВМ-Сервис', 'Новосибирск', 'Алепов АА', '53-12-45'),
	(10,'Стингер', 'Новосибирск', 'Полинов ПП', '83-12-12'),
	(11,'Комп-Граф', 'Н.Новгород', 'Тамарова ТТ', '38-32-45'),
	(12,'Айрат', 'Н.Новгород', 'Комаров КК', '50-27-19'),
	(13,'Офис-Сервис', 'Новосибирск', 'Ульянов УУ', '33-12-46'),
	(14,'АО РосКОМ', 'Новочеркасск', 'Галинова ГГ', '16-01-77'),
	(15,'АО Фигон', 'Новочеркасск', 'Арбатова МВ', '35-31-24');

INSERT INTO agents(id_ag,fam,name,otch,dater, Payment,id_pr)
VALUES
(1,'Алексеенков', 'Сергей', 'Иванович','1970-11-21', 12390, 3),
	(2,'Ткаченко', 'Марина', 'Михайловна', '1967-10-25', 13890, 7),
	(3,'Морозов', 'Юрий', 'Анатольевич', '1956-12-23', 14500, 6),
	(4,'Петров', 'Михаил', 'Николаевич', '1956-12-15', 15620, 4),
	(5,'Ильин', 'Филипп', 'Петрович', '1972-03-02', 12360, 11),
	(6,'Петров', 'Михаил', 'Николаевич', '1948-08-18', 14510, 15),
	(7,'Зайцева', 'Людмила', 'Николаевна', '1974-04-03', 13470, 8),
	(8,'Загревский', 'Андрей', 'Анатольевич', '17.05.62', 22200, 2),
	(9,'Смесов', 'Филипп', 'Петрович', '1959-02-06', 13650, 5),
	(10,'Сухарев', 'Андрей', 'Анатольевич', '1968-03-29', 24160, 2),
	(11,'Зайцева', 'Дарья', 'Михайловна', '1974-07-15', 15190, 9),
	(12,'Фролов', 'Виктор', 'Михайлович', '1957-07-27', 22500, 10),
	(13,'Свешников', 'Сергей', 'Иванович', '1989-03-30', 13400, 12),
	(14,'Уханов', 'Михаил', 'Сергеевич', '1998-11-28', 14590, 15),
	(15,'Морозов', 'Андрей', 'Михайлович', '1983-09-17', 23370, 1),
	(16,'Райский', 'Михаил', 'Сергеевич', '1964-06-22', 13520, 5);

INSERT INTO tovary(id_tov, nametov,cenaopt, cenarozn)
VALUES
(1,'CP 35 FX', 1000, 1200),
(2,'CP 40 MC', 1400, 1580), 
(3,'CP 55FX', 1800, 1950), 
(4,'CP 55 LS', 1390, 1500), 
(5,'Вист 133', 2800, 3000), 
(6,'Andy 586', 3900, 4500), 
(7,'Borland C++', 5000, 5200), 
(8,'Delphi', 2600, 2800), 
(9,'Microsoft Office', 4000, 4600), 
(10,'Microsoft Office Pro', 5000, 5900);

INSERT INTO zakaz(id_ag,id_tov,zakazano,prodano,date_zak)
VALUES
	(3, 2, 23, 21, '2020-02-01'),
	(5, 4, 12, 12, '2020-03-01'),
	(9, 7, 10, 10, '2020-02-01'),
	(5, 6, 8, 7, '2020-02-05'),
	(14, 7, 15, 15, '2020-02-13'),
	(8, 8, 15, 15, '2020-04-12'),
	(9, 5, 8, 8, '2020-09-23'),
	(14, 10, 11, 9, '2020-08-16'),
	(8, 6, 24, 24, '2020-07-27'),
	(6, 4, 10, 10, '2020-04-18'),
	(11, 3, 9, 7, '2020-06-11'),
	(3, 2, 5, 4, '2020-03-19'),
	(10, 5, 13, 11, '2020-09-08'),
	(8, 4, 22, 22, '2020-09-10'),
	(5, 8, 5, 5, '2020-04-17');


--- 1
SELECT fam, name, otch, nametov, zakazano, date_zak FROM agents a, tovary t, zakaz z WHERE a.id_ag=z.id_ag AND t.id_tov=z.id_tov;

--- 2 
SELECT fam, name, otch, zakazano FROM agents LEFT JOIN zakaz ON agents.id_ag=zakaz.id_ag;

--- 3
SELECT nazvp, avg(cenaopt) avg_opt, max(cenaopt) max_opt, min(cenaopt) min_opt, avg(cenarozn) avg_roz, max(cenarozn) max_roz, min(cenarozn) min_roz FROM zakaz z, agents a, tovary t, predpr p WHERE z.id_ag=a.id_ag AND z.id_tov=t.id_tov AND a.id_pr=p.id_pr GROUP BY nazvp;

--- 4
SELECT zakazano, prodano, date_zak, city, nazvp FROM zakaz z, agents a, predpr p where z.id_ag=a.id_ag and a.id_pr=p.id_pr and (city LIKE 'Новосибирск' OR city LIKE 'Одесса' or nazvp LIKE 'АО%');

--- 5 
SELECT fam, name, otch, (prodano * cenarozn - zakazano * cenaopt) prib from zakaz z, agents a, tovary t where z.id_ag=a.id_ag and z.id_tov=t.id_tov GROUP BY a.id_ag;

--- 6 
SELECT nametov, (zakazano-prodano) ostatok from zakaz z, tovary t where z.id_tov=t.id_tov GROUP BY nametov;

--- 7 
SELECT fam, name, otch, sum(prodano * cenarozn - zakazano * cenaopt) prib from zakaz z, agents a, tovary t where z.id_ag=a.id_ag and z.id_tov=t.id_tov GROUP BY a.id_ag ORDER BY prib LIMIT 1;

--- 8
select fam, count(*) count from zakaz JOIN agents on agents.id_ag=zakaz.id_ag group by fam order by count desc limit 1;  

--- 9 
select nametov from zakaz join tovary on zakaz.id_tov=tovary.id_tov group by nametov order by sum(zakazano) desc limit 1; 

--- 10
select MONTH(date_zak) month from zakaz group by month(date_zak) order by month desc limit 1;

--- 11
select fam, prodano*cenarozn- zakazano*cenaopt as total, iif(total<0,'Заказ убыточен',iif(total<1500,0.1*total,iif(total<3000,0.12*total, 0.15*total))) as cost from zakaz z, agents a, tovary t where a.id_ag=z.id_ag and z.id_tov=t.id_tov;

--- 12
select fam, int((now()-dater)/362.25) as vozrast, iif(vozrast MOD 10=0, 'Поздравить с круглой датой и премию дать!', ' ') as prem from agents;

--- 13
update tovary set cenaopt=cenaopt*1.1 where 20<(select sum(zakazano) as zakazi from zakaz group by id_tov);

--- 14
select distinct a2.fam as agent, z2.date_zak as data, a1.fam as ilin, z1.date_zak as dateIlin from zakaz z1, zakaz z2, agents a1, agents a2 where a1.fam='Ильин' and a1.id_ag=z1.id_ag and a2.id_ag=z2.id_ag and a1.id_ag<>a2.id_ag and month(z1.date_zak)=month(z2.date_zak) order by z1.date_zak;