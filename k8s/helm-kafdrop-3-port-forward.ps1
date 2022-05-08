$kafdrop = kubectl get pods --no-headers -o custom-columns=":metadata.name" | Select-String "kafdrop"
kubectl port-forward pod/$kafdrop 9000:9000