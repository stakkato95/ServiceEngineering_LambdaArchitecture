$podName = kubectl get pods --no-headers -o custom-columns=":metadata.name" | Select-String "service-transactional"
kubectl logs $podName