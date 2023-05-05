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

We need to:
0. Put a whole bunch of files in place
1. Place an order that meets the needs of those files
2. Run a script to complete the order
3. Have the program pick up the completed order
4. Initiate the archive process -- copy to archive
5. When copy completes, double check for the existence and equality of files
6. delete the files locally
7. Place a second order that relies on files in the archive
8. initiate the rehydration of the files in the archive
9. wait in the background until complete (this won't finish)
10. start the order when that completes
