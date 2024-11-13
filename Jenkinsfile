pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "myapp:latest" 
    }

    stages {
        stage('Clone Repository') {
            steps {
                git 'https://github.com/rzfd/tes.git' 
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

    post {
        always {
            cleanWs()
        }
    }
}
