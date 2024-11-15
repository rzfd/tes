pipeline {
    agent any
    
    environment {
        APP_NAME = "my-go-app"
        DOCKER_IMAGE = "${APP_NAME}:latest"
    }
    
    stages {
        stage('Checkout') {
            when {
                branch 'main'  // Ensures the pipeline runs only for the main branch
            }
            steps {
                // Checkout your code from version control
                git 'https://your-repository-url.git'
            }
        }
        
        stage('Build Docker Image') {
            when {
                branch 'main'  // Build the image only for the main branch
            }
            steps {
                script {
                    docker.build(DOCKER_IMAGE)
                }
            }
        }
        
        stage('Run Containers') {
            when {
                branch 'main'  // Run containers only for the main branch
            }
            steps {
                script {
                    // Use docker-compose to start services
                    sh 'docker-compose up -d'
                }
            }
        }
        
        stage('Run Tests') {
            when {
                branch 'main'  // Run tests only for the main branch
            }
            steps {
                script {
                    // Add any test commands you need to run
                    // For example:
                    // sh 'docker exec my-go-app-container go test ./...'
                }
            }
        }
    }

