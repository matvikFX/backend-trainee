package storage

const (
	createBanner = `insert into banners(content, is_active)
	values($1, $2);`

	createFeature = `insert into banner_feature(banner_id, feature_id)
	values(currval('banners_id_seq'), $1);`

	createTag = `insert into banner_tag(banner_id, tag_id)
	values(currval('banners_id_seq'), $1);`

	getUserBanner = `select 
	b.content
from banners b
join banner_tag b_t on b.id = b_t.banner_id
join banner_feature b_f on b.id = b_f.banner_id
where b_t.tag_id = $1 and 
	b_f.feature_id = $2 and 
	b.is_active = true
group by b.id, b_f.feature_id;`

	getCurrentBannerID = "select currval('banners_id_seq');"

	updateBanner = `update banners
	set content=$2, is_active=$3, updated_at=now()
	where id = $1;

update banner_feature
	set feature_id = $4
	where banner_id = $1;`

	deleteBanner = `delete from banners
	where id = $1;`
)
