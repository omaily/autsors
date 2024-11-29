CREATE TABLE account (
    id SERIAL PRIMARY KEY,
    uuid TEXT NOT NULL,
    name TEXT NOT NULL, 
    amount INT
);

INSERT INTO account (uuid, iname, amount) 
    VALUES ('00000000-0000-0000-0000-000000000001', 'Алеша', 52);
    
INSERT INTO account (uuid, iname, amount) 
    VALUES ('00000000-0000-0000-0000-000000000002', 'Степаша', 322);

INSERT INTO account (uuid, iname, amount) 
    VALUES ('66666666-6666-6666-6666-666666666666', 'Олег', 666);

INSERT INTO account (uuid, iname, amount) 
    VALUES ('00000000-0000-0000-0000-000000000003', 'Никита', 3141592);