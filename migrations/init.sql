create database "note-service-api";
create user "shem" with password '12345678';
grant all privileges on database "note-service-api" to "shem";

\c note-service-api shem

create table notes(
    id serial primary key,
    user_id bigint not null,
    classroom_id bigint not null,
    document_id bigint not null
);

alter table notes owner to shem;