# URL Shortener

A full-stack URL shortener application built with **Go** (backend), **React** (frontend), and **Redis** (database). The application allows users to shorten long URLs and provides a simple analytics dashboard to track the number of clicks on each shortened URL.

## Features

- **Shorten URLs**: Convert long URLs into short, easy-to-share links.
- **Analytics Dashboard**: Track the number of clicks for each shortened URL.
- **Responsive UI**: Built with modern React components and Tailwind CSS.
- **Scalable Backend**: Built with Go and deployed using Docker and Kubernetes.
- **Redis Integration**: Stores URLs and click counts efficiently.

## Tech Stack

- **Frontend**: React, Tailwind CSS
- **Backend**: Go (Gin framework)
- **Database**: Redis
- **Containerization**: Docker
- **Orchestration**: Kubernetes (optional)
- **Environment Management**: Docker Compose

## Project Structure
url-shortener/
├── backend/ # Backend code (Go)
│ ├── Dockerfile # Dockerfile for the backend
│ ├── go.mod # Go module file
│ ├── main.go # Entry point for the backend
│ └── ... # Other backend files
├── frontend/ # Frontend code (React)
│ ├── Dockerfile # Dockerfile for the frontend
│ ├── nginx.conf # Nginx configuration
│ ├── package.json # Frontend dependencies
│ ├── src/ # React source code
│ └── ... # Other frontend files
├── kubernetes/ # Kubernetes deployment files
│ ├── backend-deployment.yaml
│ ├── backend-service.yaml
│ ├── redis-deployment.yaml
│ ├── redis-service.yaml
│ ├── frontend-deployment.yaml
│ ├── frontend-service.yaml
│ └── ingress.yaml
├── docker-compose.yaml # Docker Compose file for local development
└── README.md # Project documentation

Copy

## Prerequisites

Before running the project, ensure you have the following installed:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- (Optional) [Kubernetes](https://kubernetes.io/docs/setup/) for production deployment

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/url-shortener.git
cd url-shortener
2. Run with Docker Compose
To run the application locally using Docker Compose:

bash
Copy
docker-compose up -d
This will start the following services:

Frontend: Accessible at http://localhost:3000

Backend: Accessible at http://localhost:8080

Redis: Running on localhost:6379

3. Access the Application
Open your browser and navigate to http://localhost:3000.

Enter a long URL and click "Shorten".

View the shortened URL and click count in the list.

4. Stop the Application
To stop the application, run:

bash
Copy
docker-compose down
Kubernetes Deployment (Optional)
To deploy the application on Kubernetes:

Navigate to the kubernetes folder:

bash
Copy
cd kubernetes
Apply the Kubernetes manifests:

bash
Copy
kubectl apply -f .
Access the application:

If using Minikube, run:

bash
Copy
minikube service frontend
If using a cloud provider, use the external IP of the frontend service.

API Endpoints
Backend API
Shorten a URL:

Method: POST

Endpoint: /shorten

Request Body:

json
Copy
{
  "long_url": "https://example.com"
}
Response:

json
Copy
{
  "short_code": "abc123",
  "long_url": "https://example.com",
  "clicks": 0
}
Redirect to Original URL:

Method: GET

Endpoint: /:short_code

Response: Redirects to the original URL.

Get All URLs:

Method: GET

Endpoint: /urls

Response:

json
Copy
[
  {
    "short_code": "abc123",
    "long_url": "https://example.com",
    "clicks": 5
  }
]
Screenshots
Home Page
Home Page

Shortened URLs List
URLs List

Contributing
Contributions are welcome! If you'd like to contribute, please follow these steps:

Fork the repository.

Create a new branch (git checkout -b feature/YourFeatureName).

Commit your changes (git commit -m 'Add some feature').

Push to the branch (git push origin feature/YourFeatureName).

Open a pull request.

License
This project is licensed under the MIT License. See the LICENSE file for details.

Made with ❤️ by Your Name

Copy

---

### **How to Use**

1. Copy the entire Markdown content above.
2. Paste it into your `README.md` file in the root of your project.
3. Replace placeholders like `your-username` with your actual GitHub username.
4. Add screenshots to the `screenshots` folder and update the paths in the `## Screenshots` section.

---

### **Example Screenshots**

If you want to add screenshots:
1. Create a `screenshots` folder in the root directory.
2. Add your screenshots (e.g., `home.png`, `list.png`).
3. Update the paths in the `## Screenshots` section:
   ```markdown
   ![Home Page](screenshots/home.png)
   ![URLs List](screenshots/list.png)