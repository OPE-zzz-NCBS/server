select first_name, last_name, isnull(father_name, '') father_name,
    case when sex = 'M' then 'Male' else 'Female' end sex,
    convert(nvarchar, birth_date, 127) birth_date,
    isnull(birth_place, '') birth_place,
    isnull(identification_data, '') identification_data,
    isnull(nationality, '') nationality,
    activity_id
from dbo.Persons where id = ?
