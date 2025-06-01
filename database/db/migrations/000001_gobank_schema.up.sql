CREATE TYPE public.transferstatus AS ENUM (
  'created',
  'processed',
  'failure',
  'success'
);

CREATE TABLE public.accounts (
  id uuid PRIMARY KEY NOT NULL,
  owner text NOT NULL,
  balance float NOT NULL DEFAULT 0,
  currency text NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE public.transfers (
  id uuid PRIMARY KEY NOT NULL,
  from_acc uuid NOT NULL,
  to_acc uuid NOT NULL,
  amount float NOT NULL,
  status transferstatus NOT NULL DEFAULT 'created',
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  FOREIGN KEY (from_acc) REFERENCES accounts (id),
  FOREIGN KEY (to_acc) REFERENCES accounts (id)
);

CREATE TABLE public.entries (
  id uuid PRIMARY KEY NOT NULL,
  account_id uuid NOT NULL,
  transfer_id uuid NOT NULL,
  amount float NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  FOREIGN KEY (account_id) REFERENCES accounts (id),
  FOREIGN KEY (transfer_id) REFERENCES transfers (id)
);

CREATE INDEX owner_idx ON accounts (owner);
CREATE INDEX accs_trns_idx ON entries (account_id, transfer_id);
CREATE INDEX from_to_idx ON transfers (from_acc, to_acc);
CREATE INDEX status_idx ON transfers (status);
