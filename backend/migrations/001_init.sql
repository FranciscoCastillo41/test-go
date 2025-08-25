create table if not exists widgets (
    id        uuid primary key,
    name      text not null,
    price     numeric(12,2) not null check (price >= 0),
    created_at timestamptz not null default now()
);
