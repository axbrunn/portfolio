CREATE TABLE IF NOT EXISTS posts (
    id         INT UNSIGNED NOT NULL AUTO_INCREMENT,
    title      VARCHAR(255) NOT NULL,
    slug       VARCHAR(255) NOT NULL,
    excerpt    TEXT         NOT NULL,
    body       LONGTEXT     NOT NULL,
    published  BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    published_at TIMESTAMP  NULL DEFAULT NULL,

    PRIMARY KEY (id),
    UNIQUE KEY uq_posts_slug (slug)
);

INSERT INTO posts (title, slug, excerpt, body, published, published_at) VALUES
(
    'Lorem Ipsum Dolor Sit Amet',
    'lorem-ipsum-dolor-sit-amet',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.\n\nDuis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.\n\nSed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.',
    TRUE,
    '2025-01-10 09:00:00'
);
