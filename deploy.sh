docker login -u $DOCKER_USER -p $DOCKER_PASSWORD
cd ./users
docker build -t slack-users .
docker tag slack-users $DOCKER_USER/users
docker push $DOCKER_USER/users
cd ..
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker pull $DOCKER_USER/users
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker stop slack-users || true
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker rm slack-users || true
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker run -name=slack-users -d -net=host $DOCKER_USER/users


cd ./chats
docker build -t slack-chats .
docker tag slack-chats $DOCKER_USER/chats
docker push $DOCKER_USER/chats
cd ..
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker pull $DOCKER_USER/chats
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker stop slack-chats || true
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker rm slack-chats || true
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker run --name=slack-chats -d  -net=host $DOCKER_USER/chats


cd ./notifications
docker build -t slack-notifications .
docker tag slack-notifications $DOCKER_USER/notifications
docker push $DOCKER_USER/notifications
cd ..
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker pull $DOCKER_USER/notifications
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker stop slack-notifications || true
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker rm slack-notifications || true
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker run --name=slack-notifications -d -net=host $DOCKER_USER/notifications


cd ./messages
docker build -t slack-messages .
docker tag slack-messages $DOCKER_USER/messages
docker push $DOCKER_USER/messages
cd ..
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker pull $DOCKER_USER/messages
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker stop slack-messages || true
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker rm slack-messages || true
ssh -i $HOME/.ssh/2019_2_CoolCode_id_rsa.pem ubuntu@95.163.209.195 sudo docker run --name=slack-messages -d -net=host $DOCKER_USER/messages
