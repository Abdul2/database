
## This is a small project to test out  GO lang DB approach 

'''
I will use this space to document, lessons I have learned and what i need to watch out for.

'''



**MySQL prep**

* db hello exists but table world did not exist in my sys (it has been a while since I have installed and used mysql) 



    * `./mysq -u root`
    * `CREATE USER 'user'@'localhost' identified by 'password';`
    * `GRANT ALL PRIVILEGES ON * . * TO 'user'@'localhost';`
    * `use hello;'
    * ' create table world (id INT, name VARCHAR(20));`
    * `insert into world (id,name) VALUES (2, "James");`
    select User,Host from mysql.user;
    
    SHOW VARIABLES LIKE 'wait_timeout';