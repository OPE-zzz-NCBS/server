select p.first_name, p.last_name, isnull(p.father_name, '') father_name,
    case when p.sex = 'M' then 'Male' else 'Female' end sex,
    convert(nvarchar, p.birth_date, 127) birth_date,
    isnull(p.birth_place, '') birth_place,
    isnull(p.identification_data, '') identification_data,
    isnull(p.nationality, '') nationality,
    p.activity_id,
    t.branch_id
from dbo.Persons p
left join dbo.Tiers t on t.id = p.id
where p.id = ?
