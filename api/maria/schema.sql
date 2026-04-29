-- phpMyAdmin SQL Dump
-- version 5.2.3
-- https://www.phpmyadmin.net/
--
-- Host: db_mariadb:3306
-- Generation Time: Feb 25, 2026 at 09:24 AM
-- Server version: 10.11.16-MariaDB-ubu2204
-- PHP Version: 8.3.26

--SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
--START TRANSACTION;
--SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `stsy_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `assets`
--

CREATE TABLE IF NOT EXISTS `assets` (
                          `filename` varchar(255) NOT NULL,
                          `created` timestamp NULL DEFAULT current_timestamp(),
                          `creator_user_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `denormalized_alarms_events_live`
--

CREATE TABLE IF NOT EXISTS `denormalized_alarms_events_live` (
                                                   `id` int(11) NOT NULL,
                                                   `modified` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                                                   `source_serial` varchar(255) DEFAULT NULL,
                                                   `device` varchar(255) NOT NULL,
                                                   `devices_alarm_group` varchar(255) NOT NULL,
                                                   `alarm` varchar(255) NOT NULL,
                                                   `message` text DEFAULT NULL,
                                                   `message_alarm_lemma` text DEFAULT NULL,
                                                   `on_ts` timestamp NOT NULL,
                                                   `ack_ts` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `meta`
--

CREATE TABLE IF NOT EXISTS `meta` (
                        `id` int(11) NOT NULL,
                        `modified` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                        `entity_id` int(11) DEFAULT NULL,
                        `entity_table` varchar(255) DEFAULT NULL,
                        `name` varchar(255) DEFAULT NULL,
                        `value` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `permissions_user_asset`
--

CREATE TABLE IF NOT EXISTS `permissions_user_asset` (
                                          `user_id` int(11) NOT NULL,
                                          `modified` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                                          `assets_filename` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `permissions_user_source`
--

CREATE TABLE IF NOT EXISTS `permissions_user_source` (
                                           `modified` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                                           `user_id` int(11) NOT NULL,
                                           `source_serial` varchar(255) NOT NULL,
                                           `read` tinyint(4) NOT NULL DEFAULT 0,
                                           `write` tinyint(4) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `sources`
--

CREATE TABLE IF NOT EXISTS `sources` (
                           `serial` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE IF NOT EXISTS `users` (
                         `id` int(11) NOT NULL,
                         `modified` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                         `username` varchar(255) DEFAULT NULL,
                         `password` varchar(255) DEFAULT NULL,
                         `enabled_bool` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users_pin`
--

CREATE TABLE IF NOT EXISTS `users_pin` (
                             `id` int(11) NOT NULL,
                             `modified` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                             `user_id` int(11) DEFAULT NULL,
                             `pin` varchar(255) DEFAULT NULL,
                             `pin_enabled_bool` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `assets`
--
ALTER TABLE `assets`
    ADD PRIMARY KEY (`filename`),
  ADD KEY `creator_user_id` (`creator_user_id`);

--
-- Indexes for table `denormalized_alarms_events_live`
--
ALTER TABLE `denormalized_alarms_events_live`
    ADD PRIMARY KEY (`id`);

--
-- Indexes for table `meta`
--
ALTER TABLE `meta`
    ADD PRIMARY KEY (`id`),
  ADD KEY `idx_entity` (`entity_id`,`entity_table`);

--
-- Indexes for table `permissions_user_asset`
--
ALTER TABLE `permissions_user_asset`
    ADD PRIMARY KEY (`user_id`,`assets_filename`);

--
-- Indexes for table `permissions_user_source`
--
ALTER TABLE `permissions_user_source`
    ADD PRIMARY KEY (`user_id`,`source_serial`);

--
-- Indexes for table `sources`
--
ALTER TABLE `sources`
    ADD PRIMARY KEY (`serial`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
    ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users_pin`
--
ALTER TABLE `users_pin`
    ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `denormalized_alarms_events_live`
--
ALTER TABLE `denormalized_alarms_events_live`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `meta`
--
ALTER TABLE `meta`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users_pin`
--
ALTER TABLE `users_pin`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `assets`
--
ALTER TABLE `assets`
    ADD CONSTRAINT `assets_ibfk_1` FOREIGN KEY (`creator_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `permissions_user_asset`
--
ALTER TABLE `permissions_user_asset`
    ADD CONSTRAINT `permissions_user_asset_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `permissions_user_source`
--
ALTER TABLE `permissions_user_source`
    ADD CONSTRAINT `permissions_user_source_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `users_pin`
--
ALTER TABLE `users_pin`
    ADD CONSTRAINT `users_pin_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;