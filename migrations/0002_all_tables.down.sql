ALTER TABLE  app_user RENAME TO member;
alter table member add column leader UUID;

drop table project;
drop table task;