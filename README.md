# Gaultier – Microservices E-Commerce Backend (Go)

## Overview
Gaultier is a microservices-based backend system for an e-commerce platform.
The system is implemented using Go and follows a distributed architecture approach with separated domain services.

The project demonstrates backend service decomposition, containerization, and inter-service communication.

## Architecture

Microservices included:
- User Service
- Product Service
- Order Service
- Cart Service
- Delivery Service
- Admin Service

Each service is independently deployable and runs inside Docker containers.

Architecture style:
Client → API → Microservices → Database

## Technologies
- Go (Golang)
- Docker & Docker Compose
- REST / gRPC communication
- Protobuf definitions
- PostgreSQL / MongoDB (if applicable)

## Key Features
- Service-based architecture
- Containerized deployment
- Domain separation
- Inter-service communication
- Scalable backend structure

## How to Run

```bash
git clone https://github.com/Dikolyandro/gaultier_git.git
cd gaultier_git
docker-compose up --build
