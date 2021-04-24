truncate go_classify.images;

insert into go_classify.images
(id, created_at, updated_at, deleted_at, row, path, url) VALUES
(1, now(), now(), null, '', '/root/image/go_classify/avatar/1779c453.png', 'http://benjaminlee.cn/nginx/image/go_classify/avatar/1779c453.png');