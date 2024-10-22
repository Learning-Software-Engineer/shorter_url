### How to set up project:
1. Install cassandra by docker: `docker pull cassandra`.
   
2. Create docker-compose.yml file follow project.
   
3. Create database flow these steps:
- ``docker exec -it cassandra cqlsh``
- ``CREATE KEYSPACE shorturl WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};``, note: you can replace shorturl by other name.
- `use shorturl`
- Create tables: this project has two tables user and url_mapping
  + `create table user (user_id int primary key, name text, email text, password text);`
  + `create table url_mapping(short_url text primary key, long_url text, creation_date timestamp, expiration_date timestamp, user_id int, click_count int);`
  
__Resource for designing project__: https://medium.com/@ashishps/design-a-url-shortener-system-design-interview-b8226c8c94af
