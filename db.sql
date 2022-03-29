CREATE DATABASE banking;
USE banking;

USE banking;

DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
    `customer_id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `data_of_birth` date NOT NULL,
    `city`  varchar(100) NOT NULL,
    `zipcode` varchar(10) NOT NULL,
    `status` tinyint(1) NOT NULL DEFAULT '1',
    PRIMARY KEY(`customer_id`)
) ENGINE = InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;

INSERT INTO `customers` VALUES
    (2000,'Jason','1997-03-01','Shanghai','210000',1),
    (2001,'Naci','1997-03-02','Shanghai','210000',1),
    (2002,'Archie','1997-03-03','Shanghai','210000',1);
    (2003,'Ben','1997-03-04','Manchester, NH','03102',0),
	(2004,'Nina','1997-03-05','Clarkston, MI','48348',1),
	(2005,'Osman','1997-03-06','Hyattsville, MD','20782',0);

DROP TABLE IF EXISTS `accounts`;

CREATE TABLE `accounts` (
    `account_id` int(11) NOT NULL AUTO_INCREMENT,
    `customer_id` int(11) NOT NULL,
    `opening_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `account_type` varchar(10) NOT NULL,
    `amount` decimal(10,2) NOT NULL,
    -- `pin` varchar(10) NOT NULL,
    `status` tinyint(4) NOT NULL DEFAULT '1',
    PRIMARY KEY(`account_id`),
    KEY `accounts_FK` (`customer_id`),
    CONSTRAINT `accounts_FK` FOREIGN  KEY (`customer_id`) REFERENCES `customers` (`customer_id`)
) ENGINE = InnoDB AUTO_INCREMENT=95471 DEFAULT CHARSET=latin1;

INSERT INTO `accounts` VALUES
    (95470,2000,'2022-03-25 10:20:06','saving',107,1),
    (95471,2002,'2022-03-25 10:21:22','saving',1255,1),
    (95472,2001,'2022-03-25 10:21:22','checking',0723,1);
    (95473,2001,'2022-03-25 10:21:22','saving',0723,1);

DROP TABLE IF EXISTS `transactions`;

CREATE TABLE `transactions`(
    `transaction_id` int(11) NOT NULL AUTO_INCREMENT,
    `account_id` int(11) NOT NULL,
    `amount` int(11) NOT NULL,
    `transaction_type` varchar(10) NOT NULL,
    `transaction_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`transaction_id`),
    KEY `transactions_FK`(`account_id`),
    CONSTRAINT `transactions_FK` FOREIGN KEY (`account_id`) REFERENCES `account` (`account_id`)
)ENGINE = InnoDB DEFAULT CHARSET=latin1;