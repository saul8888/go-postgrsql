--CREATE TABLE public.persons (
CREATE TABLE persons (
    --id SERIAL PRIMARY KEY,
    id varchar(255) NOT NULL PRIMARY KEY,
    firstname varchar(255),
    email varchar(255),
    pass varchar(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    update_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);