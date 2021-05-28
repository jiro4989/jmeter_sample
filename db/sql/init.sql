CREATE TABLE users (
  primary key (id),
  id varchar(255) not null,
  name varchar(255) not null,
  age int(3) not null
);
CREATE TABLE hobbies (
  primary key (id),
  id varchar(255) not null,
  user_id varchar(255) not null,
  name varchar(255) not null,
  foreign key fk_user_id(user_id) references users(id)
);
INSERT INTO
  users (id, name, age)
VALUES
  ('abf8f4ab-1020-4117-bd1a-4a6bdb8b6207', 'taro', 20),
  ('36548b82-bdb0-4fef-a7a8-bea7b546afc4', 'john', 22),
  (
    '5671d04d-fe93-4cba-82fa-b264fe344615',
    'hanako',
    18
  ),
  (
    '2534084b-7340-448e-bf8e-dd229e2bf8f3',
    'ichiro',
    21
  ),
  (
    '6837798d-cd0d-46bf-90f8-18707d93deee',
    'keisuke',
    25
  );
INSERT INTO
  hobbies (id, user_id, name)
VALUES
  (
    'd82ea5a1-7a0c-4bc2-8fc3-70f3012ddf55',
    'abf8f4ab-1020-4117-bd1a-4a6bdb8b6207',
    'baseball'
  ),
  (
    'cf269e0f-94bd-468d-8348-ec60dc5ab909',
    '36548b82-bdb0-4fef-a7a8-bea7b546afc4',
    'succer'
  ),
  (
    'deb3f8c3-3242-4a4a-9307-957f8bbf7fc9',
    '5671d04d-fe93-4cba-82fa-b264fe344615',
    'cooking'
  ),
  (
    '29504948-8cf0-42fd-95a8-0ce76adc55b0',
    '2534084b-7340-448e-bf8e-dd229e2bf8f3',
    'walking'
  ),
  (
    'f5c823aa-38c7-4fd1-addb-f826c1f33c0f',
    '6837798d-cd0d-46bf-90f8-18707d93deee',
    'dance'
  );
