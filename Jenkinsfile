pipeline {

    agent {
        node {
            label 'slave_node1'
        }
    }
    stages {
        
        stage('Code Checkout') {
            steps {
             checkout scm
            }
        }
         stage ('Deploying our Api'){
            steps{
             checkout scm
             sh 'docker build -t user-id-string .'
             sh 'docker tag user-id-string 603389930669.dkr.ecr.us-east-1.amazonaws.com/user-id-string:latest'
             sh '$(aws ecr get-login --no-include-email --region us-east-1) > /dev/null'
             sh 'docker push 603389930669.dkr.ecr.us-east-1.amazonaws.com/user-id-string:latest'
             sh 'aws ecs update-service --cluster pratilipi --service user-id-string --force-new-deployment --region us-east-1'
            }

         }
    }   
}
