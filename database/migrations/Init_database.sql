create database todo_activity;


CREATE TABLE person  (
                         user_id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
                         username VARCHAR(50) NOT NULL,
                         email VARCHAR(100) NOT NULL UNIQUE,
                         password_hash VARCHAR(255) NOT NULL,
                         create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE project (
                         project_id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
                         user_id INT,
                         project_name VARCHAR(100) NOT NULL,
                         created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         FOREIGN KEY (user_id) REFERENCES person(user_id)
);

CREATE TYPE priority_enum AS ENUM ('Low', 'Medium', 'High');
CREATE TYPE status_enum AS ENUM ('Not Started', 'In Progress', 'Completed');

CREATE TABLE task (
                      task_id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
                      project_id INT,
                      task_name VARCHAR(100) NOT NULL,
                      description TEXT,
                      due_date DATE,
                      priority priority_enum,
                      status status_enum,
                      created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      FOREIGN KEY (project_id) REFERENCES project(project_id)
);

CREATE TABLE comment (
                         comment_id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
                         task_id INT,
                         user_id INT,
                         comment_text TEXT NOT NULL,
                         create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         FOREIGN KEY (task_id) REFERENCES task(task_id),
                         FOREIGN KEY (user_id) REFERENCES person(user_id)
);

CREATE TABLE reminder (
                          reminder_id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
                          task_id INT,
                          reminder_date DATE NOT NULL,
                          FOREIGN KEY (task_id) REFERENCES task(task_id)
);
