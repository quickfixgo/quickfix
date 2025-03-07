// prevents build from occurring as a result of indexing:
//  if build was caused by indexing returns true in addition to setting the build result to the previous build result
//  otherwise it returns false
if (causedByIndexing()) {
  return
}

pipeline {
  agent {
    label 'linux'
  }

  stages {
    stage('Build and Test') {
      steps {
        deleteDir()
        checkout scm
        sh 'docker build --pull .'
      }
      post {
        cleanup {
          cleanWs(deleteDirs: true, disableDeferredWipeout: true, notFailBuild: true)
        }
      }
    }
  }

  post {
    success {
      slackSend(
        channel: '#eotc-jenkins',
        message: "SUCCESS: quickfix build: '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})",
        color: 'good',
        teamDomain: 'iceland',
        tokenCredentialId: '2b40808d-04b7-415b-8401-a4d500a84aab'
      )
    }
    failure {
      slackSend(
        channel: '#eotc-jenkins',
        message: "FAILURE: quickfix build: '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})",
        color: 'danger',
        teamDomain: 'iceland',
        tokenCredentialId: '2b40808d-04b7-415b-8401-a4d500a84aab'
      )
    }
  }
}
