# Person Tracker

Person Tracker is a Go application that allows you to manage a database of people. It provides commands to insert a new person into the database and query all people from the database.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go
- Docker

### Installing

Clone the repository to your local machine:

`git clone https://github.com/username/person-tracker.git`

Navigate to the project directory:

`cd person-tracker`

Build the Docker image:

`docker build -t person-tracker .`

Run the Docker container:

`docker run -it person-tracker ./main`

## Usage

To insert a new person into the database, use the insert command:

`docker run -it person-tracker ./main insert`

To query all people from the database, use the queryall command:

`docker run -it person-tracker ./main queryall`


## Available commands

- `insert`: Inserts a new person into the database.
- `queryall`: Queries all people from the database.
- `openai`: Interacts with OpenAI's API (chatgpt).


## Contributing

Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the LICENSE.md file for details


