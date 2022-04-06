(MYSQL)

create database alura_loja;

use alura_loja;

create table produtos(
    id int auto_increment primary key,
    nome varchar(100),
    descricao varchar(100),
    preco decimal,
    quantidade int
)ENGINE=INNODB;

insert into produtos(nome, descricao, preco, quantidade)
values ('Bermuda', 'Para corrida', 52, 5),
('Creme dental', 'Protege 12 horas', 7.99, 45);