CREATE TABLE IF NOT EXISTS ums_admin_user
(
    id          BIGINT             NOT NULL AUTO_INCREMENT,
    nickname    VARCHAR(25) UNIQUE NOT NULL COMMENT '昵称',
    avatar      VARCHAR(256) COMMENT '头像',
    `password`  VARCHAR(8)         NOT NULL COMMENT '密码(加密后)',
    create_time TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (id)
);