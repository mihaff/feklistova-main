CREATE TABLE if not exists departments
(
    id              serial PRIMARY KEY,
    department_name varchar
);
CREATE TABLE if not exists employees
(
    id            serial PRIMARY KEY,
    surname       varchar,
    name          varchar,
    department_id int,
    constraint dep_fk foreign key(department_id) references departments(id)
);