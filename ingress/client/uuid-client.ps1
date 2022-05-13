$id = [guid]::NewGuid().ToString()
$payload = '{ \"id\": \"747c6ae4-0eff-4f9f-b1ef-31a44ddccf59\", \"name\": \"helloworld\" }'
echo $payload
curl -X POST localhost/ingress/user -H 'Content-Type: application/json' -d $payload