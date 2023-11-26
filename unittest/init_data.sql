USE `hr`;

LOCK TABLES `admins` WRITE;
INSERT INTO `admins`
VALUES (1,'虎王',0,'{\"avatar\":\"https://my.scutbot.cn/storage/avatars/users/8af90aa5190041148cb21de3807e56ac.jpg\",\"created_at\":\"2020-06-18T11:21:50.000000Z\",\"email\":\"i@scut-robotlab.club\",\"group_info\":[\"机械\",\"电控\",\"视觉\",\"项目\",\"信息\",\"宣运\",\"教育\",\"模拟器\",\"信息模拟器\"],\"groups\":[{\"Id\":1,\"Name\":\"机械\",\"Power\":2},{\"Id\":2,\"Name\":\"电控\",\"Power\":2},{\"Id\":3,\"Name\":\"视觉\",\"Power\":2},{\"Id\":6,\"Name\":\"项目\",\"Power\":2},{\"Id\":7,\"Name\":\"信息\",\"Power\":2},{\"Id\":8,\"Name\":\"宣运\",\"Power\":2},{\"Id\":9,\"Name\":\"教育\",\"Power\":2},{\"Id\":10,\"Name\":\"模拟器\",\"Power\":2},{\"Id\":13,\"Name\":\"信息模拟器\",\"Power\":2}],\"id\":1,\"name\":\"虎王\",\"updated_at\":\"2022-10-18T04:18:44.000000Z\"}',1,0,NULL,current_timestamp(),current_timestamp());
UNLOCK TABLES;

LOCK TABLES `announce_configs` WRITE;
INSERT INTO `announce_configs`
VALUES (1, 'HaveNotAppliedForm', '您还没有提交报名表。'),
       (2, 'HaveNotSelectedTime', '{{applicant.name}}同学您好，感谢您报名华工机器人实验室{{#each intents}}{{this}}组{{#unless @last}}、{{/unless}}{{/each}}。您还没有选择报名面试时间，请进入系统选择。'),
       (3, 'HaveNotInterview', '{{applicant.name}}同学您好，感谢您报名华工机器人实验室。您报名的{{#each times}}{{this.group}}组的面试时间为{{this.date}} {{this.time}}，面试地点在{{this.location}}；{{/each}}我们期待您的到来！'),
       (4, 'HaveNotPublished', '{{applicant.name}}同学您好，面试结果尚未公布，请静候通知。'),
       (5, 'Admitted', '恭喜您！{{applicant.name}}同学，您通过了{{#each admitted}}{{this.group}}组{{#unless @last}}、{{/unless}}{{/each}}面试！'),
       (6, 'Failed', '很遗憾，{{applicant.name}}同学，您没有通过面试。');
UNLOCK TABLES;

LOCK TABLES `forms` WRITE;
INSERT INTO `forms`
VALUES (1, 'E-mail', 'email', 'text-field-email', 1, '[]', NULL, NULL, NULL, current_timestamp(), current_timestamp()),
       (2, '微信号', 'wechat', 'text-field', 1, '[]', NULL, NULL, NULL, current_timestamp(), current_timestamp()),
       (3, '年级', 'grade', 'select', 1, '[\"大一\", \"大二\", \"大三\", \"大四\", \"其他\"]', NULL, NULL, NULL, current_timestamp(), current_timestamp()),
       (4, '学院', 'school', 'select', 1, '[\"机械与汽车工程学院\", \"建筑学院\", \"土木与交通学院\", \"电子与信息学院\", \"材料科学与工程学院\", \"化学与化工学院\", \"轻工科学与工程学院\", \"食品科学与工程学院\", \"数学学院\", \"物理与光电学院\", \"经济与金融学院\", \"旅游管理系\", \"电子商务系\", \"自动化科学与工程学院\", \"计算机科学与工程学院\", \"电力学院\", \"生物科学与工程学院\", \"环境与能源学院\", \"软件学院\", \"工商管理学院（创业教育学院）\", \"公共管理学院\", \"马克思主义学院\", \"外国语学院\", \"法学院（知识产权学院）\", \"新闻与传播学院\", \"艺术学院\", \"体育学院\", \"设计学院\", \"医学院（生命科学研究院）\", \"国际教育学院\", \"峻德书院\", \"铭诚书院\", \"生物医学科学与工程学院\", \"吴贤铭智能工程学院\", \"前沿软物质学院\", \"微电子学院\", \"未来技术学院\", \"海洋科学与工程学院\"]', NULL, NULL, NULL, current_timestamp(), current_timestamp()),
       (5, '专业', 'major', 'text-field', 1, '[]', NULL, 16, NULL, current_timestamp(), current_timestamp()),
       (6, '本学年校区', 'campus', 'select', 1, '[\"五山校区\", \"大学城校区\", \"广州国际校区\"]', NULL, NULL, NULL, current_timestamp(), current_timestamp()),
       (7, '自我介绍', 'self_intro', 'textarea', 0, '[]', NULL, 3000, NULL, current_timestamp(), current_timestamp()),
       (8, '促使你来报名的最主要原因', 'apply_main_reason', 'select', 1, '[\"同学分享\", \"老师介绍\", \"微信公众号@华工机器人实验室\", \"bilibili@华工机器人实验室\", \"RoboMaster赛事\", \"中小学阶段的机器人竞赛经历\", \"校内线下招新\", \"华工机器人协会\", \"华工机器人协会校内赛\", \"不记得\", \"其它渠道\"]', NULL, NULL, NULL, current_timestamp(), current_timestamp()),
       (9, '内部推荐人（选填）', 'internal_referees', 'text-field', 0, '[]', NULL, 16, NULL, current_timestamp(), current_timestamp());
UNLOCK TABLES;

LOCK TABLES `guide` WRITE;
INSERT INTO `guide`
VALUES (1, '赛规', '{\"steps\": [{\"title\": \"#1 官方规则视频\", \"content\": [{\"href\": \"https://www.bilibili.com/video/BV1cs4y1B787\", \"text\": \"RoboMaster 2023 机甲大师超级对抗赛·规则视频\", \"type\": \"link\"}], \"subtitle\": \"请花费2分钟观看以下视频，这将帮助你概括性地了解比赛流程。30%的赛规问题可以从中找到答案。\"}, {\"title\": \"#2 真实比赛视频\", \"content\": [{\"href\": \"https://www.bilibili.com/video/BV1GP41147Gf\", \"text\": \"第69场 东北大学 TDT战队 vs 华南理工大学 华南虎战队\", \"type\": \"link\"}], \"subtitle\": \"请花费20分钟观看以下视频，这将帮助你完整理解比赛的全流程。60%的赛规问题可以从中找到答案。\"}, {\"title\": \"#3 官方规则手册\", \"content\": [{\"href\": \"https://terra-1-g.djicdn.com/b2a076471c6c4b72b574a977334d3e05/RM2023/RoboMaster%202023%20%E6%9C%BA%E7%94%B2%E5%A4%A7%E5%B8%88%E8%B6%85%E7%BA%A7%E5%AF%B9%E6%8A%97%E8%B5%9B%E6%AF%94%E8%B5%9B%E8%A7%84%E5%88%99%E6%89%8B%E5%86%8CV2.0%EF%BC%8820230704%EF%BC%89.pdf\", \"text\": \"RoboMaster 2023 机甲大师超级对抗赛 比赛规则手册V2.0\", \"type\": \"link\"}], \"subtitle\": \"有10%的题目需要查阅规则手册找到答案，请善用浏览器的搜索功能。\"}], \"subtitle\": \"不同于技术方面的考察，赛规题目只需要很少的准备时间即可获得满分。阅读下方的题库指导，你将在40分钟内掌握全部的赛规题目。\"}', NULL, current_timestamp(), current_timestamp());
UNLOCK TABLES;

LOCK TABLES `rooms` WRITE;
INSERT INTO `rooms`
VALUES (1, '31-501', '', 0, current_timestamp(), 0, '', '', '电控', 0, NULL, current_timestamp(), current_timestamp()),
       (2, '31-502', '', 0, current_timestamp(), 0, '', '', '机械', 0, NULL, current_timestamp(), current_timestamp()),
       (3, '31-503', '', 0, current_timestamp(), 0, '', '', '视觉', 0, NULL, current_timestamp(), current_timestamp()),
       (4, '31-504', '', 0, current_timestamp(), 0, '', '', '宣运', 0, NULL, current_timestamp(), current_timestamp());
UNLOCK TABLES;
