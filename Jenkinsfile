pipeline {
    agent { dockerfile true }

    environment {
        DOCKER_IMAGE = "myapp:latest"
    }

    stages {
        stage('Clone Repository') {
            steps {
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
        stage('Run Docker Container') {
            steps {
                script {
                    docker.image(DOCKER_IMAGE).run('-p 8080:8080')
                }
            }
        }
    }
}
