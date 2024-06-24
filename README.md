# Burgher Social Feed

This repository contains the backend code for the Burgher Social Feed application, developed in Go. It includes modules for handling user posts, locations, topics, and storage, among others.

## Features

- **User Management**: Handle user authentication and profiles.
- **Post Management**: Create, read, update, and delete posts.
- **Topic and Location**: Categorize posts by topics and locations.
- **Storage**: Manage media and other assets.

## Requirements

- Go 1.18+
- Docker

## Installation

1. Clone the repository:
   ```sh
    git clone https://github.com/burgher-social/feed.git
    cd feed
   ```
2. Build the Docker containers:
  ```sh
  docker-compose up --build
  ```
3. Usage
Start the server:
  ```sh
  go run server.go
  ```


Contributing
Please submit issues and pull requests for improvements and new features.

License
This project is licensed under the MIT License.
