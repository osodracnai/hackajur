CREATE TABLE IF NOT EXISTS debtors (
    id uuid PRIMARY KEY,
    fiscal_document TEXT,
    name TEXT,
    email TEXT,
    type_of_debtor TEXT,
    postal_code TEXT,
    city TEXT,
    uf TEXT,
    street TEXT,
    number TEXT,
    complement TEXT
);

CREATE TABLE IF NOT EXISTS proposes (
    id uuid PRIMARY KEY,
    debt VARCHAR,
    date TIMESTAMP,
    status jsonb,
    propose_value BIGINT,
    expiration_date BIGINT,
    payment_deadline BIGINT,
    payments jsonb,
    comunication jsonb
);


CREATE TABLE IF NOT EXISTS debts (
    id uuid PRIMARY KEY,
    debtor VARCHAR,
    origin TEXT,
    document_id TEXT,
    expiration_date TIMESTAMP,
    original_value BIGINT,
    fee jsonb,
    interest jsonb,
    other_charges jsonb,
    present_value BIGINT,
    collateral VARCHAR,
    correction jsonb
);

