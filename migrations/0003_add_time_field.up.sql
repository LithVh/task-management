alter table project add column created_at timestamp not null;
alter table project add column updated_at timestamp not null;
alter table task add column created_at timestamp not null;
alter table task add column updated_at_at timestamp not null;