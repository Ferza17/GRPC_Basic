#USE GIT BASH if u use Windows

#change these CN's to match your hosts in your environment if needed
SERVER_ON=localhost

#step 1: Generate Certificate Authority + Trust Certificate (ca.cert)
openssl genrsa -passout pass:1111 -des3 -out ca.key 4096
openssl req -passin pass:1111 -new -x509 -days 365 -key ca.key -out ca.crt -subj "//ON=${SERVER_ON}"

#Step 2: Generate the server Private Key (server.key)
openssl genrsa  -passout pass:1111 -des3 -out server.key 4096

#Step 3: Get a certificate signing Request from the CA (server.csr)
openssl req -passin pass:1111 -new -key server.key -out server.csr -subj "//ON=${SERVER_ON}"

#step 4: Sign the certificate with the CA we created (it's called self signing) - server.crt
openssl x509 -req -passin pass:1111 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt

#Step 5: Convert the server certificate to /pem format (server.pem) - usable by gRPC
openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem

