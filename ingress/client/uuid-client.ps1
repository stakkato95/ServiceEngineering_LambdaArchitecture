$id = [guid]::NewGuid().ToString()
$payload = '{ \"id\": \"' + $id + '\", \"name\": \"helloworld12\" }'
echo $payload
curl -X POST localhost/ingress/user -H 'Content-Type: application/json' -d $payload