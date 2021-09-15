# monstarlab
##monstarlab technical test

###run this
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres

Run Command
1. Docker exec -it {Container ID} bash
2. psql -h 0.0.0.0 -p 5432 -U postgres
    1. run following sql
        1.      CREATE TABLE movies (
                    movie_id SERIAL NOT NULL ,
                    movie_name varchar(50) NOT NULL,
                    genre varchar(50) NOT NULL,
                    sypnosis varchar NULL,
                    PRIMARY KEY (movie_id)
                );
        1.     INSERT INTO movies (
                    movie_id,
                    movie_name,
                    genre,
                    sypnosis
                )
                VALUES
                    ('1', 'Godzilla', 'horror', 'a story about dinasour'),
                    ('2 ', 'burnout', 'action', 'boy burn all his toys'),
                    ('3', 'need for speed', 'romance', 'a world where cars have unlimited gears');
        1.    CREATE TABLE users (
                    user_id SERIAL,
                    name varchar(100) DEFAULT '',
                    email varchar(100) DEFAULT '' UNIQUE,
                    password varchar(100) DEFAULT '',
                    updated_at timestamp NULL DEFAULT NULL,
                    created_at timestamp NULL DEFAULT NULL,
                    PRIMARY KEY (user_id)
               );
        1.     CREATE TABLE favourites(
                    fav_id SERIAL,
                    user_id int NOT NULL,
                    movie_id int NOT NULL,
                    PRIMARY KEY(fav_id),
                    CONSTRAINT fk_users
                    FOREIGN KEY(user_id) 
                    REFERENCES users(user_id)
                );
        1.      INSERT INTO favourites (
                    user_id,
                    movie_id
                )
                VALUES
                    ('1', '2'),
                    ('1 ', '3'),
                    ('1 ', '1');
                    
##Once everything has been run. go inside the monstarlab folder and run go run main.go

To use the API.
1. You need to signup (localhost:8080/user/signup) if you signup you dont need to login
2. If you already sign up you can head on to login (localhost:8080/user/login) 
3. token is automatically save in cookie. but for saving favourite you can use postman and copy the token from the browser.
4. The api accept token from cookie, param, query so any will suffice.
