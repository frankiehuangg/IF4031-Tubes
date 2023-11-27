CREATE TABLE events (
    event_id SERIAL PRIMARY KEY,
    event_name VARCHAR(256) UNIQUE NOT NULL,
    event_date TIMESTAMP NOT NULL,
    total_seat INT NOT NULL
);

CREATE TYPE status AS ENUM ('empty', 'marked', 'paid');

CREATE TABLE seats (
    seat_id SERIAL PRIMARY KEY,
    event_id INT NOT NULL REFERENCES events(event_id) ON DELETE CASCADE ,
    seat_number INT NOT NULL,
    seat_status status NOT NULL DEFAULT 'empty'
);