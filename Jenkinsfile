pipeline {
    agent any

    environment {
        // Sesuaikan dengan username Docker Hub Anda
        DOCKER_USER  = 'manggalakrida46' 
        IMAGE_NAME   = 'simple-go'
        IMAGE_TAG    = "${BUILD_NUMBER}"
        
        DOCKER_CREDS = 'dockerhub-creds'
        NAMESPACE  = "devops-academy"
    }

    stages {
        stage('Clone Repository') {
            steps {
                echo 'Mengambil kode terbaru dari GitHub...'
                // Proses checkout dilakukan otomatis jika menggunakan Pipeline from SCM
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    echo "Kompilasi aplikasi dan pembuatan Docker Image..."
                    appImage = docker.build("${DOCKER_USER}/${IMAGE_NAME}:${IMAGE_TAG}")
                }
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    echo "Mengunggah image ke Docker Hub..."
                    docker.withRegistry('', "${DOCKER_CREDS}") {
                        appImage.push()
                        appImage.push("latest")
                    }
                }
            }
        }

        // stage('Deploy Kubernetes') {

        //     steps {
        //         sh '''
        //             kubectl -n ${NAMESPACE} set image deployment/${DEPLOYMENT} \
        //             ${CONTAINER}=${IMAGE_NAME}:${IMAGE_TAG}

        //             kubectl rollout status deployment/${DEPLOYMENT} \
        //             -n ${NAMESPACE} --timeout=300s
        //         '''
        //     }
        // }


        stage('Deploy to Kubernetes') {
            steps {
                    echo 'Memodifikasi manifest Kubernetes...'
                    // Mengganti teks placeholder dengan image tag yang spesifik
                    sh """
                        sed -i 's|DOCKERHUB_USER/IMAGE_NAME:IMAGE_TAG|${DOCKER_USER}/${IMAGE_NAME}:${IMAGE_TAG}|g' k8s/deploy.yaml
                    """
                    
                    echo 'Menerapkan konfigurasi ke Cluster Kubernetes...'
                    sh 'kubectl apply -f k8s/deploy.yaml -n ${NAMESPACE}'
            }
        }
    }

    post {
        always {
            echo 'Pembersihan workspace Jenkins untuk menghemat disk...'
            cleanWs()
        }
        success {
            echo 'CI/CD Pipeline Berhasil Dieksekusi!'
        }
        failure {
            echo 'Pipeline Gagal. Harap periksa log untuk detail error.'
        }
    }
}