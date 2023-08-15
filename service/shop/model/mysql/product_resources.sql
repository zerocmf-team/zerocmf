CREATE TABLE IF NOT EXISTS product_resources
(
    id            bigint(20) AUTO_INCREMENT,
    product_id    BIGINT(20),
    resource_type TINYINT(3),
    resource_url  VARCHAR(255),
    `created_at`  bigint(20) NOT NULL DEFAULT 0,
    `updated_at`  bigint(20) NOT NULL DEFAULT 0,
    `deleted_at`  bigint(20) NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);