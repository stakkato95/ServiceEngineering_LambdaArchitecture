$id = [guid]::NewGuid().ToString()
$userName = "user12345"
$payload = '{ \"id\": \"' + $id + '\", \"name\": \"' + $userName + '\" }'
echo $payload
curl -X POST localhost/ingress/user -H 'Content-Type: application/json' -d $payload