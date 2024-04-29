pipeline {
  agent { label 'jenkins-agent' }
  tools {
    go 'Go1.20'
  }
  environment {
        SCANNER_HOME=tool 'sonar-scanner'
        APP_NAME='go-webapp'
        DOCKER_USER='muuzii'
        IMAGE_NAME="${DOCKER_USER}" + "/" + "${APP_NAME}"
        IMAGE_TAG="${env.BUILD_NUMBER}"
    }

  stages {
    stage('Clean workspace') {
      steps {
        cleanWs()
      }
  }

  stage('Checkout') {
    steps {
      git branch: 'main', credentialsId: 'github', url: 'https://github.com/muzammil-iftikhar/go-webapp.git'
    }
  }

  stage('SonarQube Analysis') {
    steps {
      
      withSonarQubeEnv('sonar-server') {
      sh '''$SCANNER_HOME/bin/sonar-scanner \
      -Dsonar.projectName=go-webapp \
      -Dsonar.projectKey=go-webapp \
      -Dsonar.sources=. '''
      }
    }
  }

    stage('Quality Gate') {
    steps {
      script {
        waitForQualityGate abortPipeline: true, credentialsId: 'jenkins-sonar-token'
        }
      }
    }

    stage('Docker build and push') {
      steps {
        withDockerRegistry([credentialsId: 'docker', url: '']) {
          sh 'docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .'
          sh 'docker push ${IMAGE_NAME}:${IMAGE_TAG}'
        }
      }
    }

    stage('Trivy scan') {
      steps {
    sh ('docker run -v /var/run/docker.sock:/var/run/docker.sock aquasec/trivy image ${IMAGE_NAME}:${IMAGE_TAG} --no-progress --scanners vuln --exit-code 0 --severity HIGH,CRITICAL --format table')
      }
    }

    stage('Clean up images') {
      steps {
        sh 'docker rmi ${IMAGE_NAME}:${IMAGE_TAG}'
      }
    }
  }
}
