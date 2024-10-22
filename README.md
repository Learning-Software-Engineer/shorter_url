### Techstack:
- __Language:__ `Go`
- __Database:__ Cassandra

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


### There are two approach for generating short url
#### Approach 1: Hashing and Encoding
A common approach for generating short URLs is to use a hash function, such as MD5 or SHA-256 to generate a fixed-length hash of the original URL.
This hash is then encoded into a shorter form using Base62.
Base62 uses alphanumeric characters (A-Z, a-z, 0–9), which are URL-friendly and provide a dense encoding space.
The length of the short URL is determined by the number of characters in the Base62 encoded string.
1. User submits a request to generate short url for the long url: https://www.example.com/some/very/long/url/that/needs/to/be/shortened
2. Generate an MD5 hash of the long URL. MD5 produces a 128-bit hash, typically a 32-character hexadecimal string: 1b3aabf5266b0f178f52e45f4bb430eb
3. Instead of encoding the entire 128-bit hash, we typically use a portion of the hash (e.g., the first few bytes) to create a more manageable short URL.
4. First 6 bytes of the hash: 1b3aabf5266b
5. Convert these bytes to decimal: 1b3aabf5266b (hexadecimal) → 47770830013755 (decimal)
6. Encode the result into a Base62 encoded string: DZFbb43