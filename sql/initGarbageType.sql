truncate go_classify.garbage_types;

insert into go_classify.garbage_types
(created_at, updated_at, deleted_at, row, name, parent_type_id, image_id, garbage_detail_id)
values (now(), now(), null, '', 'all', null, null, null);

insert into go_classify.garbage_types
(created_at, updated_at, deleted_at, row, name, parent_type_id, image_id, garbage_detail_id)
values (now(), now(), null, '', '厨余垃圾', 1, null, null);

insert into go_classify.garbage_types
(created_at, updated_at, deleted_at, row, name, parent_type_id, image_id, garbage_detail_id)
values (now(), now(), null, '', '可回收物', 1, null, null);

insert into go_classify.garbage_types
(created_at, updated_at, deleted_at, row, name, parent_type_id, image_id, garbage_detail_id)
values (now(), now(), null, '', '有害垃圾', 1, null, null);

insert into go_classify.garbage_types
(created_at, updated_at, deleted_at, row, name, parent_type_id, image_id, garbage_detail_id)
values (now(), now(), null, '', '其他垃圾', 1, null, null);