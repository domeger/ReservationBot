-- reservation_bot.STATES tables setup
DESCRIBE reservation_bot.STATES;
SELECT * FROM reservation_bot.STATES;
INSERT INTO reservation_bot.STATES (ID, NAME, DESCRIPTION) VALUES (1, 'unavailable', 'unavailable for check-out');
INSERT INTO reservation_bot.STATES (ID, NAME, DESCRIPTION) VALUES (2, 'available', 'available for check-out');
INSERT INTO reservation_bot.STATES (ID, NAME, DESCRIPTION) VALUES (3, 'disabled', 'disabled resource');

-- reservation_bot.RESOURCES
DESCRIBE reservation_bot.RESOURCES;
SELECT * FROM reservation_bot.RESOURCES;
INSERT INTO reservation_bot.RESOURCES (NAME, DESCRIPTION, USER, STATE) VALUES ('stage01', 'Staging Environment 01', 0, 2);
INSERT INTO reservation_bot.RESOURCES (NAME, DESCRIPTION, USER, STATE) VALUES ('int01', 'Integration Environment 01', 0, 2);
INSERT INTO reservation_bot.RESOURCES (NAME, DESCRIPTION, USER, STATE) VALUES ('int02', 'Integration Environment 02', 0, 2);
INSERT INTO reservation_bot.RESOURCES (NAME, DESCRIPTION, USER, STATE) VALUES ('int03', 'Integration Environment 03', 0, 2);
INSERT INTO reservation_bot.RESOURCES (NAME, DESCRIPTION, USER, STATE) VALUES ('int04', 'Integration Environment 04', 0, 2);
INSERT INTO reservation_bot.RESOURCES (NAME, DESCRIPTION, USER, STATE) VALUES ('uswest1a', 'Prod West', 0, 2);
INSERT INTO reservation_bot.RESOURCES (NAME, DESCRIPTION, USER, STATE) VALUES ('useast1a', 'Prod East', 0, 2);

-- reservation_bot.PERMISSIONS
DESCRIBE reservation_bot.PERMISSIONS;
SELECT * FROM reservation_bot.PERMISSIONS;
INSERT INTO reservation_bot.PERMISSIONS (NAME, DESCRIPTION) VALUES (1, 'checkin-checkout', '');

-- reservation_bot.USERS
DESCRIBE reservation_bot.USERS;
SELECT * FROM reservation_bot.USERS;
INSERT INTO reservation_bot.USERS (NAME, PERMISSION) VALUES ('ablack', 1);
INSERT INTO reservation_bot.USERS (NAME, PERMISSION) VALUES ('domeger', 1);
INSERT INTO reservation_bot.USERS (NAME, PERMISSION) VALUES ('bdowns', 1);

