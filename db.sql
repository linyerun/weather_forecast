/*
SQLyog Ultimate v12.08 (64 bit)
MySQL - 5.7.20 : Database - weather_system
*********************************************************************
*/

CREATE DATABASE /*!32312 IF NOT EXISTS*/`weather_system` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `weather_system`;

/*Table structure for table `city_msg` */

CREATE TABLE `city_msg` (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                            `city` varchar(50) DEFAULT NULL,
                            `citykey` varchar(50) DEFAULT NULL,
                            `parent` varchar(50) DEFAULT NULL,
                            `update_time` varchar(10) DEFAULT NULL,
                            `shidu` varchar(10) DEFAULT NULL,
                            `pm25` varchar(10) DEFAULT NULL,
                            `pm10` varchar(10) DEFAULT NULL,
                            `quality` varchar(10) DEFAULT NULL,
                            `wendu` varchar(10) DEFAULT NULL,
                            `ganmao` varchar(50) DEFAULT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `forecast_weather` */

CREATE TABLE `forecast_weather` (
                                    `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                    `city_id` bigint(20) DEFAULT NULL,
                                    `ymd` datetime(3) DEFAULT NULL,
                                    `date` varchar(20) DEFAULT NULL,
                                    `high` varchar(10) DEFAULT NULL,
                                    `low` varchar(10) DEFAULT NULL,
                                    `week` varchar(10) DEFAULT NULL,
                                    `sunrise` varchar(20) DEFAULT NULL,
                                    `sunset` varchar(20) DEFAULT NULL,
                                    `aqi` bigint(20) DEFAULT NULL,
                                    `fx` varchar(10) DEFAULT NULL,
                                    `fl` varchar(10) DEFAULT NULL,
                                    `type` varchar(10) DEFAULT NULL,
                                    `notice` varchar(50) DEFAULT NULL,
                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4;