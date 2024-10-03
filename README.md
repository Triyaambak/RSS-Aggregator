# RSS Aggregator

The RSS Aggregator is a Go-based application that fetches and aggregates RSS feeds from various sources into a unified platform.

## Key Features

-> Go: Improve code reliability and maintainability with Go's strong typing and concurrency support, reducing bugs and enhancing development efficiency.

-> Docker: Containerize the application for consistent environments and streamlined deployment, ensuring scalability and ease of management across different setups.

-> PostgreSQL: Use PostgreSQL as the relational database for reliable data storage and complex query capabilities.

-> GitHub Actions: Automate workflows and streamline CI/CD processes with GitHub Actions, enabling efficient testing, building, and deployment of applications, while enhancing collaboration and integration across teams.

-> AWS EC2: Utilize Amazon EC2 for scalable and flexible cloud computing, allowing for quick deployment of applications on virtual servers with customizable configurations.

-> AWS ECS: Leverage Amazon ECS for efficient container orchestration, enabling seamless management and scaling of Docker containers in a fully managed environment.

-> AWS ECR: Use Amazon ECR for secure and scalable container image storage, simplifying the process of storing, managing, and deploying Docker images within the AWS ecosystem.

## Getting Started

```bash
// HTTP
> git clone https://github.com/Triyaambak/RSS-Aggregator.git

// SSH
> git clone git@github.com:Triyaambak/RSS-Aggregator.git

```
## Setting up docker-compose.yml env variables

```bash
> touch .env
```

### Your .env file should contain the following 

```env
POSTGRES_PORT=xxxx
FRONTEND_PORT=xxxx
BACKEND_PORT=xxxx
REDIS_PORT=xxxx
```

## Now after everything is set

```bash
> docker compose up --watch
```

