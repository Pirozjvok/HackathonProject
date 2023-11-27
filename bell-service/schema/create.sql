create table bells
(
    id serial primary key,
    status          varchar(10),
    calldate        timestamp,
    operatorfio    varchar,
    clientphone     varchar,
    contactaudiourl varchar,
    statusbell      varchar
);