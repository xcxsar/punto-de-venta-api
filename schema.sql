create sequence hobby_id_seq
    as integer;

alter sequence hobby_id_seq owner to punto_de_venta_admin;

create table hobby
(
    id    integer default nextval('hobby_id_seq'::regclass) not null
        constraint hobby_pk
            primary key,
    emoji text,
    name  text
);

alter table hobby
    owner to punto_de_venta_admin;

alter sequence hobbie_id_seq owned by hobby.id;

create table completion
(
    date_completed date,
    hobby_id       bigint
        constraint completion_hobby_id_fk
            references hobby
            on delete cascade
);

alter table completion
    owner to punto_de_venta_admin;

