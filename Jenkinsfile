pipeline {
    agent { docker { image 'golang' } }   
    stages {
        stage('install deps') {
            steps {
                sh 'go get -d github.com/magefile/mage'
                sh 'cd $GOPATH/src/github.com/magefile/mage && go run bootstrap.go install'


                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/github.com/yantrashala/prefab'

                // Copy all files in our Jenkins workspace to our project directory.                
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/github.com/yantrashala/prefab'
                sh 'GO111MODULE=on CGO_ENABLED=0 GOOS=linux mage -v get'
            }
        }
        stage('Test') {
            steps {
                // Remove cached test results.
                sh 'go clean -cache'

                // Run all Tests.
                sh 'go test ./... -v'
            }
        }
    }
}