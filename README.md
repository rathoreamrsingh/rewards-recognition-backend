# Rewards and Recognition Internal Project
## Overview
TODO - just for testing

## Features
TODO

## 🛠️ Usage

### Prerequisites

* Ensure you have Docker and Docker Compose installed.
* A MongoDB instance is required. This guide uses a Docker container for MongoDB.

### Steps to Run the Application

1.  **Create `.env` File:**

    * In the project's base directory, create a file named `.env`.
    * Add the following configuration, replacing the bracketed placeholders with your actual values:

        ```
        PORT=8080
        APP_ENV=local
        APPNAME=rewardsAndRecognition
        PROTOCOL=mongodb
        USERNAME=<your_mongodb_username>
        PASSWORD=<your_mongodb_password>
        HOST=localhost:27017
        ```

2.  **Start MongoDB Container:**

    * From the project's base directory, run:

        ```bash
        make up
        ```

        This command starts a MongoDB container as defined in the `docker-compose.yml` file.

3.  **Initialize MongoDB:**

    * To set up the database and collections:
        * Connect to the MongoDB instance using a tool like MongoDB Compass.
        * Create a database named "rewardsAndRecognition".
        * Create two collections within this database: "users" and "points".
        * Import the necessary data into the "users" and "points" collections from the files provided in the `db-data` folder.

4.  **Run the Application:**

    * Navigate to the "main" directory.
    * Execute the following command:

        ```bash
        go run .
        ```

## Configuration
TODO

## Contributing
TODO

## License
TODO

## Acknowledgements
TODO

