#!groovy

pipeline {
	agent {
	    label 'swarm'
	}
	options {
		timestamps()
		buildDiscarder(logRotator(numToKeepStr: '5'))
	}

	stages {
		stage("Build") {
			steps {
				sh """
					docker compose -f docker-compose.build.yaml run --rm builder
				"""
			}
		}
		stage("Build docker images") {
			steps {
				sh """
						docker buildx build --platform linux/amd64 -t ghcr.io/iv1310/trial-app-service:${BUILD_NUMBER} .
				"""
			}
		}
		stage("Security check for the image") {
		    steps {
		        sh """
		            trivy image ghcr.io/iv1310/trial-app-service:${BUILD_NUMBER}
		        """
		    }
		}
		stage("Push the image to registry") {
		    steps {
		        withCredentials([string(credentialsId: 'CR_PAT', variable: 'CR_PAT')]) {
		            sh """
		                echo ${CR_PAT} | docker login ghcr.io -u iv1310 --password-stdin
		                docker push ghcr.io/iv1310/trial-app-service:${BUILD_NUMBER}
		            """
		        }
		    }
		}
	}
	post {
		always {
			cleanWs()
		}
	}
}
