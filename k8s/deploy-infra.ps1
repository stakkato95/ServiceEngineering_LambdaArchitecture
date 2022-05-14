./nginx-controller/helm-nginx-controller.ps1 > $null

./cassandra/helm-cassandra.ps1 > $null
./mysql/helm-mysql.ps1 > $null

./kafka/deploy/helm-1-kafka.ps1 > $null
# kafdrop should be installed via ubuntu in advance 
helm ls
./kafka/deploy/helm-4-kafdrop-port-forward.ps1