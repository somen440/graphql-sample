-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- ホスト: db
-- 生成日時: 2020 年 4 月 05 日 05:42
-- サーバのバージョン： 5.7.29-log
-- PHP のバージョン: 7.4.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- データベース: `test_database`
--

-- --------------------------------------------------------

--
-- テーブルの構造 `channel`
--

CREATE TABLE `channel` (
  `channel_id` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Channel manage table';

--
-- テーブルのデータのダンプ `channel`
--

INSERT INTO `channel` (`channel_id`, `name`, `created_at`, `updated_at`) VALUES
('787bc3a7-cc8d-4700-ba1e-993acad412cf', 'Tech チャンネル', '2020-04-05 14:21:13', '2020-04-05 14:21:13'),
('cf6441b9-6a1a-49a7-ba53-941d8d030188', 'News チャンネル', '2020-04-05 14:21:13', '2020-04-05 14:21:13'),
('ebf8cb44-8ed0-45e6-a60c-d775143aa70d', 'Cook チャンネル', '2020-04-05 14:21:13', '2020-04-05 14:21:13');

-- --------------------------------------------------------

--
-- テーブルの構造 `followed_channels`
--

CREATE TABLE `followed_channels` (
  `followed_channel_id` int(11) NOT NULL,
  `listener_id` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `channel_id` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Followed management table';

--
-- テーブルのデータのダンプ `followed_channels`
--

INSERT INTO `followed_channels` (`followed_channel_id`, `listener_id`, `channel_id`, `created_at`, `updated_at`) VALUES
(1, '82ad9338-18c9-4700-83c2-b6904052b7d8', '787bc3a7-cc8d-4700-ba1e-993acad412cf', '2020-04-05 14:23:26', '2020-04-05 14:23:26'),
(2, '82ad9338-18c9-4700-83c2-b6904052b7d8', 'cf6441b9-6a1a-49a7-ba53-941d8d030188', '2020-04-05 14:23:26', '2020-04-05 14:23:26'),
(3, '024127fb-5902-49a6-862d-a4981644f940', '787bc3a7-cc8d-4700-ba1e-993acad412cf', '2020-04-05 14:23:26', '2020-04-05 14:23:26'),
(4, '024127fb-5902-49a6-862d-a4981644f940', 'cf6441b9-6a1a-49a7-ba53-941d8d030188', '2020-04-05 14:23:26', '2020-04-05 14:23:26'),
(5, '024127fb-5902-49a6-862d-a4981644f940', 'ebf8cb44-8ed0-45e6-a60c-d775143aa70d', '2020-04-05 14:23:26', '2020-04-05 14:23:26'),
(6, 'd596715b-76fc-11ea-9935-0242ac160002', 'ebf8cb44-8ed0-45e6-a60c-d775143aa70d', '2020-04-05 14:23:26', '2020-04-05 14:23:26');

-- --------------------------------------------------------

--
-- テーブルの構造 `listener`
--

CREATE TABLE `listener` (
  `listener_id` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='Listener management table';

--
-- テーブルのデータのダンプ `listener`
--

INSERT INTO `listener` (`listener_id`, `name`, `created_at`, `updated_at`) VALUES
('024127fb-5902-49a6-862d-a4981644f940', 'うえの', '2020-04-05 14:19:19', '2020-04-05 14:19:19'),
('82ad9338-18c9-4700-83c2-b6904052b7d8', 'よしむら', '2020-04-05 14:19:19', '2020-04-05 14:19:19'),
('d596715b-76fc-11ea-9935-0242ac160002', 'ながた', '2020-04-05 14:18:07', '2020-04-05 14:18:07');

--
-- ダンプしたテーブルのインデックス
--

--
-- テーブルのインデックス `channel`
--
ALTER TABLE `channel`
  ADD PRIMARY KEY (`channel_id`);

--
-- テーブルのインデックス `followed_channels`
--
ALTER TABLE `followed_channels`
  ADD PRIMARY KEY (`followed_channel_id`),
  ADD UNIQUE KEY `listener_id` (`listener_id`,`channel_id`),
  ADD KEY `channel_id` (`channel_id`);

--
-- テーブルのインデックス `listener`
--
ALTER TABLE `listener`
  ADD PRIMARY KEY (`listener_id`);

--
-- ダンプしたテーブルのAUTO_INCREMENT
--

--
-- テーブルのAUTO_INCREMENT `followed_channels`
--
ALTER TABLE `followed_channels`
  MODIFY `followed_channel_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- ダンプしたテーブルの制約
--

--
-- テーブルの制約 `followed_channels`
--
ALTER TABLE `followed_channels`
  ADD CONSTRAINT `channel_id` FOREIGN KEY (`channel_id`) REFERENCES `channel` (`channel_id`),
  ADD CONSTRAINT `listener_id` FOREIGN KEY (`listener_id`) REFERENCES `listener` (`listener_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
