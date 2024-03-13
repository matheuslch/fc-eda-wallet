CREATE TABLE IF NOT EXISTS balances (
    id varchar(255) PRIMARY KEY,
    account_id varchar(255),
    balance float,
    updated_at datetime
);