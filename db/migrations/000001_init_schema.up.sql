CREATE TABLE IF NOT EXISTS public.account
(
    id character varying(255) COLLATE pg_catalog."default" NOT NULL,
    name character varying(255) COLLATE pg_catalog."default",
    created_date timestamp without time zone,
    modified_date timestamp without time zone,
    CONSTRAINT account_pkey PRIMARY KEY (id)
    )
                           WITH (
                               OIDS = FALSE
                               )
    TABLESPACE pg_default;

-- Index: index_account_id

CREATE INDEX index_account_id
    ON public.account USING btree
    (id COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;

-- Table: public.currency

CREATE TABLE IF NOT EXISTS public.currency
(
    id character varying(5) COLLATE pg_catalog."default" NOT NULL,
    "desc" text COLLATE pg_catalog."default",
    CONSTRAINT currency_pkey PRIMARY KEY (id)
    )
    WITH (
        OIDS = FALSE
        )
    TABLESPACE pg_default;

insert into public.currency values ('USD', 'US Dollar')
on conflict (id) do nothing;

-- Table: public.account_balance

CREATE TABLE IF NOT EXISTS public.account_balance
(
    id character varying(255) COLLATE pg_catalog."default" NOT NULL,
    balance bigint,
    currency character varying(5) COLLATE pg_catalog."default",
    CONSTRAINT account_balance_pkey PRIMARY KEY (id),
    CONSTRAINT currency_pkey FOREIGN KEY (currency)
    REFERENCES public.currency (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID,
    CONSTRAINT balance_check CHECK (balance >= 0) NOT VALID
    )
    WITH (
        OIDS = FALSE
        )
    TABLESPACE pg_default;
-- Index: index_account_balance_id_currency

CREATE INDEX index_account_balance_id_currency
    ON public.account_balance USING btree
    (id COLLATE pg_catalog."default" ASC NULLS LAST, currency COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: index_account_balance_idx

CREATE INDEX index_account_balance_idx
    ON public.account_balance USING btree
    (id COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;

-- Table: public.payment_data

CREATE TABLE IF NOT EXISTS public.payment_data
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    debit bigint,
    credit bigint,
    account_id character varying(255) COLLATE pg_catalog."default",
    last_balance bigint,
    current_balance bigint,
    direction character varying(11) COLLATE pg_catalog."default",
    to_account_id character varying(255) COLLATE pg_catalog."default",
    from_account_id character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT id_primary_key PRIMARY KEY (id),
    CONSTRAINT account_fkey FOREIGN KEY (account_id)
    REFERENCES public.account (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION,
    CONSTRAINT from_account_fkey FOREIGN KEY (from_account_id)
    REFERENCES public.account (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID,
    CONSTRAINT to_account_fkey FOREIGN KEY (to_account_id)
    REFERENCES public.account (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID,
    CONSTRAINT last_balance_check CHECK (last_balance >= 0),
    CONSTRAINT current_balance_check CHECK (current_balance >= 0),
    CONSTRAINT debit_check CHECK (debit >= 0),
    CONSTRAINT credit_check CHECK (credit >= 0)
    )
    WITH (
        OIDS = FALSE
        )
    TABLESPACE pg_default;