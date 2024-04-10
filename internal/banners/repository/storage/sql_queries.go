package storage

const (
	createBanner = `insert into banners(content, is_active) 
	values($1, $2);`

	createFeature = `insert into banner_feature(banner_id, feature_id)
	values(currval('banners_id_seq'), $1);`

	createTag = `insert into banner_tag(banner_id, tag_id)
	values(currval('banners_id_seq'), $1);`

	getCurrentBannerID = `select currval('banners_id_seq');`

	getTagsByID = `select tag_id from banner_tag where banner_id = $1;`

	getBannerWithoutTags = `select b.id, b.content, b.is_active, 
	b.created_at, b.updated_at, b_f.feature_id
from banners b
join banner_tag b_t on b.id = b_t.banner_id
join banner_feature b_f on b.id = b_f.banner_id`

	getByID = `select b.id, b.content, b.is_active,
	b.created_at, b.updated_at, b_f.feature_id
from banners b
join banner_feature b_f on b.id = b_f.banner_id
where b.id = $1;`

	getContent = `select b.content
from banners b
join banner_tag b_t on b.id = b_t.banner_id
join banner_feature b_f on b.id = b_f.banner_id
where b_t.tag_id = $1 and b_f.feature_id = $2 and is_active = true;`

	updateBanner = `update banners
	set content=$2, is_active=$3, updated_at=now()
	where id = $1;`

	updateFeature = `update banner_feature
	set feature_id = $2
	where banner_id = $1;`

	deleteBanner = `delete from banners where id = $1;`

	deleteTags = `delete from banner_tag where banner_id = $1`
)
