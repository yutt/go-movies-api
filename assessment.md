
## Objective:


An online stream platform has requested us to create an API to manage films
and user's favourites list. As the backend developer of this integration, you will be
responsible of the design and implementation of the solution.
You should use GO to implement the proposed solution and take advantage of
any existing library that could be useful.

## Requirements:

- Logged users should authenticate their requests via JWT token
- All endpoints related to movies should only be able to be used by
registered users
- Data should be stored in a database of your choice (MySQL,
MongoDB...)
- Create the following endpoints:
  - Register
    - Username (alphanumeric starting with letter)
    - Password (add some validations like maximum number of
characters...)
  - Login via username and password (should return the JWT token with
expiration time)
  - Create film. Suggested fields (Title, director, release date, cast, genre,
synopsis...)
    - The film should have a reference to the user who created it.
    - The films should be unique (different titles)
  - Get films list. Include optional custom filters as title, release date, genre)
  - Get film details. Include the creator user
  - Update a film
    - The film can only be modified by the user who created it but all
users can see and search for it. (Not a personal list)
  - Delete a film
    - The film can only be deleted by the user who created it.
 - Create a simple migration script to create database structure and initial
data
 - This are the minimum suggested requirements, feel free to expand them.



### Required:

- Clean solution architecture
- Good code quality. Keep it simple and readable
- Security mechanism: Authentication, authorisation, prevent SQL
injection...
- Proper use of common libraries
- Upload the project to a git repository server like Github, Gitlab or
Bitbucket.
- Use any HTTP API routing library (depict the reason of your choice)
- Use environment variable to configure your system

### Nice to have:

- ORM
- Proper Error handling
- Logging system
- Deployment instructions
- Possibility to be run automatically using Docker (docker composer)
- Use a code formatter library.
