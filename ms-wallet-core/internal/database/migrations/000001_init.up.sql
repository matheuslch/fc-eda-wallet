CREATE TABLE clients (
    id varchar(255) NOT NULL PRIMARY KEY, 
    name varchar(255), 
    email varchar(255), 
    created_at datetime
);
    
CREATE TABLE accounts (
    id varchar(255) NOT NULL PRIMARY KEY, 
    client_id varchar(255), 
    balance float, 
    created_at datetime
);
    
CREATE TABLE transactions (
    id varchar(255) NOT NULL PRIMARY KEY, 
    account_id_from varchar(255), 
    account_id_to varchar(255), 
    amount float, 
    created_at datetime
);