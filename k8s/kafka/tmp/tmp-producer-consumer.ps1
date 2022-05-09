kubectl run kafka-client-producer --restart='Never' --image docker.io/bitnami/kafka:3.1.0-debian-10-r89 --namespace default --command -- sleep infinity

# kafka-console-consumer.sh --bootstrap-server kafka.default.svc.cluster.local:9092 --topic test --from-beginning