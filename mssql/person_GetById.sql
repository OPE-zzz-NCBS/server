select p.first_name, p.last_name, isnull(p.father_name, '') father_name,
    case when p.sex = 'M' then 'Male' else 'Female' end sex,
    convert(nvarchar, p.birth_date, 127) birth_date,
    isnull(p.birth_place, '') birth_place,
    isnull(p.identification_data, '') identification_data,
    isnull(p.nationality, '') nationality,
    p.activity_id,
    t.branch_id,
    isnull(t.home_phone, '') home_phone,
    isnull(t.personal_phone, '') personal_phone,
    isnull(c1.id, 0) city_id1,
    isnull(t.address, '') address1,
    isnull(t.zipCode, '') postal_code1,
    isnull(c2.id, 0) city_id2,
    isnull(t.secondary_address, '') address2,
    isnull(t.secondary_zipCode, '') postal_code2
from dbo.Persons p
left join dbo.Tiers t on t.id = p.id
left join dbo.City c1 on c1.name = t.city
    and c1.district_id = t.district_id
    and c1.deleted = 0
left join dbo.City c2 on c2.name = t.secondary_city
    and c2.district_id = t.secondary_district_id
    and c2.deleted = 0
where p.id = ?
