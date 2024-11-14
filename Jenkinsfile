pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "myapp:latest"
    }

    stages {
        stage('Clone Repository') {
            steps {
                git url: 'https://github.com/rzfd/tes.git'
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
            script {
                // Clean up Docker containers and images to ensure a clean environment
                def containers = docker.ps('--filter ancestor=' + DOCKER_IMAGE).trim().split('\n')
                if (containers.size() > 1) {
                    for (int i = 1; i < containers.size(); i++) {
                        docker.command("stop ${containers[i]}")
                        docker.command("rm ${containers[i]}")
                    }
                }
                def images = docker.images(DOCKER_IMAGE).trim().split('\n')
                if (images.size() > 1) {
                    for (int i = 1; i < images.size(); i++) {
                        docker.command("rmi ${images[i]}")
                    }
                }
            }
        }
    }
}
