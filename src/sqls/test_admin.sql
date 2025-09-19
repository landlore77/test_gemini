DROP TABLE IF EXISTS admin_auth;

CREATE TABLE admin_auth (
    id INT AUTO_INCREMENT PRIMARY KEY,
    group_id INT DEFAULT -1,
    user_name VARCHAR(64),
    description VARCHAR(64),
    password VARCHAR(64),
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_ip VARCHAR(15),
    last_login INT DEFAULT 0,
    INDEX (user_name),
    INDEX (create_time)
);
