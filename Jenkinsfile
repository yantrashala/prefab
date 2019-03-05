pipeline {
    agent { docker ' golang:1.11' }
    stages {
        stage('install deps') {
            steps {
                sh 'go get -d github.com/magefile/mage'
                sh 'cd $GOPATH/src/github.com/magefile/mage && go run bootstrap.go install'
                sh 'GO111MODULE=on CGO_ENABLED=0 GOOS=linux mage -v get'
            }
        }

    }
}