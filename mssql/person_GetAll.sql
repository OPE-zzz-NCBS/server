select t.id,
	convert(nvarchar(36), p.uuid) uuid,
	p.first_name,
	p.last_name,
	isnull(p.father_name, '') father_name,
	case when p.sex = 'M' then 'Male' else 'Female' end sex,
	convert(nvarchar, p.birth_date, 127) birth_date,
	isnull(p.birth_place, '') birth_place,
	isnull(p.identification_data, '') identification_data,
	isnull(p.nationality, '') nationality,
	p.activity_id,
	ti.branch_id,
	isnull(ti.home_phone, '') home_phone,
	isnull(ti.personal_phone, '') personal_phone,
	isnull(c1.id, 0) city_id1,
	isnull(ti.address, '') address1,
	isnull(ti.zipCode, '') postal_code1,
	isnull(c2.id, 0) city_id2,
	isnull(ti.secondary_address, '') address2,
	isnull(ti.secondary_zipCode, '') postal_code2,
	isnull(t.custom_field_id, 0) custom_field_id,
	isnull(t.custom_field_value, '') custom_field_value
from (
	select p.id, cfv.field_id custom_field_id, cfv.value custom_field_value, dense_rank() over (order by p.id) num
	from dbo.Persons p
	left join 
	(
		select cfv.owner_id, cfv.field_id, cfv.value
		from dbo.CustomFieldsValues cfv
		left join dbo.CustomFields cf on cf.id = cfv.field_id
		where len(value) > 0 and cf.[type] != 'Table' and cf.deleted = 0
	) cfv on cfv.owner_id = p.id
) t
left join dbo.Persons p on p.id = t.id
left join dbo.Tiers ti on ti.id = p.id
left join dbo.City c1 on c1.name = ti.city
	and c1.district_id = ti.district_id
	and c1.deleted = 0
left join dbo.City c2 on c2.name = ti.secondary_city
	and c2.district_id = ti.secondary_district_id
	and c2.deleted = 0
where t.num between ? and ?
order by t.id
