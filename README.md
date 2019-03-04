1. Creating the Database
   As our application is simple, we will create only one table called users with the following fields:

   id : is the primary key.
   firstname : is the first-name of the user.
   lastname : is the last-name of the user.
   age : is the age of the user.

    CREATE DATABASE api_task;
    USE api_task;
    CREATE TABLE users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        firstname VARCHAR(30) NOT NULL,
        lastname VARCHAR(30) NOT NULL,
        age INT NOT NULL
    );

2. Set GOPATH and GOROOT and create a workspace for Go source code. 

3. Clone the git library to your system.

4. Build the Go code and Run.

5. Server will up on the localhost and port :8080

6. I have used MySQL as a database with username as "root" and password as "s".

7. API Specification
   ------------------

 * Create a new user in response to a valid POST request at /user,
 * Update a user in response to a valid PUT request at /user/{id},
 * Delete a user in response to a valid DELETE request at /user/{id},
 * Fetch a user in response to a valid GET request at /user/{id}, and
 * Fetch a list of users in response to a valid GET request at /users.
The {id} will determine which user the request will work with.


