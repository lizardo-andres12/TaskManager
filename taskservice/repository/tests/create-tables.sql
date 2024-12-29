DROP TABLE IF EXISTS tasks, workers;

CREATE TABLE tasks (
    taskId  BIGINT UNSIGNED AUTO_INCREMENT NOT NULL,
    title VARCHAR(100) NOT NULL,
    priority TINYINT(1) NOT NULL,
    creatorId BIGINT UNSIGNED NOT NULL,
    assigneeId BIGINT UNSIGNED DEFAULT 0,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`taskId`)
);

CREATE TABLE workers (
    userId BIGINT UNSIGNED NOT NULL,
    taskId BIGINT UNSIGNED NOT NULL,
    username VARCHAR(16),
    PRIMARY KEY (`userId`, `taskId`)
);

INSERT INTO tasks
    (title, priority, creatorId, assigneeId)
VALUES
    ('Pair Program', 0, 1, 3),
    ('Write Documentation', 1, 2, 3),
    ('Run Tests', 1, 1, 3),
    ('Hire Interns', 0, 4, 0),
    ('Fix Bugs', 0, 5, 0);