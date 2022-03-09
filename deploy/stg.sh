APP_NAME=webbe

docker load -i ${APP_NAME}.tar
docker rm -f ${APP_NAME}

sudo mkdir /opt/assets

docker run -d --name ${APP_NAME} \
  -e VIRTUAL_HOST="groupbar.me" \
  -e LETSENCRYPT_HOST="groupbar.me" \
  -e LETSENCRYPT_EMAIL="duongptryu@gmail.com" \
  -v /opt/assets/:/tmp \
  --net my-net \
  -p 8080:8080 \
  --expose 8080 \
  ${APP_NAME}