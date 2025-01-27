## Surge proxy

Generate a certificate for the following domains:

```
rm -rf tmp
mkdir tmp

mkcert -key-file tmp/key.pem -cert-file tmp/cert.pem gql-fed.reddit.com premium-noneu.truecaller.com api.busuu.com

mv "$(mkcert -CAROOT)"/rootCA.pem tmp
rm -rf "$(mkcert -CAROOT)"/rootCA-key.pem
```

Use the following docker-compose file:

```
services:
  surge-proxy:
    image: ghcr.io/labac-dev/surge-proxy:main
    container_name: surge-proxy
    restart: unless-stopped
    volumes:
      - ${PWD}/tmp:/tmp
    ports:
      - 443:443
```
