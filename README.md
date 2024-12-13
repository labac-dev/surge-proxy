## Surge proxy

Generate a certificate for the following domains:

```
mkcert -key-file tmp/key.pem -cert-file tmp/cert.pem gql-fed.reddit.com youtubei.googleapis.com
```

Run a container with the generated certificate:

```
docker run -d --name surge-proxy --restart unless-stopped -v $(pwd)/tmp:/tmp -p 443:443 ghcr.io/labac-dev/surge-proxy:main
```
