# Burgher Social Feed

This repository contains the backend code for the Burgher Social Feed application, developed in Go. It includes modules for handling user posts, locations, topics, and storage, among others.

## Features

- **User Management**: Handle user authentication and profiles.
- **Post Management**: Create, read, update, and delete posts.
- **Location**: Index posts by locations.

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

### Sequence Diagram
![Sequence Diagram](./docs/image.png?raw=true "Burgher sequence diagram")


### Walk through For feed generation
-- Location Update:
- Fetch recent User Posts at the old location.
- Delete those posts from the old locations.
- Create posts with new location.

-- Refresh User feed after timeout:
- Fetch all posts by filtering with radius.
- Calculate relevancy score for user for each post.
- Flush existing feed.
- Insert all posts in the user feed.

-- User fetches feed:
- Return existing posts from user feed.
- Client stores posts in SQFLite database.
- Client manages sorting of posts based on score, decay with time and brings unseen posts on top.
