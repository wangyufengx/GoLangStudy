### 查看已经开放的端口

```
firewall-cmd --list-ports
```

### 开放端口

```
firewall-cmd --zone=public --add-port=80/tcp --permanent
```

### 重启防火墙

```
systemctl reload firewalld
```
