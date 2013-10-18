package main

func create_database() error {
	return nil
}

var INITIAL_DB = `
    /*
        URL entries
     */
    CREATE TABLE shortly_urls(
        id BIGSERIAL PRIMARY KEY,
        code TEXT,
        url TEXT NOT NULL,

        checked BOOLEAN,
        author INET,
        created TIMESTAMP default statement_timestamp()
    );


    /*
        Domains that we won't create short-links for. This doesn't include the
        current instance, as we don't know it yet, but this is handled in code.
     */
    CREATE TABLE shortly_disallowed_domains(
        id BIGSERIAL PRIMARY KEY,
        domain TEXT NOT NULL
    );


    /*
        Resolution requests, we create one of these each time a link is resolved
     */
    CREATE TABLE shortly_record(
        id BIGSERIAL PRIMARY KEY,
        url_id BIGINT references shortly_urls(id),

        request_ip INET,
        referer TEXT,
        useragent TEXT,
        headers JSON,

        created TIMESTAMP default statement_timestamp()
    );

    /*
        Indexeses
     */
    CREATE INDEX idx_shortly_urls_code ON shortly_urls(code);
    CREATE INDEX idx_shortly_disallowed_domains_domain ON shortly_disallowed_domains(domain);
    /* TODO: Indexes on JSON types ... */

    /*
      Initial data for known url shorteners to make sure nobody pastes a URL to these.
     */
    INSERT INTO shortly_disallowed_domains(domain) VALUES('bit.ly');
    INSERT INTO shortly_disallowed_domains(domain) VALUES('t.co');
`
