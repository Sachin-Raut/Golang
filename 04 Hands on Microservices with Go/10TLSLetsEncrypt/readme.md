

### Build with target OS and Processor Architecture

```
In terminal, navigate to 10TLSLetsEncrypt & execute

$ env GOOS=linux GOARCH=amd64 go build -o https-test .

(this cretes executable file "https-test" in 10TLSLetsEncrypt)

Move this exe file to live server (lets say our live server path is "/home/Projects")

Now save "https-test" in "home/Projects"

Now run this exe file. In terminal navigate to live server path

/home/Projects $ sudo ./https-test

Now in Postman, set method = GET & execute

https://httpstest.metonymie.com

(we will receive the response)

Now lets test our configuration on ssllabs.com ---> test your server ---> enter your hostname (i.e httpstest.metonymie.com)

(we will get overall ratings & other details as well.) 

```

## Learn More

[Go Autocert Package](golang.org/x/crypto/acme/autocert)

[Let's Encrypt](https://letsencrypt.org/)

[Let's Encrypt - How it works](https://letsencrypt.org/how-it-works/)

[Let's Encrypt - Rate Limits](https://letsencrypt.org/docs/rate-limits/)

[dotGo 2016 - Matthew Holt - Go with ACME](https://www.youtube.com/watch?v=KdX51QJWQTA)

[SSL Server Test](https://www.ssllabs.com/ssltest/)

[How To Build Go Executables for Multiple Platforms on Ubuntu 16.04](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04)


