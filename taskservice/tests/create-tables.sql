DROP TABLE IF EXISTS task, assignee;

CREATE TABLE task (
    taskId BIGINT UNSIGNED AUTO_INCREMENT NOT NULL,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(250) NOT NULL,
    deadline TIMESTAMP NOT NULL,
    status TINYINT(255) NOT NULL,
    priority BOOL NOT NULL,
    creatorId BIGINT UNSIGNED NOT NULL,
    teamId BIGINT UNSIGNED,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`taskId`)
);

CREATE TABLE assignee (
    id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL,
    taskId BIGINT UNSIGNED NOT NULL,
    assigneeId BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO task
    (title, description, deadline, status, priority, creatorId, teamId)
VALUES
    ('Program', 'Help gavin code', '2025-12-12 00:00:00', 0, 0, 1, NULL),
    ('Eat', 'cook pasta', '2025-12-11 00:00:00', 1, 1, 2, 4),
    ('Write some code', 'refactor service layer', '2025-11-12 00:00:00', 0, 1, 1, 2),
    ('Drive home', 'get all things and go home', '2025-12-12 12:00:00', 0, 1, 1, NULL),
    ('Run', 'run a mile', '2025-12-12 00:00:30', 0, 1, 1, 2),
    ('Drink water', 'make sure to drink a gallon of water', '2025-12-1 00:00:00', 2, 0, 1, NULL);

INSERT INTO assignee
    (taskId, assigneeId)
VALUES
    (1, 3),
    (2, 6),
    (1, 4),
    (3, 3),
    (5, 9);