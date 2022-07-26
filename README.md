## 部署指南


## 本地准备
##### 镜像文件准备

##### 1.镜像文件本地准备
1. 在library/session/manager.go 97行修改cookie的domain ip
2. make docker 生成docker 镜像 形如 harbor.cn/group/fastgo:NO_TAG.79.f80fcea
3. docker save -o bomb.tar harbor.cn/group/fastgo:NO_TAG.79.f80fcea 镜像文件
##### 2. 服务相关文件
1. config.toml为配置文件 需要修改数据库配置以及对方控制系统的配置
```toml
[Server]
Listen = ":19098"
AppName = "fastgo"
Env = "TEST"
[Auth]
Enable = false
Secret = "7ec164764dacaea974eae08966deef87"
Skips = [
]
[Auth.Accounts]
#key = secret
e76cb1d7079fdaf49422fe711bfd98c6 = "e10adc3949ba59abbe56e057f20f883e"

[Database]
Enable = true
UserPassword = "root:123456"
DB = "bomb"
[Database.Write]
HostPort = "tcp(127.0.0.1:33061)"
[Database.Read]
HostPort = "tcp(127.0.0.1:33061)"
[Database.Conn]
MaxLifeTime = 7200
MaxIdle = 50
MaxOpen = 50

[Log]
FilePath = "./log"
FileName = "fastgo"


[Wechat]
ApiKey = "wx12580e7fcadd5e7d"
ApiSecret = "676efdf15192f6524b6bc8e17fdee24b"


[Redis]
Host = "127.0.0.1:6379"
Password = "123456"
```   
2. docker-compose.yml 为部署文件
```yaml
version: '3.2'

services:
  appointment:
    container_name: bomb
    restart: always
    #image: harbor.aibee.cn/ap/parking-dev:0.0.4
    image: harbor.cn/group/fastgo:NO_TAG.78.42ab615
    network_mode: "host"
    logging:
      options:
        max-size: "1g"
        max-file: "2"
    volumes:
      - ./config.toml:/root/config.toml
      - ./log:/root/log
      - ./files:/root/files
    entrypoint:
      - fastgo
      - server
    environment:
      - TZ=Asia/Shanghai
```

##### 3. 数据库相关准备
1. 镜像准备
```
   docker pull mysql:8.0
   docker save -o mysql.tar
```

2. 部署文件
```yaml
version: '3'

services:
  mysql-db-screenserver:
    container_name: mysql-docker-screenserver       # 指定容器的名称
    image: mysql:8.0                   # 指定镜像和版本
    restart: always
    ports:
      - "33061:3306"
    environment:
      MYSQL_ROOT_PASSWORD: 123456
    env_file:
      - .env
    volumes:
      - ./data:/var/lib/mysql"           # 挂载数据目录
      - ./conf.d:/etc/mysql/conf.d"      # 挂载配置文件目录
      - ./my.cnf:/etc/my.conf
```

3. 数据库数据填充
```sql

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `account`
--

DROP TABLE IF EXISTS `account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `account` (
                           `account_id` varchar(255) DEFAULT NULL,
                           `password` varchar(255) DEFAULT NULL,
                           `id` int unsigned NOT NULL AUTO_INCREMENT,
                           `created_at` datetime DEFAULT NULL,
                           `updated_at` datetime DEFAULT NULL,
                           `deleted_at` datetime DEFAULT NULL,
                           PRIMARY KEY (`id`),
                           KEY `idx_account_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account`
--

LOCK TABLES `account` WRITE;
/*!40000 ALTER TABLE `account` DISABLE KEYS */;
INSERT INTO `account` VALUES ('nobody','e10adc3949ba59abbe56e057f20f883e',1,NULL,NULL,NULL);
/*!40000 ALTER TABLE `account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `move_unit`
--

DROP TABLE IF EXISTS `move_unit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `move_unit` (
                             `move_unit_sn` varchar(255) DEFAULT NULL,
                             `soc` int DEFAULT NULL,
                             `status` int DEFAULT NULL,
                             `speed` double DEFAULT NULL,
                             `current_station_code` varchar(255) DEFAULT NULL,
                             `is_in_station` int DEFAULT NULL,
                             `ring_angle` double DEFAULT NULL,
                             `ring_status` int DEFAULT NULL,
                             `work_duration` int DEFAULT NULL,
                             `production_line_id` int DEFAULT NULL,
                             `id` int unsigned NOT NULL AUTO_INCREMENT,
                             `created_at` datetime DEFAULT NULL,
                             `updated_at` datetime DEFAULT NULL,
                             `timestamp` bigint DEFAULT NULL,
                             `move_unit_id` int DEFAULT NULL,
                             `work_status` int DEFAULT '1' COMMENT '0停用 1启用',
                             `deleted` int DEFAULT NULL,
                             PRIMARY KEY (`id`),
                             UNIQUE KEY `move_unit_sn` (`move_unit_sn`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `move_unit`
--

LOCK TABLES `move_unit` WRITE;
/*!40000 ALTER TABLE `move_unit` DISABLE KEYS */;
INSERT INTO `move_unit` VALUES ('ZY-101',19,1,10,'1-缓存工位1',NULL,14.1,NULL,NULL,1,1,NULL,NULL,NULL,NULL,0,NULL),('ZY-102',19,1,10,'1-缓存工位1',NULL,14.1,NULL,NULL,1,2,NULL,NULL,NULL,NULL,0,NULL),('ZY-103',19,1,10,'1-缓存工位1',NULL,14.1,NULL,NULL,1,3,NULL,NULL,NULL,NULL,0,NULL),('ZY-104',19,1,10,'1-缓存工位1',NULL,14.1,NULL,NULL,1,4,NULL,NULL,NULL,NULL,0,NULL),('ZY-105',19,1,10,'1-缓存工位1',NULL,14.1,NULL,NULL,1,5,NULL,NULL,NULL,NULL,0,NULL),('ZY-106',19,1,10,'1-缓存工位1',NULL,14.1,NULL,NULL,1,6,NULL,NULL,NULL,NULL,0,NULL),('ZY-107',19,1,10,'1-缓存工位1',NULL,14.1,NULL,NULL,1,7,NULL,NULL,NULL,NULL,0,NULL),('ZY-108',19,1,10,'1-缓存工位1',NULL,14.1,NULL,NULL,1,8,NULL,NULL,NULL,NULL,0,NULL),('ZY-109',19,1,10,'1-缓存工位1',NULL,14.1,NULL,NULL,1,9,NULL,NULL,NULL,NULL,0,NULL),('ZY-110',19,1,10,'1-缓存工位1',NULL,14.1,NULL,NULL,1,10,NULL,NULL,NULL,NULL,0,NULL),('ZY-201',19,1,10,'2-缓存工位3',NULL,14.1,NULL,NULL,2,11,NULL,NULL,NULL,NULL,1,NULL),('ZY-202',19,1,10,'2-缓存工位3',NULL,14.1,NULL,NULL,2,12,NULL,NULL,NULL,NULL,1,NULL),('ZY-203',19,1,10,'2-缓存工位3',NULL,14.1,NULL,NULL,2,13,NULL,NULL,NULL,NULL,1,NULL),('ZY-204',19,1,10,'2-缓存工位3',NULL,14.1,NULL,NULL,2,14,NULL,NULL,NULL,NULL,1,NULL),('ZY-205',19,1,10,'2-缓存工位3',NULL,14.1,NULL,NULL,2,15,NULL,NULL,NULL,NULL,1,NULL),('ZY-206',19,1,10,'2-缓存工位3',NULL,14.1,NULL,NULL,2,16,NULL,NULL,NULL,NULL,1,NULL),('ZY-207',19,1,10,'2-缓存工位3',NULL,14.1,NULL,NULL,2,17,NULL,NULL,NULL,NULL,1,NULL),('ZY-208',19,1,10,'2-缓存工位3',NULL,14.1,NULL,NULL,2,18,NULL,NULL,NULL,NULL,1,NULL),('ZY-209',19,1,10,'2-缓存工位3',NULL,14.1,NULL,NULL,2,19,NULL,NULL,NULL,NULL,1,NULL),('ZY-210',19,1,10,'2-缓存工位3',NULL,14.1,NULL,NULL,2,20,NULL,NULL,NULL,NULL,1,NULL);
/*!40000 ALTER TABLE `move_unit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `production_lines`
--

DROP TABLE IF EXISTS `production_lines`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `production_lines` (
                                    `production_line_id` int DEFAULT NULL,
                                    `production_line_name` varchar(255) DEFAULT NULL,
                                    `id` int unsigned NOT NULL AUTO_INCREMENT,
                                    `created_at` datetime DEFAULT NULL,
                                    `updated_at` datetime DEFAULT NULL,
                                    `deleted_at` datetime DEFAULT NULL,
                                    PRIMARY KEY (`id`),
                                    KEY `idx_production_lines_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `production_lines`
--

LOCK TABLES `production_lines` WRITE;
/*!40000 ALTER TABLE `production_lines` DISABLE KEYS */;
INSERT INTO `production_lines` VALUES (1,'产线1',1,NULL,NULL,NULL),(2,'产线2',2,NULL,NULL,NULL);
/*!40000 ALTER TABLE `production_lines` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stations`
--

DROP TABLE IF EXISTS `stations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `stations` (
                            `station_code` varchar(255) DEFAULT NULL,
                            `production_line_id` int DEFAULT NULL,
                            `station_id` varchar(255) DEFAULT NULL,
                            `id` int unsigned NOT NULL AUTO_INCREMENT,
                            `created_at` datetime DEFAULT NULL,
                            `updated_at` datetime DEFAULT NULL,
                            `deleted_at` datetime DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `idx_stations_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=125 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stations`
--

LOCK TABLES `stations` WRITE;
/*!40000 ALTER TABLE `stations` DISABLE KEYS */;
INSERT INTO `stations` VALUES ('1-缓存工位1',1,'1',1,NULL,NULL,NULL),('1-工位1',1,'2',2,NULL,NULL,NULL),('1-工位2',1,'3',3,NULL,NULL,NULL),('1-工位3',1,'4',4,NULL,NULL,NULL),('1-工位4',1,'5',5,NULL,NULL,NULL),('1-工位5',1,'6',6,NULL,NULL,NULL),('1-工位6',1,'7',7,NULL,NULL,NULL),('1-工位7',1,'8',8,NULL,NULL,NULL),('1-工位8',1,'9',9,NULL,NULL,NULL),('1-工位9',1,'10',10,NULL,NULL,NULL),('1-工位10',1,'11',11,NULL,NULL,NULL),('1-缓存工位2',1,'12',12,NULL,NULL,NULL),('2-缓存工位3',2,'1',13,NULL,NULL,NULL),('2-工位1',2,'2',14,NULL,NULL,NULL),('2-工位2',2,'3',15,NULL,NULL,NULL),('2-工位3',2,'4',16,NULL,NULL,NULL),('2-工位4',2,'5',17,NULL,NULL,NULL),('2-工位5',2,'6',18,NULL,NULL,NULL),('2-工位6',2,'7',19,NULL,NULL,NULL),('2-工位7',2,'8',20,NULL,NULL,NULL),('2-工位8',2,'9',21,NULL,NULL,NULL),('2-工位9',2,'10',22,NULL,NULL,NULL),('2-工位10',2,'11',23,NULL,NULL,NULL),('2-缓存工位4',2,'12',24,NULL,NULL,NULL);
/*!40000 ALTER TABLE `stations` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
```




#### 远程操作
1. docker load < mysql.tar 
2. docker load < bomb.tar 
3. docker-compose up -d 执行mysql的 docker-compose.yml
4. 修改config.toml中的数据库或者控制系统api的配置
5. docker-compose up -d 执行服务的 docker-compose.yml
