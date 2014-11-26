select p.first_name, p.last_name, isnull(p.father_name, '') father_name,
    case when p.sex = 'M' then 'Male' else 'Female' end sex,
    convert(nvarchar, p.birth_date, 127) birth_date,
    isnull(p.birth_place, '') birth_place,
    isnull(p.identification_data, '') identification_data,
    isnull(p.nationality, '') nationality,
    p.activity_id,
    t.branch_id,
    isnull(c.id, 0) city_id,
    isnull(t.address, '') address,
    isnull(t.zipCode, '') postal_code,
    isnull(t.home_phone, '') home_phone,
    isnull(t.personal_phone, '') personal_phone
from dbo.Persons p
left join dbo.Tiers t on t.id = p.id
left join dbo.City c on c.name = t.city
    and c.district_id = t.district_id
    and c.deleted = 0
where p.id = ?
