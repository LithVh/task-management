ALTER TABLE  member RENAME TO app_user;
alter table app_user drop column leader;

create table project (
    ID bigserial primary key,
    User_id UUID references app_user(ID) not null,
    Name varchar(255) not null,
    Description varchar(255)
);

create table task (
    ID bigserial primary key,
    Project_id bigint references project(ID) not null,
    Title varchar(255) not null,
    Status varchar(20) not null,
    Priority varchar(20)
);

