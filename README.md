# KIDS23-Team1

## install postgres11:
```
sudo yum update -y
sudo reboot
sudo yum install https://download.postgresql.org/pub/repos/yum/reporpms/EL-7-x86_64/pgdg-redhat-repo-latest.noarch.rpm
sudo yum -y install postgresql11-server postgresql11
sudo /usr/pgsql-11/bin/postgresql-11-setup initdb
sudo systemctl start postgresql-11
sudo systemctl enable postgresql-11
```

## install migrate
```
wget https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz
tar -xf migrate.linux-amd64.tar.gz
mv migrate /usr/local/bin
```

## setup adminuser
```
sudo su postgres
createuser fileadmin
psql -c "alter user fileadmin with password 'test1234'"
psql -c "alter user fileadmin createdb"
```


