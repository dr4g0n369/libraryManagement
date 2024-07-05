# Library Management System

Welcome to the Library Management System! This project aims to streamline the management of a library's inventory, including functionalities for adding, viewing, and removing books.

## Getting Started

### Prerequisites

Before you begin, ensure you have met the following requirements:

- You have Docker installed on your system.

### Installation

```
git clone https://github.com/dr4g0n369/libraryManagement.git
cd libraryManagement
cp config/.env .env
docker-compose build --no-cache
```
## Usage

After completing the installation you can start the server using the below command

```
docker-compose up -d
```

This will launch the application. You can visit it by going to `http://localhost:3000`.

In order to shut it down run the below command

```
docker-compose down
```