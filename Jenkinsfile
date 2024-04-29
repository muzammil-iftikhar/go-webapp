pipeline {
  agent { label 'jenkins-agent' }
  tools {
    go 'Go1.20'
  }
  environment {
        SCANNER_HOME=tool 'sonar-scanner'
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
  stage('Build') {
    steps {
      sh 'go build -o main .'
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
}
}
