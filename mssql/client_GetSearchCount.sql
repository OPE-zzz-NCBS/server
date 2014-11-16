select count(*)
from
(
	select id, first_name + ' ' + last_name search_name from dbo.Persons
	union all
	select id, name from dbo.Corporates
	union all
	select id, name from dbo.Groups
	union all
	select id, name from dbo.Villages
) t
where t.search_name like '%' + ? + '%'
