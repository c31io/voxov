Install databases.

'''
sudo snap install redis
sudo apt install postgresql
'''

Configure firewall.

'''
sudo ufw deny 6379
sudo ufw deny 5432
'''

Login to postegres.

'''
sudo -i -u postgres
psql
'''

Create database and user.

'''
create database pdb;
create user voxov with encrypted password 'forLocalDevOnly';
grant all privileges on database pdb to voxov;
'''
