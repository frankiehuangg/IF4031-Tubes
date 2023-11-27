CREATE TABLE events (
    event_id INT PRIMARY KEY,
    event_name VARCHAR(256) UNIQUE NOT NULL,
    event_date TIMESTAMP NOT NULL
);

CREATE TYPE status AS ENUM ('empty', 'marked', 'paid');

CREATE TABLE seats (
    event_id INT NOT NULL REFERENCES events(event_id),
    client_id INT NOT NULL,
    seat_number INT NOT NULL,
    seat_status status NOT NULL,
    PRIMARY KEY (event_id, client_id, seat_number)
);