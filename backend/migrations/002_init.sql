-- Users table to store user data tied to Supabase Auth
create table if not exists users (
    id uuid primary key default gen_random_uuid(),
    auth_user_id uuid not null unique, -- Supabase Auth user ID (from JWT 'sub' claim)
    email text not null unique,
    full_name text,
    avatar_url text,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

-- Index for fast lookups by auth_user_id
create index if not exists idx_users_auth_user_id on users(auth_user_id);

-- Index for email lookups
create index if not exists idx_users_email on users(email);

-- Function to update updated_at timestamp
create or replace function update_updated_at_column()
returns trigger as $$
begin
    new.updated_at = now();
    return new;
end;
$$ language plpgsql;

-- Trigger to automatically update updated_at
drop trigger if exists update_users_updated_at on users;
create trigger update_users_updated_at
    before update on users
    for each row
    execute function update_updated_at_column();
