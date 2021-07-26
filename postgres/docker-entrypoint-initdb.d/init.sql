create table author
(
    id   integer,
    name varchar(99)
);

create table book
(
    id        integer,
    title     varchar(99),
    price     integer,
    author_id integer
);

insert into author (id, name)
values (101, 'Mat Ryer');
insert into author (id, name)
values (102, 'Katherine Cox-Buday');
insert into author (id, name)
values (103, 'Thorsten Ball');

insert into book (id, title, price, author_id)
values (301, 'Go言語によるWebアプリケーション開発', 3520, 101);
insert into book (id, title, price, author_id)
values (302, 'Go言語による並行処理', 3080, 102);
insert into book (id, title, price, author_id)
values (303, 'Go言語でつくるインタプリタ', 3740, 103);
