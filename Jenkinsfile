
pipeline {

    agent any

    environment {
        IMAGE_NAME = 'go-api'
        CONTAINER_NAME = 'go-api'
    }

    stages {

        stage('Checkout') {
            steps {
                echo "Building branch: ${env.BRANCH_NAME}"
                checkout scm
            }
        }

        stage('Go Test') {
            steps {
                echo 'Running Go tests...'
                sh 'go test ./...'
            }
        }

        stage('Docker Build') {
            steps {
                echo 'Building Docker image...'
                sh 'docker compose build'
            }
        }

        stage('Deploy') {
            when {
                branch 'main'
            }

            steps {
                echo 'Deploying application from main branch...'
                sh 'docker compose up -d'
            }
        }
    }

    post {
        success {
            echo "Pipeline successful for ${env.BRANCH_NAME}"
        }

        failure {
            echo "Pipeline failed for ${env.BRANCH_NAME}"
        }
    }
}
