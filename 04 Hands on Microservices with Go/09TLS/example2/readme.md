

## Generate Self Signed Certificates

1. create client.key, client.pem, server.key, server.pem

openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key

openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650 -subj "/C=AR/ST=CABA/L=CABA/O=Example Org/OU=IT Department/CN=*"

openssl genrsa -out client.key 2048
openssl ecparam -genkey -name secp384r1 -out client.key

openssl req -new -x509 -sha256 -key client.key -out client.pem -days 3650 -subj "/C=GB/ST=LONDON/L=LONDON/O=Another Org/OU=IT Department/CN=*"




2. In terminal, navigate to "server"

$ go run main.go




3. In terminal, navigate to "client"

$ go run main.go



4. We will get the reply from server without any problem







**Warning:** Self signed certificates are useful for testing & they should never be used in production. Use a CA signed certificate on production applications.


## Learn More

#### Public Key Cryptography

[Wikipedia - Public Key Cryptography](https://en.wikipedia.org/wiki/Public-key_cryptography)

[IBM - Public Key Cryptography](https://www.ibm.com/support/knowledgecenter/en/SSB23S_1.1.0.13/gtps7/s7pkey.html)

[IBM - Digital Signatures](https://www.ibm.com/support/knowledgecenter/SSB23S_1.1.0.13/gtps7/s7dsign.html)

#### TLS

[Mozilla Guide lines for TLS](https://wiki.mozilla.org/Security/Server_Side_TLS)

[Wikipedia - Public Key Certificate](https://en.wikipedia.org/wiki/Public_key_certificate)

[IBM - An overview of the TLS Handshake](https://www.ibm.com/support/knowledgecenter/en/SSFKSJ_7.1.0/com.ibm.mq.doc/sy10660_.htm)

[Tech Radar - How SSL and TLS Works](https://www.techradar.com/news/software/how-ssl-and-tls-works-1047412)

[The SSL/TLS Handshake: an Overview](https://www.ssl.com/article/ssl-tls-handshake-overview/)

[Caddy](https://caddyserver.com/)
