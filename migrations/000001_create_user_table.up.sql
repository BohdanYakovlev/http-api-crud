Create TABLE  users(
    id serial NOT NULL PRIMARY KEY,
    username varchar(255) not null,
    phone varchar(255) not null
);

alter table users
add constraint uniqueName unique (username)