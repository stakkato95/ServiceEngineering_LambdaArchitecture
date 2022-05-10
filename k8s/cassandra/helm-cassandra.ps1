# https://artifacthub.io/packages/helm/bitnami/cassandra
# why --generate-name https://github.com/bitnami/charts/issues/1838?msclkid=845a5a10d05011ec8c136ad6ca75af36
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install bitnami/cassandra --generate-name --set dbUser.user=cassandra --set dbUser.password=cassandra