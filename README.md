# monstarlab
##monstarlab technical test

1. setup postgres docker as localDB
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
        
