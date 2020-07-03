# binder-api

## JWTトークン用の鍵生成
### 秘密鍵の生成
```
$ openssl genrsa 1024 > ./rsa/private-key.pem
```
### 公開鍵の生成
```
$ openssl rsa -in ./rsa/private-key.pem -pubout -out ./rsa/public-key.pem
```

## Install MINIO
### macOS
```
$ brew install minio
$ brew services start minio
```
するとどっかにアクセスキーが生成される  
(自分の場合は、/usr/local/var/log/minio.logに入っていた )

### Linux
```
$ wget https://dl.min.io/server/minio/release/linux-amd64/minio
$ chmod +x minio
$ ./minio server /data
```
起動すると、アクセスキー等が表示されるはず

## nginxでminioのリバプロ周辺
作業予定中