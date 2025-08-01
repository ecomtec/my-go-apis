pipeline {
    agent any

    environment {
        IMAGE_NAME = "my-go-apis"
        CONTAINER_NAME = "go-app"
        DOCKER_PORT = "8080:8080"
    }

    stages {
        stage('Clone Repository') {
            steps {
                git url: 'https://github.com/ecomtec/my-go-apis.git', branch: 'main'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build -t $IMAGE_NAME .'
            }
        }

        stage('Stop Existing Container') {
            steps {
                script {
                    // Stop and remove container if running
                    sh """
                        docker ps -q --filter name=$CONTAINER_NAME | grep -q . && docker stop $CONTAINER_NAME && docker rm $CONTAINER_NAME || true
                    """
                }
            }
        }

        stage('Run New Container') {
            steps {
                sh 'docker run -d --name $CONTAINER_NAME -p $DOCKER_PORT $IMAGE_NAME'
            }
        }

        stage('Deploy with Ansible') {
            steps {
                sh 'ansible-playbook deploy.yml'
            }
        }
    }
}
