# 生成普通证书

# ca
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=example Inc./CN=apulis' -keyout rootCA.key -out rootCA.crt
# server
openssl req -out server.csr -newkey rsa:2048 -nodes -keyout server.key -subj "/CN=apedge.apulis.cn/O=httpbin organization"
openssl x509 -req -sha256 -days 365 -CA rootCA.crt -CAkey rootCA.key -set_serial 0 -in server.csr -out server.crt
# client
openssl req -out client.csr -newkey rsa:2048 -nodes -keyout client.key -subj "/CN=apedge.apulis.cn/O=httpbin's client organization"
openssl x509 -req -sha256 -days 365 -CA rootCA.crt -CAkey rootCA.key -set_serial 0 -in client.csr -out client.crt

## server
rm /etc/kubeedge/ca/*
rm /etc/kubeedge/certs/*
cp rootCA.* /etc/kubeedge/ca
cp server.* /etc/kubeedge/certs

## client
cd /etc/kubeedge
export SSL_PATH=ssl-test
scp -r root@192.168.3.149:/root/KubeEdge/${SSL_PATH} /etc/kubeedge

rm /etc/kubeedge/ca/*
rm /etc/kubeedge/certs/*
cp ${SSL_PATH}/rootCA.crt ca
cp ${SSL_PATH}/client.crt certs/server.crt
cp ${SSL_PATH}/client.key certs/server.key


# 生成SAN证书
# ca
openssl genrsa -des3 -out rootCA.key 2048
openssl req -sha256 -new -x509 -days 365 -key rootCA.key -out rootCA.crt \
    -subj "/C=CN/ST=GD/L=SZ/O=lee/OU=study/CN=apulis"

# server
openssl genrsa -des3 -out server.key 2048

openssl req -new \
    -sha256 \
    -key server.key \
    -subj "/C=CN/ST=GD/L=SZ/O=lee/OU=study/CN=apedge.apulis.cn" \
    -reqexts SAN \
    -config <(cat ./openssl.cnf \
        <(printf "[SAN]\nsubjectAltName=DNS:*.apulis.cn")) \
    -out server.csr

mkdir newcerts
touch index.txt
echo 00 > serial
openssl ca -in server.csr \
    -md sha256 \
    -keyfile rootCA.key \
    -cert rootCA.crt \
    -extensions SAN \
    -config <(cat ./openssl.cnf \
        <(printf "[SAN]\nsubjectAltName=DNS:*.apulis.cn")) \
    -out server.crt

# server bash
rm /etc/kubeedge/ca/*
rm /etc/kubeedge/certs/*
tree /etc/kubeedge
cp rootCA.* /etc/kubeedge/ca
cp server.* /etc/kubeedge/certs
tree /etc/kubeedge


cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=kubernetes kubeedge-csr.json | cfssljson -bare edge