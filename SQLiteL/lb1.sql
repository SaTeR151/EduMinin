CREATE TABLE adresbook (
  `Chelid` int(11) NOT NULL,
  `Fam` varchar(30) NOT NULL,
  `Name` varchar(15) NOT NULL,
  `Otch` varchar(30) DEFAULT NULL,
  `Town` varchar(20) NOT NULL,
  `dateR` date NOT NULL,
  `work` varchar(25) DEFAULT NULL,
  PRIMARY KEY (`Chelid`)
);

ALTER TABLE adresbook ADD pol VARCHAR(2);
ALTER TABLE adresbook ADD oklad NUMERIC(8,2);

INSERT INTO `adresbook` (`Chelid`, `Fam`, `Name`, `Otch`, `Town`, `dateR`, `work`, `pol`, `oklad`) VALUES
(1, 'Ильин', 'Роман', 'Алексеевич', 'Ярославль', '1990-02-11', 'ЗАО Яринтерком', 'м.', '25000.00'),
(2, 'Полякова', 'Тамара', 'Игоревна', 'Москва', '2001-10-23', 'ТОО Звезда', 'ж.', '19000.00'),
(3, 'Петров', 'Сергей', 'Ильич', 'Москва', '1987-09-18', '', 'м.', '28000.00'),
(4, 'Классен', 'Роланд', '', 'Нижний Новгород', '1989-10-29', 'ЗАО Тер-маль', 'м.', '19000.00'),
(5, 'Ермилова', 'Тамара', 'Дмитриевна', 'Москва', '2003-10-30', 'ООО Ольха', 'ж.', '22000.00'),
(6, 'Николаев', 'Олег', 'Иванович', 'Нижний Новгород', '1970-08-16', 'ТОО Лад', 'м.', '23000.00'),
(7, 'Горелова', 'Маргарита', 'Романовна', 'Ярославль', '2004-06-26', '', 'ж.', '10000.00'),
(8, 'Рокунов', 'Илья', 'Петрович', 'Рязань', '1978-10-28', 'воен.часть', 'м.', '27000.00'),
(9, 'Ермакова', 'Татьяна', 'Алексеевна', 'Москва', '2000-11-20', 'ООО Апрель', 'ж.', '22000.00');

--- 1
select Fam, Name, dateR from adresbook;

--- 2
select Fam, Name, Otch from adresbook where Town like 'М%'; 

--- 3
select * from adresbook where pol='м.'and Town='Москва';

--- 4 
select Town, avg(oklad) oklad from adresbook group by Town;   

--- 5
select * from adresbook order by Town; 