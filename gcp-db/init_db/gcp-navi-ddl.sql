CREATE TABLE IF NOT EXISTS questions (
    id SERIAL primary key,
    question_id varchar(36) NOT NULL,
    question VARCHAR(1000) NOT NULL,
    category_id VARCHAR(10) NOT NULL
);

CREATE INDEX ON questions (question_id);

CREATE TABLE IF NOT EXISTS answers (
    id BIGSERIAL primary key,
    question_id VARCHAR(36) not null,
    correct_answer INTEGER not null,
    answer1 VARCHAR(1000) not null,
    answer2 VARCHAR(1000),
    answer3 VARCHAR(1000),
    answer4 VARCHAR(1000),
    answer5 VARCHAR(1000),
    answer6 VARCHAR(1000),
    answer7 VARCHAR(1000),
    answer8 VARCHAR(1000),
    answer9 VARCHAR(1000),
    answer10 VARCHAR(1000)
);

CREATE INDEX on answers (question_id);

CREATE TABLE IF NOT EXISTS category (
    id BIGSERIAL primary key,
    category_id VARCHAR(10) NOT NULL,
    category VARCHAR(100) NOT NULL
);