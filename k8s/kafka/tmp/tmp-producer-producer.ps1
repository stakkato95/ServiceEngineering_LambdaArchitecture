kubectl run kafka-client-producer --restart='Never' --image docker.io/bitnami/kafka:3.1.0-debian-10-r89 --namespace default --command -- sleep infinity

# kafka-console-producer.sh --broker-list kafka-0.kafka-headless.default.svc.cluster.local:9092 --topic test