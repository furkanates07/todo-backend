
# Todo Backend Project

This is a **Todo Management Backend** application implemented in Go, utilizing **Supabase** as the database. The project includes features for user authentication, task creation, and management.

---

## Features

### Authentication
- **User Registration**  
  Allows new users to register with an email and password.  
  Passwords are securely hashed using `bcrypt`.  
  JWT tokens are used for user authentication.

- **User Login**  
  Authenticates users with their email and password.  
  Returns a JWT token upon successful login for session management.

### Todo Management
- **Create Todo**  
  Allows logged-in users to create a new todo task with a title and description.  
  Defaults to a `PENDING` status upon creation.

- **Retrieve Todos**  
  Fetches all todos associated with a specific user.

- **Retrieve a Single Todo**  
  Fetches details of a specific todo task by its ID.

- **Update Todo**  
  Updates the title, description, or any other field of a specific todo task.

- **Update Todo Status**  
  Updates the status (`PENDING`, `COMPLETED`, etc.) of a specific todo task.

- **Delete Todo**  
  Deletes a specific todo task by its ID.

---
## Prerequisites
###  1. Go Programming Language Install Go (version 1.19 or higher). You can download it from the [official website](https://golang.org/dl/). 

### 2. Supabase Project Set up a Supabase project at [Supabase](https://supabase.io/). 
Create the following tables in your Supabase database: 
#### Users Table
 ```sql 
 CREATE TABLE users ( id UUID PRIMARY KEY DEFAULT gen_random_uuid(), email TEXT NOT NULL UNIQUE, password TEXT NOT NULL, created_at TIMESTAMP DEFAULT NOW() );
```
#### Todo Table
 ```sql 
 CREATE TABLE todos ( id UUID PRIMARY KEY DEFAULT gen_random_uuid(), user_id UUID NOT NULL REFERENCES users(id), title TEXT NOT NULL, description TEXT, status TEXT DEFAULT 'PENDING', created_at TIMESTAMP DEFAULT NOW() );
```
### 3. Environment Variables  
Create a `.env` file in the project root directory with the following variables:  
```env
SUPABASE_URL=your_supabase_project_url
SUPABASE_KEY=your_supabase_api_key
JWT_SECRET=your_jwt_secret_key
````
-   Replace `your_supabase_project_url` with your Supabase project's URL.
-   Replace `your_supabase_api_key` with your Supabase API key.
-   Replace `your_jwt_secret_key` with a secure key for signing JWT tokens.

### Installation

Follow the steps below to set up and run the project:

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/todo-backend.git
   cd todo-backend
   ````
2. Install dependencies:
```bash
   go mod tidy
   ````
3. Set up environment variables: Create a `.env` file in the root directory and configure the required variables
4. Run the application:
. ```bash
   go run main.go
   ````
5. Access the API endpoints: Use a tool like Postman or cURL to interact with the API. By default, the server will run on `http://localhost:8080`.


### Technologies Used

- **Programming Language:** Go  
- **Database:** Supabase  
- **Authentication:** JWT (JSON Web Token)  
- **Password Hashing:** bcrypt  

## Contributing

Feel free to submit issues or pull requests if you find any bugs or want to add new features.
