create table useraccount (
    id serial primary key,
    email varchar unique,
    fullname varchar,
    password varchar,
    phone_number varchar unique,
    sex int,
    biography varchar,
    location varchar,
    date_of_birth date,
    profile_photo varchar,
    created_at timestamp default now(),
    updated_at timestamp default null
);

create index concurrently useraccountemail on useraccount(email);