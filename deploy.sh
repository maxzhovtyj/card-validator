pm2 stop card-validator
pm2 start card-validator-linux-amd64 --name=card-validator -- -configPath=./config.yml