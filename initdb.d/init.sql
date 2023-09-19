CREATE TABLE persons (
  person_id SERIAL PRIMARY KEY,
  first_name VARCHAR(50),
  last_name VARCHAR(50),
  email VARCHAR(100),
  birth_date DATE,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  last_reserved_at TIMESTAMP
);

INSERT INTO persons (first_name, last_name, email, birth_date, created_at, last_reserved_at)
VALUES
('Alice', 'Yamada', 'alice.yamada@example.com', '1987-02-15', '2023-09-11 09:00:00+09', '2023-09-10 08:00:00'),
('Bob', 'Suzuki', 'bob.suzuki@example.com', '1992-05-28', '2023-09-11 10:00:00+09', NULL),
('Charlie', 'Tanaka', 'charlie.tanaka@example.com', '1976-11-08', '2023-09-11 11:00:00+09', '2023-09-11 09:00:00');
