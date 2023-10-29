# Project "Trueconf-task"

## Table of Contents
- [Description](#description)
- [Constraints](#constraints)
- [Changes](#changes)
- [Future Improvements](#future-improvements)
- [Usage Instructions](#usage-instructions)
- [Authors](#authors)

## Description
This project is a simple web application for managing a list of users. It provides the following functionality:
- Search for users
- Create a new user
- Get user information by identifier
- Update user information
- Delete a user

## Constraints
The project maintains all the constraints that were set for the original version:

- User data is still stored in a JSON file.
- The user structure remains unchanged.
- The application does not lose its existing functionality.

## Changes
To improve the project and ensure better maintainability, I have made the following changes:
- Restructured the project by separating business logic, error handling, and routing.
- Added error handling using custom user-defined errors.
- Implemented logging using the Zap library.
- Created a Makefile for convenient project building and management.
- Added loggers for monitoring to the code.

## Future Improvements
In the future, could further enhance this application in the following ways:
- Add an authentication system and action restrictions for security.
- Move data storage to a database, such as MongoDB.
- Write tests to ensure reliability and stability.
- Implement metric collection for performance monitoring.
- Improve the project's architecture by separating it into more detail layers.
- Implement API versioning to ensure backward compatibility, added authorization.

## Usage Instructions
### Prerequisites
- Install Go on your computer (https://golang.org/dl/).
- Clone the repository: `git clone https://github.com/DmitriiKumancev/refactor-project.git`.

### Building
To build the application, run:
```
make build
```

### Running
To run the application, use the following command:
```
make run
```
### Cleaning
To clean up the project, execute:
```
make clean
```

### Interacting with the API
You can use an API client, such as [Postman](https://www.postman.com/), to interact with the API. The `http` folder contains files for performing various HTTP requests.

## Authors

- [Dmitrii Kumancev](https://github.com/DmitriiKumancev)