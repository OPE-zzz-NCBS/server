declare @id int = ?
select 'first_name' name, 'TEXT' data_type, 'First name' caption, isnull(first_name, '') value, '' extra
from dbo.Persons where id = @id

union all

select 'father_name', 'TEXT', 'Father name', isnull(father_name, ''), ''
from dbo.Persons where id = @id

union all

select 'last_name', 'TEXT', 'Last name', isnull(last_name, ''), ''
from dbo.Persons where id = @id

union all

select 'birth_date', 'DATE', 'Birth date', convert(nvarchar, birth_date, 127), ''
from dbo.Persons where id = @id

union all

select 'birth_place', 'TEXT', 'Brith place', isnull(birth_place, ''), ''
from dbo.Persons where id = @id

union all

select 'sex', 'LIST', 'Sex',
	case when sex = 'M' then 'Male' else 'Female' end,
	'Male:Female'
from dbo.Persons where id = @id

union all

select 'identification_data', 'TEXT', 'ID document', isnull(identification_data, ''), ''
from dbo.Persons where id = @id

union all

select 'nationality', 'TEXT', 'Nationality', isnull(nationality, ''), ''
from dbo.Persons where id = @id
