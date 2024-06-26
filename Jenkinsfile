pipeline {
    agent any
    
    environment {
        DOCKER_IMAGE_NAME = 'prady0t/minor-2'
    }

    stages {
        stage("start") {
            steps {
                echo "We start here."
            }
        }

        stage("Checkout") {
            steps {
                script {
                    checkout([$class: 'GitSCM', branches: [[name: 'main']], userRemoteConfigs: [[url: 'https://github.com/prady0t/Minor-2']]])
                }
            }
        }

        stage('Run Tests') {
            steps {
                // Install dependencies (if needed)
                script {
                    sh 'echo placeholder to run tests'
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

        stage('Update Deployment File'){
        environment {
            GIT_REPO_NAME = "Minor-2"
            GIT_USER_NAME = "prady0t"
        }
        steps {
            withCredentials([string(credentialsId: 'Github', variable: 'GITHUB_TOKEN')]) {
                sh '''
                    git config user.email "pradyot605@gmail.com"
                    git config user.name "prady0t"
                    BUILD_NUMBER=${BUILD_NUMBER}
                    sed -i "s|prady0t/minor-2:[0-9]*|prady0t/minor-2:${BUILD_NUMBER}|g" manifest/deployment.yaml
                    git add .
                    git commit -m "Update deployment image to version ${BUILD_NUMBER}"
                    git push https://${GITHUB_TOKEN}@github.com/${GIT_USER_NAME}/${GIT_REPO_NAME} HEAD:refs/heads/main

                '''
            }
        }
    }
    }
}

