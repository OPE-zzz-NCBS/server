select count(*)
from
(
	select id from dbo.Persons
	union all
	select id from dbo.Corporates
	union all
	select id from dbo.Groups
	union all
	select id from dbo.Villages
) t

