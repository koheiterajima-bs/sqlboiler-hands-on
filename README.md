# 何をやっているか
- SQL Boilerの使い方を学ぶ

## やっていること
- dockerでMySQLのコンテナ作成
`docker run --name mysql-sqlboiler -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=sqlboiler_demo -p 3308:3306 -d mysql:latest`

- MySQLに接続
`docker exec -it mysql-sqlboiler mysql -uroot -ppassword`

- モデルの生成
`sqlboiler mysql`

