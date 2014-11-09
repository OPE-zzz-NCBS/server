select id, name, client_type
from (
	select *, row_number() over (order by t.id asc) num
	from (
		select id, last_name + ', ' + first_name name, 'PERSON' client_type from dbo.Persons
		union all 
		select id, name, 'COMPANY' from dbo.Corporates
		union all
		select id, name, 'GROUP' from dbo.Groups
		union all
		select id, name, 'VILLAGE_BANK' from dbo.Villages
	) t
) t 
where t.num between ? and ?

