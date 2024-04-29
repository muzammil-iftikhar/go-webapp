pipeline {
  agent { label 'jenkins-agent' }
  tools {
    jkd 'Java17'
    go 'Go1.20'
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
}
}
