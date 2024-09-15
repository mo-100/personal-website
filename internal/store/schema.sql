CREATE TABLE skills (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE projects (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  description text NOT NULL
);

CREATE TABLE educations (
  id BIGSERIAL,
  name text NOT NULL
);

CREATE TABLE projects_skills (
  p_id BIGSERIAL,
  s_id BIGSERIAL,
  PRIMARY KEY (p_id, s_id),
  FOREIGN KEY (p_id) REFERENCES projects(id),
  FOREIGN KEY (s_id) REFERENCES skills(id)
);

CREATE TABLE languages (name text NOT NULL);

CREATE TABLE userinfo (
  phonenumber text NOT NULL,
  email text NOT NULL,
  user_location text NOT NULL,
  linkedin text NOT NULL,
  github text NOT NULL
);