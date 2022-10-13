CREATE TABLE Transaction(
    id VARCHAR(64) NOT NULL,
    issue_date INT NOT NULL,
    due_date INT NOT NULL,
    return_date INT,
    book_id VARCHAR(64) NOT NULL,
    user_id VARCHAR(64) NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY(user_id) REFERENCES users(id),
    on delete cascade,
    FOREIGN KEY(book_id) REFERENCES books(id)
    on delete cascade
    );