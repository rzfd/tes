pipeline {
    agent any
    
    environment {
        APP_NAME = "my-go-app"
        DOCKER_IMAGE = "${APP_NAME}:latest"
    }
    
    stages {
        stage('Checkout') {
            steps {
                // Checkout your code from version control
                git branch: 'main', url: 'https://github.com/rzfd/tes.git'
            }
        }
        
        stage('Build Docker Image') {
            steps {
                script {
                    docker.build(DOCKER_IMAGE)
                }
            }
        }
        
        stage('Run Containers') {
            steps {
                script {
                    // Use docker-compose to start services
                    sh 'docker-compose up -d'
                }
            }
        }
        
        stage('Run Tests') {
            steps {
                script {
                    sh 'docker --version'
                }
            }
        }
    }
}
