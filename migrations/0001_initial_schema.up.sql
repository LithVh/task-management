create table member (
    ID  UUID    primary key default gen_random_uuid(),
    name    varchar(255) not null,
    email   varchar(255) unique not null,
    password_hash varchar(255) not null,
    created_at timestamp,
    leader UUID
)
