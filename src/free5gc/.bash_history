cd free5gc-stage-1/
ls
vim lib/nextepc/core/src/Makefile.am 
vim lib/nextepc/core/test/Makefile.am 
vim lib/sbiJson/Makefile.am 
cat src/app/Makefile.am
vim src/app/Makefile.am
vim src/smf/Makefile.am 
vim src/amf/Makefile.am 
vim src/smf/Makefile.am 
vim src/amf/Makefile.am 
vim src/smf/Makefile.am 
vim src/amf/Makefile.am 
vim test/Makefile.am 
top
ps aux | grep free
killall -9 free5gc
ps aux | grep free
kill -9 4769
kill -9 4784
ps aux | grep free
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan --debug apply
ifc
ping 8.8.8.8
tail -f /var/log/syslog 
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan --debug apply
tail -f /var/log/syslog 
ping 8.8.8.8

ping 8.8.8.8
sudo apt-get update --fix-missing
sudo apt-get upgrade
sudo apt-get dist-upgrade 
ping 8.8.8.8
sudo apt-get dist-upgrade 
sudo apt-get upgrade
sudo apt-get update 
sudo apt-get upgrade
sudo apt-get dist-upgrade 
sudo apt-get install  net-tools 
sudo apt-get install mongodb wget git -y
sudo systemctl start mongodb
sudo apt-get install golang
go get -u -v "github.com/gorilla/mux"
go get -u -v "github.com/x/net/http2"
go get -u -v "github.com/gorilla/mux"
go get -u -v "github.com/x/net/http2"
go get -u -v "golang.org/x/net/http2"
go get -u -v "golang.org/x/sys/unix"
ls -al /dev/net/tun 
sudo sh -c "cat << EOF > /etc/systemd/network/99-free5gc.netdev
[NetDev]
Name=uptun
Kind=tun
EOF"
sudo systemctl enable systemd-networkd
sudo systemctl restart systemd-networkd
ifconfig 
ifconfig  -a
sysctl -n net.ipv6.conf.uptun.disable_ipv6
sudo sh -c "cat << EOF > /etc/systemd/network/99-free5gc.network
[Match]
Name=uptun
[Network]
Address=45.45.0.1/16
Address=cafe::1/64
EOF"
sudo systemctl enable systemd-networkd
sudo systemctl restart systemd-networkd
ifconfig 
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan --dubug apply
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan --dubug apply
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan --dubug apply
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan --dubug apply
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan --dubug apply
sudo nano /etc/netplan/50-cloud-init.yaml 
ifconfig 
sudo systemctl enable systemd-networkd
sudo systemctl disable systemd-networkd
sudo systemctl enable systemd-networkd
sudo systemctl restart systemd-networkd
ifconfig 
sudo apt-get -y install autoconf libtool gcc pkg-config git flex bison libsctp-dev libgnutls28-dev libgcrypt-mingw-w64-dev libssl-dev libidn11-dev libmongoc-dev  libbson-dev libyaml-dev 
git clone https://bitbucket.org/nctu_5g/free5gc-stage-1.git
cd free5gc-stage-1/
sudo nano env_setup.sh
chmod a+x env_setup.sh 
sudo chmod a+x env_setup.sh 
sudo ./env_setup.sh 
ifconfig 
sudo ./env_setup.sh 
sudo nano env_setup.sh
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan --debug apply
ifconfig 
sudo nano env_setup.sh
sudo ./env_setup.sh 
ifconfig 
cat /proc/sys/net/ipv4/ip_forward
autoreconf -iv
./configure --prefix=`pwd`/install
sudo apt-get install libgcrypt-dev
./configure --prefix=`pwd`/install
make -j `nproc`
nano lib/freeDiameter-1.2.1/libfdcore/sctp.c 
sudo apt-get install vim
vim lib/freeDiameter-1.2.1/libfdcore/sctp.c 
nano lib/freeDiameter-1.2.1/libfdcore/sctp.c 
make -j `nproc`
nano lib/ipfw/dummynet.c 
nano lib/ipfw/Makefile.am 
autoreconf -iv
./configure --prefix=`pwd`/install
make -j `nproc`
nano lib/nextepc/Makefile.am 

make -j `nproc`
nano lib/nextepc/Makefile
vim lib/nextepc/Makefile
cd lib/nextepc/
grep -rn "Werror"
make -j `nproc`
cd ../..
./configure --prefix=`pwd`/install
make -j `nproc`
./configure --prefix=`pwd`/install
make -j `nproc`
./configure --prefix=`pwd`/install
make -j `nproc`
./configure --prefix=`pwd`/install
make -j `nproc`
./configure --prefix=`pwd`/install
make -j `nproc`
./configure --prefix=`pwd`/install
make -j `nproc`
./configure --prefix=`pwd`/install
make -j `nproc`
./configure --prefix=`pwd`/install
make -j `nproc`
./configure --prefix=`pwd`/install
make -j `nproc`
./configure --prefix=`pwd`/install
make -j `nproc`
./configure --prefix=`pwd`/install
make -j `nproc`
./configure --prefix=`pwd`/install && make -j `nproc`
make install
./free5gc-ngcd 
ifconfig 
sudo ./env_setup.sh 
ifconfig 
./test/testngc -f install/etc/free5gc/test/free5gc.testngc.conf 
ifconfig 
sudo ./env_setup.sh 
ifconfig 
init 0
cd free5gc-stage-1/
sudo ./env_setup.sh 
ls
ifconfig 
./free5gc-ngcd 
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan --debug apply
tail -f /var/log/syslog 
ping 8.8.8.8
tail -f /var/log/syslog 
ping 8.8.8.8
tail -f /var/log/syslog 
sudo nano /etc/netplan/50-cloud-init.yaml 
tail -f /var/log/syslog 
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan --debug apply
ifconfig 
tail -f /var/log/syslog 
init 6
ping 10.0.0.146
ifconfig 
telnet 10.0.0.2
sudo netplan apply
telnet 192.188.2.4
ifconfig 
tail -f /var/log/syslog 
sudo tcpdump -i eth0
sudo tcpdump -i eth0 -vv
sudo tcpdump -i eth0 
ifconfig 
nano free5gc-stage-1/install/etc/free5gc/free5gc.conf 
sudo nano free5gc-stage-1/env_setup.sh 
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo tcpdump -i eth0
ifconfig 
cd free5gc-stage-1/
cd install/
ls
cd etc/free5gc/
ls
nano free5gc.conf 
top
ifconfig 
cd ../..
cd ..
ls
sudo nano env_setup.sh 
sudo ./env_setup.sh 
ifconfig 
sudo ./env_setup.sh 
ifconfig 
sudo iptables-S
sudo iptables -S
sudo iptables -F
ifconfig 
sudo ./env_setup.sh 
top
ifconfig 
top
ifconfig 
sudo nano env_setup.sh 
nano install/etc/free5gc/free5gc.conf 
ifconfig 
sudo ./env_setup.sh 
ifconfig 
sudo ifconfig eth0 0.0.0.0
sudo ./env_setup.sh 
sudo iptables -F
sudo ./env_setup.sh 
ifconfig 
cat env_setup.sh 
ip addr 
sudo ifconfig eth0 0.0.0.0
ifconfig 
sudo ifconfig eth0 10.0.0.0 netmask 255.255.255.0
ifconfig 
sudo ifconfig eth0 0.0.0.0
sudo ifconfig eth0 10.0.0.2 netmask 255.255.255.0
ifconfig 
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan apply
sudo nano /etc/netplan/50-cloud-init.yaml 
nano install/etc/free5gc/free5gc.conf 
sudo nano env_setup.sh 
ifconfig 
sudo netplan apply 
ifconfig 
top
cat /proc/net/sctp/eps 
clear
top
cat /proc/net/sctp/eps 
ifconfig 
cat /proc/net/sctp/eps 
webui/
ls
cd webui/
l
ls
npm install
sudo apt-get install npm
npm install
npm run dev
ifconfig 
tail -f /var/log/syslog 
killall wpa_supplicant 
sudo killall wpa_supplicant 
sudo netplan apply
tail -f /var/log/syslog 
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan --debug try
sudo netplan apply
sudo nano /etc/netplan/50-cloud-init.yaml 
tail -f /var/log/syslog 
ifconfig 
sudo killall wpa_supplicant 
tail -f /var/log/syslog 
ifconfig 
tail -f /var/log/syslog 
sudo systemctl restart netplan-wpa@wlan0.service
tail -f /var/log/syslog 
ifconfig 
tail -f /var/log/syslog 
sudo wpa_cli wlan0
sudo wpa_cli -i wlan0
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan apply
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan apply
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan ap
sudo systemctl enable network-manager
sudo systemctl enable NetworkManager.service
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan apply
sudo systemctl enable NetworkManager.service
history
history | grep disable
sudo systemctl enable systemd-networkd
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan apply
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan apply
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan apply
sudo systemctl disable systemd-networkd
sudo netplan apply
sudo systemctl restart netplan-wpa@wlan0.service
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan apply
sudo ifconfig wlan0 down
sudo ifconfig wlan0 up
sudo systemctl restart netplan-wpa@wlan0.service
cat /etc/hosts
sudo nano /etc/hosts
sudo nano /etc/hostsna
sudo nano /etc/hostname 
sudo systemctl restart netplan-wpa@wlan0.service
ifconfig 
ifconfig -a
sudo ifconfig eth1 up
ifconfig -a
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan apply
ping 8.8.8.8
sudo ifconfig wlan0 down
ping 8.8.8.8
cd free5gc-stage-1/
sudo ./env_setup.sh 
./free5gc-ngcd 
ifconfig 
./free5gc-ngcd 
ps aux | grep 
lsmod sctp
lsmod
lsmod | grep sctp
sudo modprobe -r sctp
lsmod
./free5gc-ngcd 
ifconfig 
./free5gc-ngcd 
init 0
tail -f /var/log/syslog
sudo tcpdump -i eth0
cd free5gc-stage-1/
sudo ./env_setup.sh 
ifconfig 
ifconfig 
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan apply
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan apply
sudo nano /etc/netplan/50-cloud-init.yaml 
sudo netplan apply
ifconfig 
cd free5gc-stage-1/
sudo nano env_setup.sh 
ifconfig 
./free5gc-ngcd 
ifconfig 
telnet 192.188.2.4
./free5gc-ngcd 
init 0
ping 8.8.8.8
ifconfig
dhclient eth0
sudo dhclient eth0
ifconfig
sudo dhclient eth0
ping 8.8.8.8
ls
ifconfig
sudo dhclient eth0
df -h
ls
curl ifconfig.co
lscpu
git clone https://ipadaxe@bitbucket.org/free5gc-team/gofree5gc.git
ls
lsblk
cd /media/
ls
sudo mkdir usb
sudo mount /dev/sda usb
cd usb/
ls
cp gofree5gc_old.zip ~
cd ..
sudo umount usb
cd ~
ls
unzip gofree5gc_old.zip .
sudo apt install unzip
ping 8.8.8.8
sudo dhclient eth0 
ping 8.8.8.8
sudo apt install unzip
ping 8.8.8.8
sudo apt update
sudo apt install unzip
unzip gofree5gc_old.zip 
ls
rm -rf gofree5gc
rm gofree5gc_old gofree5gc
mv gofree5gc_old gofree5gc
cd gofree5gc/
git branch
git checkout Release-v2.0.2
git checkout go.sum
git checkout Release-v2.0.2
git branch 
ls
cd infra/
ls
cd ..
cd infra
ls
vim git-script.sh 
vim pre-commit 
cd ..
ls
cd /media/
ls
lsblk
sudo mount /dev/sda usb
cd usb/
ls
cp release.sh ~
cd ~
ls
vim release.sh
cp release.sh roger_release.sh
ls
go
go version
which go
vim .bashrc 
vim .profile 
vim .profile ls
ls
ls -a
vim .bashrc 
vim .profile 
go env
vim .bashrc 
which go
ls
cd go
ls
cd ..
vim /etc/profile
vim /etc/bash.bashrc 
vim .bashrc 
go version
which go
vim .profile 
vim .bash_logout 
vim .bash_history 
ls
cd free5gc-stage-1/
ls
which go
cd ..
vim .bashrc 
vim /etc/profile
vim /etc/profile.d/bash_completion.sh 
vim /etc/profile.d/01-locale-fix.sh 
ls
ls -a
vim .profile 
cd /usr/bin/
ls
ls go
go
cd ..
ls
la
ls
go env
cd ~
ls
echo $GOPATH
echo $GOROOT
echo GOROOT
uname -r
uname -a
go version
wget https://dl.google.com/go/go1.12.9.linux-arm64.tar.gz
ls
udo tar -C /usr/local/ -zxvf go1.12.9.linux-arm64.tar.
udo tar -C /usr/local/ -zxvf go1.12.9.linux-arm64.tar.gz
sudo tar -C /usr/local/ -zxvf go1.12.9.linux-arm64.tar.gz
vim .bashrc 
source .bashrc 
go version
echo $GOPATH
which go
cd /usr/bin/
ls go*
sudo mv go go_old
go version
echo $PATH | grep go
go version
logout
go version
ls
gofmt version
gofmt
gofmt -ve
gofmt -v
which gofmt
cd /usr/bin/
ls
sudo mv gofmt gofmt_old
logout
