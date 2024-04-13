DROP TABLE IF EXISTS banners CASCADE;
DROP TABLE IF EXISTS banner_feature CASCADE;
DROP TABLE IF EXISTS banner_tag CASCADE;

CREATE SEQUENCE IF NOT EXISTS banners_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1
    OWNED BY banners.id;

CREATE TABLE IF NOT EXISTS banners(
    id integer NOT NULL DEFAULT nextval('banners_id_seq'::regclass),
    content jsonb NOT NULL,
    is_active boolean NOT NULL DEFAULT true,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT banners_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS banner_tag(
	banner_id integer NOT NULL,
    tag_id integer NOT NULL,
    CONSTRAINT banner_tag_pkey PRIMARY KEY (banner_id, tag_id),
    CONSTRAINT banner_tag_banner_id_fkey FOREIGN KEY (banner_id)
        REFERENCES banners (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS banner_feature(
	banner_id integer NOT NULL,
    feature_id integer NOT NULL,
    CONSTRAINT banner_feature_pkey PRIMARY KEY (banner_id, feature_id),
    CONSTRAINT banner_feature_banner_id_fkey FOREIGN KEY (banner_id)
        REFERENCES banners (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
