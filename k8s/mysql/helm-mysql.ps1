# https://artifacthub.io/packages/helm/bitnami/mysql
# github 
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install mysql bitnami/mysql --set auth.rootPassword=root --set auth.database=transactional