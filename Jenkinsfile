pipeline {
    agent any
    
    environment {
        DOCKER_IMAGE_NAME = 'prady0t/minor-2'
    }

    stages {
        stage("echo") {
            steps {
                echo "This is a placeholde. We van run tests, builds etc etc like this"
            }
        }

        stage("Checkout") {
            steps {
                script {
                    checkout([$class: 'GitSCM', branches: [[name: 'main']], userRemoteConfigs: [[url: 'https://github.com/prady0t/Minor-2']]])
                }
            }
        }

        stage("Build Image") {
            steps {
                script {
                    def dockerImage = docker.build("${DOCKER_IMAGE_NAME}:${BUILD_NUMBER}", " . ")
                }
            }
        }
        stage("Push Image") {
            steps {
                // Authenticate with Docker Hub
                withCredentials([usernamePassword(credentialsId: 'dockerhub_cred', passwordVariable: 'DOCKERHUB_PASSWORD', usernameVariable: 'DOCKERHUB_USERNAME')]) {
                    // Log in to Docker Hub
                    sh "docker login -u ${DOCKERHUB_USERNAME} -p ${DOCKERHUB_PASSWORD}"

                    // Push the Docker image to Docker Hub
                    sh "docker push ${DOCKER_IMAGE_NAME}:${BUILD_NUMBER}"
                }
            }
        }

    }
}
