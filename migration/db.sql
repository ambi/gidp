CREATE DATABASE IF NOT EXISTS gidp
    CHARACTER SET utf8
    COLLATE utf8_general_ci;

USE gidp;

-- DROP TABLE users;
-- DROP TABLE tenants;

CREATE TABLE IF NOT EXISTS tenants (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    status VARCHAR(100) NOT NULL DEFAULT('active')
);

CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    tenant_id VARCHAR(36) NOT NULL,
    display_name VARCHAR(1000) NOT NULL,
    INDEX (tenant_id),
    FOREIGN KEY (tenant_id) REFERENCES tenants(id) ON DELETE CASCADE
);

-- SET @tenant_uuid1 = UUID(),
--     @tenant_uuid2 = UUID(),
--     @user_uuid1 = UUID(),
--     @user_uuid2 = UUID(),
--     @user_uuid3 = UUID();

-- INSERT INTO tenants (id)
--     VALUES (@tenant_uuid1),
--            (@tenant_uuid2);

-- INSERT INTO users (id, tenant_id, display_name)
--     VALUES (@user_uuid1, @tenant_uuid1, "test1@example.com"),
--            (@user_uuid2, @tenant_uuid1, "test2@example.com"),
--            (@user_uuid3, @tenant_uuid2, "test3@example.jp");
