INSERT INTO skills (id, name) VALUES
(1, 'JavaScript'),
(2, 'Python'),
(3, 'SQL');


INSERT INTO projects (id, name, description) VALUES
(1, 'Website Development', 'Developed a responsive company website using HTML, CSS, and JavaScript.'),
(2, 'Data Analysis', 'Performed data analysis using Python to extract insights from large datasets.'),
(3, 'Database Design', 'Designed and optimized SQL databases for e-commerce platforms.');


INSERT INTO projects_skills (p_id, s_id) VALUES
(1, 1),  -- Website Development uses JavaScript
(2, 2),  -- Data Analysis uses Python
(3, 3);  -- Database Design uses SQL


INSERT INTO educations (id, name) VALUES
(1, 'Bachelor of Computer Science'),
(2, 'Master of Data Science'),
(3, 'Online Bootcamp in Web Development');


INSERT INTO languages (name) VALUES
('English'),
('Spanish'),
('French');


INSERT INTO userinfo (phonenumber, email, user_location, linkedin, github) VALUES
('123-456-7890', 'user1@example.com', 'New York, USA', 'https://linkedin.com/in/user1', 'https://github.com/user1');
