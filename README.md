### DDL
```sql
CREATE DATABASE task_db;

USE task_db;

CREATE TABLE tasks (
 id INT AUTO_INCREMENT PRIMARY KEY,
 title VARCHAR(255) NOT NULL,
 description TEXT NOT NULL,
 status VARCHAR(50) NOT NULL
);

-- in case i forgot
CREATE USER 'username'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* to 'username'@'%';
```

### Restore DB
```sh
# mysqldump -u username -ppasword task_db > db.sql
mysql -u userame -ppassword task_db < db.sql
```

### Run
```sh
# go get github.com/gofiber/fiber/v2 github.com/go-sql-driver/mysql github.com/joho/godotenv
go mod tidy
go run main.go
```

### etc in case i forgot
```sh
go mod init rizaldyaristyo-fiber-boiler
```